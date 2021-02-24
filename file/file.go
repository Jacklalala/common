package file

import (
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"common/logger"
)

const (
	atomicWriteFilePrefix              = "write-file-atomic-"
	atomicWriteFileMaxNumConflicts     = 5
	atomicWriteFileMaxNumWriteAttempts = 1000
	lcgA                               = 6364136223846793005
	lcgC                               = 1442695040888963407
	atomicWriteFileFlag                = os.O_WRONLY | os.O_CREATE | os.O_SYNC | os.O_TRUNC | os.O_EXCL
)

var (
	atomicWriteFileRand   uint64
	atomicWriteFileRandMu sync.Mutex
)

// CreateDirIfMissing creates a dir for dirPath if not already exists. If the dir is empty it returns true
func CreateDirIfMissing(dirPath string) (bool, error) {
	// if dirPath does not end with a path separator, it leaves out the last segment while creating directories
	if !strings.HasSuffix(dirPath, "/") {
		dirPath = dirPath + "/"
	}
	logDirStatus("Before creating dir", dirPath)
	err := os.MkdirAll(path.Dir(dirPath), 0755)
	if err != nil {
		return false, err
	}
	logDirStatus("After creating dir", dirPath)
	return DirEmpty(dirPath)
}

// DirEmpty returns true if the dir at dirPath is empty
func DirEmpty(dirPath string) (bool, error) {
	f, err := os.Open(dirPath)
	if err != nil {
		logger.Debugf("Error opening dir [%s]: %+v", dirPath, err)
		return false, err
	}
	defer f.Close()

	_, err = f.Readdir(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err
}

// ListSubdirs returns the subdirectories
func ListSubdirs(dirPath string) ([]string, error) {
	subdirs := []string{}
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		if f.IsDir() {
			subdirs = append(subdirs, f.Name())
		}
	}
	return subdirs, nil
}

// Exists checks whether the given file exists.
// If the file exists, this method also returns the size of the file.
func Exists(filePath string) (bool, int64, error) {
	fileInfo, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false, 0, nil
	}
	if err != nil {
		return false, 0, err
	}
	return true, fileInfo.Size(), nil
}

// WriteFileAtomic creates a temporary file with data and provided perm and
// swaps it atomically with filename if successful.
func WriteFileAtomic(filename string, data []byte, perm os.FileMode) (err error) {
	var (
		dir = filepath.Dir(filename)
		f   *os.File
	)

	nconflict := 0
	i := 0
	for ; i < atomicWriteFileMaxNumWriteAttempts; i++ {
		name := filepath.Join(dir, atomicWriteFilePrefix+randWriteFileSuffix())
		f, err = os.OpenFile(name, atomicWriteFileFlag, perm)
		if os.IsExist(err) {
			if nconflict++; nconflict > atomicWriteFileMaxNumConflicts {
				atomicWriteFileRandMu.Lock()
				atomicWriteFileRand = writeFileRandReseed()
				atomicWriteFileRandMu.Unlock()
			}
			continue
		} else if err != nil {
			return err
		}
		break
	}
	if i == atomicWriteFileMaxNumWriteAttempts {
		return errors.New("Could not create atomic write file after %d attempts")
	}
	defer os.Remove(f.Name())
	defer f.Close()

	if n, err := f.Write(data); err != nil {
		return err
	} else if n < len(data) {
		return io.ErrShortWrite
	}
	f.Close()

	return os.Rename(f.Name(), filename)
}

func randWriteFileSuffix() string {
	atomicWriteFileRandMu.Lock()
	r := atomicWriteFileRand
	if r == 0 {
		r = writeFileRandReseed()
	}
	r = r*lcgA + lcgC
	atomicWriteFileRand = r
	atomicWriteFileRandMu.Unlock()
	// Can have a negative name, replace this in the following
	suffix := strconv.Itoa(int(r))
	if string(suffix[0]) == "-" {
		suffix = strings.Replace(suffix, "-", "0", 1)
	}
	return suffix
}

func writeFileRandReseed() uint64 {
	return uint64(time.Now().UnixNano() + int64(os.Getpid()<<20))
}

func logDirStatus(msg string, dirPath string) {
	exists, _, err := Exists(dirPath)
	if err != nil {
		logger.Errorf("Error checking for dir existence")
	}
	if exists {
		logger.Debugf("%s - [%s] exists", msg, dirPath)
	} else {
		logger.Debugf("%s - [%s] does not exist", msg, dirPath)
	}
}
