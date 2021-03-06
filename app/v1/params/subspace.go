package params

import (
	"github.com/irisnet/irishub/app/v1/params/subspace"
)

// re-export types from subspace
type (
	Subspace         = subspace.Subspace
	ReadOnlySubspace = subspace.ReadOnlySubspace
	ParamSet         = subspace.ParamSet
	ParamSets        = subspace.ParamSets
	KeyValuePairs    = subspace.KeyValuePairs
	TypeTable        = subspace.TypeTable
)

// re-export functions from subspace
func NewTypeTable(keytypes ...interface{}) TypeTable {
	return subspace.NewTypeTable(keytypes...)
}
