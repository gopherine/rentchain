package types

const (
	// ModuleName defines the module name
	ModuleName = "assets"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_assets"
)

var (
	ParamsKey = []byte("p_assets")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
