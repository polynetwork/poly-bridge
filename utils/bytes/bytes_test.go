package bytes

import (
	"github.com/polynetwork/poly/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseRune(t *testing.T) {
	// reverse poly hash
	origin := "fc9b1e82fca03cecae5499a07206797345d398ae98f2369c39a0ac2de034ebc2"
	hash, err := common.Uint256FromHexString(origin)
	assert.NoError(t, err)

	reversed := ReverseRune(hash[:])
	final, err := common.Uint256ParseFromBytes(reversed)
	assert.NoError(t, err)

	t.Logf("final hash %s", final.ToHexString())
}
