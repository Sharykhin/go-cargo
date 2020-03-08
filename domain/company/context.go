package company

type (
	//ContextKey is an alias of string that wiil be used with context.WithValue to prevent potential collisions
	ContextKey string
)

var (
	// SQLTransactionTX it's a context key regarding sql transaction struct
	SQLTransactionTX = ContextKey("sql-transaction-tx")
)
