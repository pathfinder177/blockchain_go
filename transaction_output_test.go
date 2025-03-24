package main

import (
	"reflect"
	"testing"
)

func TestNewTXOutput(t *testing.T) {
	type args struct {
		value   int
		address string
	}
	tests := []struct {
		name string
		args args
		want *TXOutput
	}{
		{
			name: "constructor",
			args: args{
				value:   5,
				address: "1AYmpvb95P8m7je8SZzTnLx7Z6sjRr6PK",
			},
			want: &TXOutput{
				Value: 5,
				PubKeyHash: []byte{
					0x01, 0xCE, 0x44, 0x2C, 0xDE, 0x67, 0xDA, 0x83,
					0xC8, 0x7B, 0x08, 0x7C, 0x1F, 0x4B, 0x23, 0xF2,
					0x77, 0xE1, 0xA4, 0x12,
				}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTXOutput(tt.args.value, tt.args.address); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTXOutput() = %v, want %v", got, tt.want)
			}
		})
	}
}
