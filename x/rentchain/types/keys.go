package types

const (
	// ModuleName defines the module name
	ModuleName = "rentchain"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_rentchain"
)

var (
	ParamsKey = []byte("p_rentchain")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
