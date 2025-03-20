package legacybech32

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/ledger"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Deprecated: TestBeach32ifPbKey exists only for backward compatibility with legacy bech32
// pubkey formats. It will be removed in a future release along with the rest of this package.
//
// This test requires both 'ledger' and 'test_ledger_mock' build tags to run, otherwise it will be skipped.
// The ledger mock will use the test mnemonic to generate keys that match the hardcoded test vector.
//
// For convenience:
// go test -tags "ledger test_ledger_mock" ./types/bech32/legacybech32/... -v
func TestBeach32ifPbKey(t *testing.T) {
	_, err := ledger.NewPrivKeySecp256k1Unsafe(*hd.NewFundraiserParams(0, sdk.CoinType, 0))
	if err != nil {
		t.Skip("ledger support is not available")
	}

	require := require.New(t)
	path := *hd.NewFundraiserParams(0, sdk.CoinType, 0)
	priv, err := ledger.NewPrivKeySecp256k1Unsafe(path)
	require.Nil(err, "%s", err)
	require.NotNil(priv)

	pubKeyAddr, err := MarshalPubKey(AccPK, priv.PubKey())
	require.NoError(err)
	require.Equal("cosmospub1addwnpepqd87l8xhcnrrtzxnkql7k55ph8fr9jarf4hn6udwukfprlalu8lgw0urza0",
		pubKeyAddr, "Is your device using test mnemonic: %s ?", testdata.TestMnemonic)
}
