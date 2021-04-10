package leveldb

import (
	"fmt"

	"github.com/btcsuite/goleveldb/leveldb"
)

type LevelDBImpl struct {
	db   *leveldb.DB
	name string
}

func NewLevelDBInstance(dir string) *LevelDBImpl {
	d := &LevelDBImpl{}
	d.name = dir
	db, err := leveldb.OpenFile(dir, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("open leveldb %s\r\n", dir)
	d.db = db
	return d
}

func (d *LevelDBImpl) Name() string {
	return d.name
}

func (d *LevelDBImpl) Set(k, v []byte) error {
	return d.db.Put(k, v, nil)
}

func (d *LevelDBImpl) Get(k []byte) ([]byte, error) {
	return d.db.Get(k, nil)
}
