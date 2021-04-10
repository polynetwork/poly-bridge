package leveldb

import (
	"testing"
)

func TestInsertStat(t *testing.T) {
	cfg := &Config{
		Ip:   "172.168.3.219",
		Port: 3306,
		User: "root",
		Pwd:  "123456",
		Db:   "txstat",
	}
	NewLevelDBInstance(cfg)
	err := InsertStat("127.0.0.1", 30001, "0xjflaksdjfoi23rnlasdf", 10, 10)
	if err != nil {
		t.Fatal(err)
	}
}
