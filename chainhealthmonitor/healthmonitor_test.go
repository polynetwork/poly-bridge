package chainhealthmonitor

import (
	"poly-bridge/conf"
	"testing"
)

func TestEthNodeMonitor(t *testing.T) {
	config := conf.NewConfig("../prod.json")
	type args struct {
		config *conf.Config
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{config: config},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			EthNodeMonitor(tt.args.config)
		})
	}
}
