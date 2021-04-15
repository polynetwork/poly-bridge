package params

import (
	"poly-bridge/models"
)

type FetchRequestParams struct {
	TokenId *models.BigInt
	Url     string
}

