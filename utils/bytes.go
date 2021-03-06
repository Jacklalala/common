package utils

import (
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"unsafe"
)

func ToHex(b []byte) string {
	hex := bytes2Hex(b)
	// Prefer output of "0x0" instead of "0x"
	if len(hex) == 0 {
		hex = "0"
	}
	return "0x" + hex
}

func FromHex(s string) []byte {
	if len(s) > 1 {
		if s[0:2] == "0x" || s[0:2] == "0X" {
			s = s[2:]
		}
	}
	if len(s)%2 == 1 {
		s = "0" + s
	}
	return hex2Bytes(s)
}

// Copy bytes
//
// Returns an exact copy of the provided bytes
func CopyBytes(b []byte) (copiedBytes []byte) {
	if b == nil {
		return nil
	}
	copiedBytes = make([]byte, len(b))
	copy(copiedBytes, b)

	return
}

func hasHexPrefix(str string) bool {
	return len(str) >= 2 && str[0] == '0' && (str[1] == 'x' || str[1] == 'X')
}

func isHexCharacter(c byte) bool {
	return ('0' <= c && c <= '9') || ('a' <= c && c <= 'f') || ('A' <= c && c <= 'F')
}

func isHex(str string) bool {
	if len(str)%2 != 0 {
		return false
	}
	for _, c := range []byte(str) {
		if !isHexCharacter(c) {
			return false
		}
	}
	return true
}

func bytes2Hex(d []byte) string {
	return hex.EncodeToString(d)
}

func hex2Bytes(str string) []byte {
	h, _ := hex.DecodeString(str)
	return h
}

func Hex2BytesFixed(str string, flen int) []byte {
	h, _ := hex.DecodeString(str)
	if len(h) == flen {
		return h
	}
	if len(h) > flen {
		return h[len(h)-flen:]
	}
	hh := make([]byte, flen)
	copy(hh[flen-len(h):flen], h[:])
	return hh
}

func RightPadBytes(slice []byte, l int) []byte {
	if l <= len(slice) {
		return slice
	}

	padded := make([]byte, l)
	copy(padded, slice)

	return padded
}

func LeftPadBytes(slice []byte, l int) []byte {
	if l <= len(slice) {
		return slice
	}

	padded := make([]byte, l)
	copy(padded[l-len(slice):], slice)

	return padded
}

func StringToPointer(str string) (uint64, error) {
	strPointerInt := fmt.Sprintf("%d", unsafe.Pointer(&str))
	return strconv.ParseUint(strPointerInt, 10, 0)
}

func PointerToString(pointer uint64) string {
	var s *string
	s = *(**string)(unsafe.Pointer(&pointer))
	str := *(*string)(unsafe.Pointer(s))
	b := CopyBytes([]byte(str))
	return string(b)
}

func Uint64ToBytes(value uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, value)
	return b
}

func Uint64SetBytes(data []byte) uint64 {
	index := binary.BigEndian.Uint64(data)
	return index
}

func Uint32ToBytes(value uint32) []byte {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, value)
	return b
}

func Uint32SetBytes(data []byte) uint32 {
	index := binary.BigEndian.Uint32(data)
	return index
}

func JsonString(v interface{}, format bool) string {
	if format {
		data, err := json.MarshalIndent(v, "", "    ")
		if err != nil {
			fmt.Println(err)
		}
		return string(data)
	} else {
		data, err := json.Marshal(v)
		if err != nil {
			fmt.Println(err)
		}
		return string(data)
	}
}
