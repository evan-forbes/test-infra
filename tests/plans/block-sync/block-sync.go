package blocksync

import (
	"context"

	"github.com/celestiaorg/test-infra/testkit"
	blocksynclatest "github.com/celestiaorg/test-infra/tests/helpers/block-sync/latest"
	"github.com/testground/sdk-go/run"
	"github.com/testground/sdk-go/runtime"
)

// BlockSyncLatest represents a testcase of W validator, X bridges, Y full nodes that
// are trying to sync the latest block from Bridge Nodes and among themselves
// using either ShrexGetter only, IPLDGetter only or the default CascadeGetter (_see compositions/cluster-k8s/blocksync-latest/*/*-{getter}.toml)
// More information under docs/test-plans/005-Block-Sync
func BlockSyncLatest(runenv *runtime.RunEnv, initCtx *run.InitContext) (err error) {
	switch runenv.StringParam("role") {
	case "validator":
		err = blocksynclatest.RunValidator(runenv, initCtx)
	}

	if err != nil {
		runenv.RecordFailure(err)
		initCtx.SyncClient.MustSignalAndWait(context.Background(), testkit.FinishState, runenv.TestInstanceCount)
		return err
	}

	runenv.RecordSuccess()
	return nil
}
