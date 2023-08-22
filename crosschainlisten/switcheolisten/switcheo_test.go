package switcheolisten

import (
	"encoding/json"
	"fmt"
	"poly-bridge/conf"
	"testing"
)

func Test_switcheo_listen(t *testing.T) {
	swthlis := NewSwitcheoChainListen(&conf.ChainListenConfig{
		ChainName:  "switcheo",
		ChainId:    5,
		ListenSlot: 5,
		Nodes: []*conf.Restful{
			{
				Url: "https://rpc.carbon.intsol.guru:443",
			},
		},
	})
	for _, height := range []uint64{45740156, 45753484} {
		fmt.Println("----------", height, "--------")
		wrapperTransactions, srcTransactions, polyTransactions, dstTransactions, wrapperDetails, polyDetails, locks, unlocks, err := swthlis.HandleNewBlock(height)
		if err != nil {
			t.Fatal(fmt.Sprintf("HandleNewBlock %d err: %v", height, err))
		}
		jsonMarshalIndent(wrapperTransactions)
		jsonMarshalIndent(srcTransactions)
		jsonMarshalIndent(polyTransactions)
		jsonMarshalIndent(dstTransactions)
		jsonMarshalIndent(wrapperDetails)
		jsonMarshalIndent(polyDetails)
		fmt.Println(locks, unlocks)
	}
}

func jsonMarshalIndent(x interface{}) {
	jsonx, _ := json.MarshalIndent(x, "", "	")
	fmt.Println(string(jsonx))
}
