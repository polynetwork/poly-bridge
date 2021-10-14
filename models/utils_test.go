package models

import (
	"poly-bridge/conf"
	"testing"
)

func TestGetL1BlockNumberOfArbitrumTx(t *testing.T) {
	type args struct {
		hash string
	}
	test := struct {
		name    string
		args    args
		want    uint64
		wantErr bool
	}{
		name: "TestGetL1BlockNumberOfArbitrumTx",
		args: args{
			"94bed197d7fabd52c2abbf01c4daf54231dfba68d1fe7e4c388aa43860a29314",
		},
		want: 9457494,
	}
	config := conf.NewConfig("./../conf/config_testnet.json")
	conf.GlobalConfig = config

	t.Run(test.name, func(t *testing.T) {
		got, err := GetL1BlockNumberOfArbitrumTx(test.args.hash)
		if (err != nil) != test.wantErr {
			t.Errorf("GetL1BlockNumberOfArbitrumTx() error = %v, wantErr %v", err, test.wantErr)
			return
		}
		if got != test.want {
			t.Errorf("GetL1BlockNumberOfArbitrumTx() got = %v, want %v", got, test.want)
		}
	})
}
