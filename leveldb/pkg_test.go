package leveldb

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testDBPath = "./leveldb"

type testDBEnv struct {
	t    *testing.T
	path string
	db   *DB
}

type testDBProviderEnv struct {
	t        *testing.T
	path     string
	provider *Provider
}

func newTestDBEnv(t *testing.T, path string) *testDBEnv {
	testDBEnv := &testDBEnv{t: t, path: path}
	testDBEnv.cleanup()
	testDBEnv.db = CreateDB(&Conf{path})
	return testDBEnv
}

func newTestProviderEnv(t *testing.T, path string) *testDBProviderEnv {
	testProviderEnv := &testDBProviderEnv{t: t, path: path}
	testProviderEnv.cleanup()
	testProviderEnv.provider = NewProvider(&Conf{path})
	return testProviderEnv
}

func (dbEnv *testDBEnv) cleanup() {
	if dbEnv.db != nil {
		dbEnv.db.Close()
	}
	assert.NoError(dbEnv.t, os.RemoveAll(dbEnv.path))
}

func (providerEnv *testDBProviderEnv) cleanup() {
	if providerEnv.provider != nil {
		providerEnv.provider.Close()
	}
	assert.NoError(providerEnv.t, os.RemoveAll(providerEnv.path))
}
