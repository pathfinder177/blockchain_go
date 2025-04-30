package entity

type HistoricalTransaction struct {
	From      string
	To        string
	Currency  string
	Amount    int
	Timestamp int64
	// Status    string
}
