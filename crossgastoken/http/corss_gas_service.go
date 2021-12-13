package http

import (
	"poly-bridge/crossgastoken/gasmodels"

func MakeGetChainResp(crossNativeMaps []*gasmodels.CrossGasMap) []*gasmodels.GetChainResp {
	getChainResps := make([]*gasmodels.GetChainResp, 0)
	for _, v := range crossNativeMaps {
		getChainResps = append(getChainResps, &gasmodels.GetChainResp{
			v.SrcTokenHash,
			v.DstToken.TokenBasicName,
			v.DstChainId,
			v.DstTokenHash,
			v.DstToken.TokenBasicName,
		})
	}
	return getChainResps
}