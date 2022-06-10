package standard

import (
	"encoding/json"
	"poly-bridge/models"
)

type Profile struct {
	Image       string        `json:"image"`
	ExternalUrl string        `json:"external_url"`
	Description string        `json:"description,omitempty"`
	Name        string        `json:"name"`
	Attributes  []interface{} `json:"attributes,omitempty"`
}

func (p *Profile) Marshal() ([]byte, error) {
	return json.Marshal(p)
}

func (p *Profile) Unmarshal(raw []byte) error {
	return json.Unmarshal(raw, p)
}

func (p *Profile) Convert(assetName string, tokenId string, url string) (*models.NFTProfile, error) {
	np := new(models.NFTProfile)
	np.TokenBasicName = assetName
	np.Name = p.Name
	np.Url = url
	np.Image = p.Image
	np.Description = p.Description
	np.NftTokenId = tokenId

	raw, err := p.Marshal()
	if err != nil {
		return nil, err
	}
	np.Text = string(raw)
	return np, nil
}
