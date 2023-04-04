package main

import (
	bigblocks "github.com/celestiaorg/test-infra/tests/plans/big-blocks"
	pfdgsbn "github.com/celestiaorg/test-infra/tests/plans/pfd-gsbn"
	"github.com/testground/sdk-go/run"
)

var testcases = map[string]interface{}{
	// Big Blocks Plan
	"001-val-large-txs": bigblocks.ValSubmitLargeTxs,
	// Pay For Blob & Get Shares by Namespace Plan
	// PayForBlobAndGetShares is tracking TestCase key to know
	// when to do shares checker scenario
	"pay-for-blob":            pfdgsbn.PayForBlobAndGetShares,
	"get-shares-by-namespace": pfdgsbn.PayForBlobAndGetShares,
}

func main() {
	run.InvokeMap(testcases)
}
