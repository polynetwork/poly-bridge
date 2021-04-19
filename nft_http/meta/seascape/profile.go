package seascape

import (
	"encoding/json"
	"fmt"
	"math/big"
	"poly-bridge/models"
)

type Profile struct {
	Image       string        `json:"image"`
	ExternalUrl string        `json:"external_url"`
	Description string        `json:"description"`
	Name        string        `json:"name"`
	Attributes  []interface{} `json:"attributes,omitempty"`
}

func (p *Profile) Marshal() ([]byte, error) {
	return json.Marshal(p)
}

func (p *Profile) Unmarshal(raw []byte) error {
	return json.Unmarshal(raw, p)
}

func (p *Profile) Convert(assetName string, tokenId string) (*models.NFTProfile, error) {
	tid, ok := new(big.Int).SetString(tokenId, 10)
	if !ok {
		return nil, fmt.Errorf("invalid token id string %s", tokenId)
	}
	np := new(models.NFTProfile)
	np.TokenBasicName = assetName
	np.Name = p.Name
	np.Url = p.ExternalUrl
	np.Image = p.Image
	np.Description = p.Description
	np.NftTokenId = models.NewBigInt(tid)

	raw, err := p.Marshal()
	if err != nil {
		return nil, err
	}
	np.Text = string(raw)
	return np, nil
}
