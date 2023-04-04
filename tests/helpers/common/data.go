package common

import (
	"bytes"

	appns "github.com/celestiaorg/celestia-app/pkg/namespace"
	"github.com/celestiaorg/nmt/namespace"
	tmrand "github.com/tendermint/tendermint/libs/rand"
)

// DefaultNameId is used in cases where we only have 1 Namespace.ID used
// across all nodes that submit pfd and get shares by this ID
var DefaultNameId = namespace.ID{100, 100, 150, 150, 200, 200, 250, 255}

// TODO(@Bidon15): We need to start testing gas mechanism sooner than later
const gasLimit uint64 = 2000000

// GetRandomNamespace returns a random namespace.ID per each call made by
// each instance of node type
func GetRandomNamespace() namespace.ID {
	for {
		s := tmrand.Bytes(8)
		if bytes.Compare(s, appns.MaxReservedNamespace.ID) > 0 {
			return s
		}
	}
}

// GenerateNamespaceID returns a namespace ID based on runenv.StringParams defined in the composition file
// TODO(@Bidon15): We actually need to refactor this out using runenv.IntParam()
func GenerateNamespaceID(amount string) namespace.ID {
	if amount == "1" {
		return DefaultNameId
	}
	return GetRandomNamespace()
}

// GetRandomMessageBySize returns a random []byte per each call made by
// each instance of node type. The size is defined in the .toml file
func GetRandomMessageBySize(size int) []byte {
	return tmrand.Bytes(size)
}
