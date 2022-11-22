package types

const (
	// ModuleName defines the module name
	ModuleName = "lending"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_lending"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	PairPoolKey      = "PairPool/value/"
	PairPoolCountKey = "PairPool/count/"
)

const (
	PoolEventId     = "id"
	PoolEventAmount = "amount"

	PoolAddedEventType        = "pool-added"
	PoolAddedEventAssetDenom  = "asset-denom"
	PoolAddedEventShadowDenom = "shadow-denom"

	PoolDepositedEventType = "pool-deposited"

	PoolWithdrawedEventType = "pool-withdrawed"

	PoolBorrowedEventType = "pool-borrowed"

	PoolRepayedEventType = "pool-repayed"
)
