package entity

type Transaction struct {
	From      string
	To        string
	Asset     string
	Amount    int
	Timestamp int64
	Status    string
}
