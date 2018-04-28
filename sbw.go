package sbw

import "time"

// Account types.
const (
	AccountTypeChecking = "checking"
	AccountTypeSavings  = "saving"
	AccountTypeCredit   = "credit_card"
	AccountTypeLoan     = "loan"
)

// Account represents an account.
type Account struct {
	Merchant    string     `json:"merchant"` // i.e. AmericanExpress
	Name        string     `json:"name"`     // i.e. Double Rewards
	Type        string     `json:"type"`     // credit_card
	Balance     float64    `json:"balance"`
	LastUpdated *time.Time `json:"last_updated"`
}

// Bill ...
type Bill struct {
	Name      string     `json:"name"`
	Type      string     `json:"type"` // utility, loan, credit card, COGS
	Amount    float64    `json:"amount"`
	Paid      bool       `json:"paid,omitempty"`
	DueDate   *time.Time `json:"due_date"`
	Frequency string     `json:"frequency,omitempty"` // monthly, weekly, bimonthly
}

// Transaction represents a bank or credit card transaction.
type Transaction struct {
	Date        time.Time `json:"date"`
	Merchant    string    `json:"lender"`
	Type        string    `json:"type"` // TODO debt, credit, refund, etc
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
}

// LoanPayment ...
// type LoanPayment struct {
// 	*Bill
// 	MonthsRemaining int // if applicable
// }
