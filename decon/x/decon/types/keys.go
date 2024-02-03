package types

const (
	// ModuleName defines the module name
	ModuleName = "decon"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_decon"
)

var (
	ParamsKey = []byte("p_decon")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
