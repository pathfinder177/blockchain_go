package main

import (
	"reflect"
	"testing"
)

func TestNewCoinbaseTX(t *testing.T) {
	type args struct {
		to   string
		data string
	}
	tests := []struct {
		name string
		args args
		want *Transaction
	}{
		{
			name: "constructor",
			args: args{
				to:   "1AYmpvb95P8m7je8SZzTnLx7Z6sjRr6PK",
				data: "random_data",
			},
			want: &Transaction{
				ID: []byte{
					0x56, 0xE9, 0xC7, 0xE1, 0xC3, 0xDA, 0x4F, 0xCD,
					0x6E, 0x04, 0x9F, 0x0D, 0xF0, 0xC4, 0x14, 0x64,
					0x39, 0x72, 0x77, 0x02, 0xC5, 0xFC, 0x76, 0x58,
					0x2C, 0xCF, 0xFD, 0xEC, 0x12, 0x5B, 0x4A, 0xC8,
				},
				Vin:  []TXInput{{[]byte{}, -1, nil, []byte("random_data")}},
				Vout: []TXOutput{*NewTXOutput(subsidy, "1AYmpvb95P8m7je8SZzTnLx7Z6sjRr6PK")},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCoinbaseTX(tt.args.to, tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCoinbaseTX() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransaction_IsCoinbase(t *testing.T) {
	type fields struct {
		ID   []byte
		Vin  []TXInput
		Vout []TXOutput
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "want_true",
			fields: fields{
				ID: []byte{
					0x56, 0xE9, 0xC7, 0xE1, 0xC3, 0xDA, 0x4F, 0xCD,
					0x6E, 0x04, 0x9F, 0x0D, 0xF0, 0xC4, 0x14, 0x64,
					0x39, 0x72, 0x77, 0x02, 0xC5, 0xFC, 0x76, 0x58,
					0x2C, 0xCF, 0xFD, 0xEC, 0x12, 0x5B, 0x4A, 0xC8,
				},
				Vin:  []TXInput{{[]byte{}, -1, nil, []byte("random_data")}},
				Vout: []TXOutput{*NewTXOutput(subsidy, "1AYmpvb95P8m7je8SZzTnLx7Z6sjRr6PK")},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := Transaction{
				ID:   tt.fields.ID,
				Vin:  tt.fields.Vin,
				Vout: tt.fields.Vout,
			}
			if got := tx.IsCoinbase(); got != tt.want {
				t.Errorf("Transaction.IsCoinbase() = %v, want %v", got, tt.want)
			}
		})
	}
}
