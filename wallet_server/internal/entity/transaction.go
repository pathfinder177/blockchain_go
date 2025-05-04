package entity

import "fmt"

type HistoricalTransaction struct {
	From      string
	To        string
	Currency  string
	Amount    int
	Timestamp int64
	// Status    string
}

func (ht HistoricalTransaction) String() string {
	return fmt.Sprintf("Timestamp %d, From %s, To %s, Amount %d, Currency %s", ht.Timestamp, ht.From, ht.To, ht.Amount, ht.Currency)
}
