package main

import (
	"reflect"
	"testing"
)

func TestIntToHex(t *testing.T) {
	type args struct {
		num int64
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "timestamp",
			args: args{
				num: 1711292345,
			},
			want: []byte{0x00, 0x00, 0x00, 0x00, 0x66, 0x00, 0x3f, 0xb9},
		},
		{
			name: "target_bits",
			args: args{
				num: int64(targetBits),
			},
			want: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x08},
		},
		{
			name: "nonce",
			args: args{
				num: int64(1),
			},
			want: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntToHex(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntToHex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseBytes(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "nonce",
			args: args{
				data: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01},
			},
			want: []byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReverseBytes(tt.args.data)
			if got := tt.args.data; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
