package expense

import "errors"

var (
	ErrAmountIsZero = errors.New("amount cannot be 0.0")
)

type Expense struct {
	amount float64
	desc   string
}
