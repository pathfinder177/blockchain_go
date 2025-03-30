package tests

import (
	"blockchain/internal/app"
	"crypto/ecdsa"
	"reflect"
	"testing"
)

var subsidy int = 10

// func TestNewCoinbaseTX(t *testing.T) {
// 	subsidy := 10

// 	type args struct {
// 		to   string
// 		data string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want *app.Transaction
// 	}{
// 		{
// 			name: "constructor",
// 			args: args{
// 				to:   "1AYmpvb95P8m7je8SZzTnLx7Z6sjRr6PK",
// 				data: "random_data",
// 			},
// 			want: &app.Transaction{
// 				ID: []byte{
// 					0x56, 0xE9, 0xC7, 0xE1, 0xC3, 0xDA, 0x4F, 0xCD,
// 					0x6E, 0x04, 0x9F, 0x0D, 0xF0, 0xC4, 0x14, 0x64,
// 					0x39, 0x72, 0x77, 0x02, 0xC5, 0xFC, 0x76, 0x58,
// 					0x2C, 0xCF, 0xFD, 0xEC, 0x12, 0x5B, 0x4A, 0xC8,
// 				},
// 				Vin: []app.TXInput{
// 					{
// 						Txid:      []byte{},
// 						Vout:      -1,
// 						Signature: nil,
// 						PubKey:    []byte("random_data")},
// 				},
// 				Vout: []app.TXOutput{*app.NewTXOutput(subsidy, "1AYmpvb95P8m7je8SZzTnLx7Z6sjRr6PK")},
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := app.NewCoinbaseTX(tt.args.to, tt.args.data); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("NewCoinbaseTX() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestTransaction_IsCoinbase(t *testing.T) {
	type fields struct {
		ID   []byte
		Vin  []app.TXInput
		Vout []app.TXOutput
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
				Vin: []app.TXInput{
					{
						Txid:      []byte{},
						Vout:      -1,
						Signature: nil,
						PubKey:    []byte("random_data")},
				},
				Vout: []app.TXOutput{*app.NewTXOutput(subsidy, "1AYmpvb95P8m7je8SZzTnLx7Z6sjRr6PK")},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := app.Transaction{
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

// func TestTransaction_TrimmedCopy(t *testing.T) {
// 	type fields struct {
// 		ID   []byte
// 		Vin  []app.TXInput
// 		Vout []app.TXOutput
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		want   app.Transaction
// 	}{
// 		{
// 			name: "want_true",
// 			fields: fields{
// 				ID: []byte{
// 					0x56, 0xE9, 0xC7, 0xE1, 0xC3, 0xDA, 0x4F, 0xCD,
// 					0x6E, 0x04, 0x9F, 0x0D, 0xF0, 0xC4, 0x14, 0x64,
// 					0x39, 0x72, 0x77, 0x02, 0xC5, 0xFC, 0x76, 0x58,
// 					0x2C, 0xCF, 0xFD, 0xEC, 0x12, 0x5B, 0x4A, 0xC8,
// 				},
// 				Vin: []app.TXInput{
// 					{
// 						Txid:      []byte{},
// 						Vout:      -1,
// 						Signature: nil,
// 						PubKey:    []byte("random_data")},
// 				},
// 				Vout: []app.TXOutput{*app.NewTXOutput(subsidy, "1AYmpvb95P8m7je8SZzTnLx7Z6sjRr6PK")},
// 			},
// 			want: app.Transaction{
// 				ID: []byte{
// 					0x56, 0xE9, 0xC7, 0xE1, 0xC3, 0xDA, 0x4F, 0xCD,
// 					0x6E, 0x04, 0x9F, 0x0D, 0xF0, 0xC4, 0x14, 0x64,
// 					0x39, 0x72, 0x77, 0x02, 0xC5, 0xFC, 0x76, 0x58,
// 					0x2C, 0xCF, 0xFD, 0xEC, 0x12, 0x5B, 0x4A, 0xC8,
// 				},
// 				Vin: []app.TXInput{
// 					{Txid: []byte{},
// 						Vout:      -1,
// 						Signature: nil,
// 						PubKey:    []byte("random_data")},
// 				},
// 				Vout: []app.TXOutput{*app.NewTXOutput(subsidy, "1AYmpvb95P8m7je8SZzTnLx7Z6sjRr6PK")},
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tx := &app.Transaction{
// 				ID:   tt.fields.ID,
// 				Vin:  tt.fields.Vin,
// 				Vout: tt.fields.Vout,
// 			}
// 			if got := tx.TrimmedCopy(); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Transaction.TrimmedCopy() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// Keys are mocked
// PrivateKey is mocked and it is in little-endian order.
// To overcome off-by-one error, PubKey in TXInput was changed to get proper X value of PubKey
func TestTransaction_Sign_Verify(t *testing.T) {
	type fields struct {
		ID   []byte
		Vin  []app.TXInput
		Vout []app.TXOutput
	}
	type args struct {
		privKey ecdsa.PrivateKey
		prevTXs map[string]app.Transaction
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "want_true",
			want: true,
			fields: fields{
				ID: []uint8{
					0x8C, 0xEA, 0xD9, 0x66, 0x9E, 0x47, 0x94, 0x92,
					0x51, 0xB5, 0x76, 0x11, 0x71, 0x5E, 0x5F, 0xB5,
					0x4A, 0x75, 0x76, 0x6E, 0xCD, 0xC1, 0xBE, 0x6B,
					0x17, 0x21, 0xF4, 0xB8, 0xE9, 0x00, 0xD2, 0x91,
				},
				Vin: []app.TXInput{
					{
						Txid: []uint8{
							0x8A, 0x44, 0x03, 0xA1, 0x06, 0xD5, 0xD5, 0x31,
							0xA7, 0x9C, 0x38, 0xA8, 0xCE, 0xC3, 0x5C, 0x92,
							0x1F, 0x23, 0x8F, 0x85, 0x28, 0x96, 0x31, 0x52,
							0xDF, 0x24, 0x46, 0x3C, 0x54, 0xD8, 0xA4, 0xCB,
						},
						Vout:      0,
						Signature: []uint8{},
						PubKey: []uint8{
							0xD2, 0xF3, 0x44, 0xE7, 0x45, 0x56, 0x33, 0x12,
							0xC9, 0xD4, 0x16, 0xF9, 0xB3, 0x40, 0x34, 0xDC,
							0x96, 0x37, 0x5A, 0x4D, 0x79, 0x91, 0xBB, 0x6E,
							0xFD, 0xFE, 0xFF, 0xF9, 0x4A, 0x13, 0x95, 0x28,
							0xAD, 0xA8, 0x06, 0x46, 0xB3, 0x43, 0x79, 0xA2,
							0x18, 0xA9, 0xD5, 0x13, 0x32, 0x7B, 0x15, 0x4D,
							0xAA, 0xFC, 0x6F, 0x54, 0xEE, 0xDC, 0x12, 0x7A,
							0x20, 0x34, 0xAA, 0xA3, 0x08, 0x15, 0x5E, 0xB4,
						},
					},
				},
				Vout: []app.TXOutput{
					{
						Value: 1,
						PubKeyHash: []uint8{
							0x20, 0x65, 0x7D, 0xC8, 0xFB, 0x2E, 0xFF, 0x3D,
							0xBC, 0xEE, 0x67, 0x18, 0x85, 0xE5, 0x9C, 0x70,
							0xA1, 0x4A, 0xF3, 0x75,
						},
					},
					{
						Value: 9,
						PubKeyHash: []uint8{
							0x4B, 0xD8, 0xC9, 0x8D, 0xD6, 0x27, 0x2A, 0xE5,
							0x2D, 0x49, 0x3B, 0x54, 0x80, 0x27, 0x65, 0x10,
							0xE8, 0x2B, 0x19, 0xEB,
						},
					},
				},
			},
			args: args{
				privKey: ecdsa.PrivateKey{},
				prevTXs: map[string]app.Transaction{
					"8a4403a106d5d531a79c38a8cec35c921f238f8528963152df24463c54d8a4cb": {
						ID: []uint8{
							0x8A, 0x44, 0x03, 0xA1, 0x06, 0xD5, 0xD5, 0x31,
							0xA7, 0x9C, 0x38, 0xA8, 0xCE, 0xC3, 0x5C, 0x92,
							0x1F, 0x23, 0x8F, 0x85, 0x28, 0x96, 0x31, 0x52,
							0xDF, 0x24, 0x46, 0x3C, 0x54, 0xD8, 0xA4, 0xCB,
						},
						Vin: []app.TXInput{
							{
								Txid:      []uint8{},
								Vout:      -1,
								Signature: []uint8{},
								PubKey: []uint8{
									0x54, 0x68, 0x65, 0x20, 0x54, 0x69, 0x6D, 0x65,
									0x73, 0x20, 0x30, 0x33, 0x2F, 0x4A, 0x61, 0x6E,
									0x2F, 0x32, 0x30, 0x30, 0x39, 0x20, 0x43, 0x68,
									0x61, 0x6E, 0x63, 0x65, 0x6C, 0x6C, 0x6F, 0x72,
									0x20, 0x6F, 0x6E, 0x20, 0x62, 0x72, 0x69, 0x6E,
									0x6B, 0x20, 0x6F, 0x66, 0x20, 0x73, 0x65, 0x63,
									0x6F, 0x6E, 0x64, 0x20, 0x62, 0x61, 0x69, 0x6C,
									0x6F, 0x75, 0x74, 0x20, 0x66, 0x6F, 0x72, 0x20,
									0x62, 0x61, 0x6E, 0x6B, 0x73,
								},
							},
						},
						Vout: []app.TXOutput{
							{
								Value: 10,
								PubKeyHash: []uint8{
									0x4B, 0xD8, 0xC9, 0x8D, 0xD6, 0x27, 0x2A, 0xE5,
									0x2D, 0x49, 0x3B, 0x54, 0x80, 0x27, 0x65, 0x10,
									0xE8, 0x2B, 0x19, 0xEB,
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := &app.Transaction{
				ID:   tt.fields.ID,
				Vin:  tt.fields.Vin,
				Vout: tt.fields.Vout,
			}
			tt.args.privKey = *test_privateKey

			tx.Sign(tt.args.privKey, tt.args.prevTXs)
			if got := tx.Verify(tt.args.prevTXs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transaction.Verify() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func TestTransaction_Serialize(t *testing.T) {
// 	type fields struct {
// 		ID   []byte
// 		Vin  []app.TXInput
// 		Vout []app.TXOutput
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		want   []byte
// 	}{
// 		{
// 			name: "cmp",
// 			fields: fields{
// 				ID: []byte{
// 					0x56, 0xE9, 0xC7, 0xE1, 0xC3, 0xDA, 0x4F, 0xCD,
// 					0x6E, 0x04, 0x9F, 0x0D, 0xF0, 0xC4, 0x14, 0x64,
// 					0x39, 0x72, 0x77, 0x02, 0xC5, 0xFC, 0x76, 0x58,
// 					0x2C, 0xCF, 0xFD, 0xEC, 0x12, 0x5B, 0x4A, 0xC8,
// 				},
// 				Vin: []app.TXInput{
// 					{
// 						Txid:      []byte{},
// 						Vout:      -1,
// 						Signature: nil,
// 						PubKey:    []byte("random_data")},
// 				},
// 				Vout: []app.TXOutput{*app.NewTXOutput(subsidy, "1AYmpvb95P8m7je8SZzTnLx7Z6sjRr6PK")},
// 			},
// 			want: []byte{50, 127, 3, 1, 1, 11, 84, 114, 97, 110, 115, 97, 99, 116, 105, 111, 110, 1, 255, 128, 0, 1, 3, 1, 2, 73, 68, 1, 10, 0, 1, 3, 86, 105, 110, 1, 255, 132, 0, 1, 4, 86, 111, 117, 116, 1, 255, 136, 0, 0, 0, 29, 255, 131, 2, 1, 1, 14, 91, 93, 109, 97, 105, 110, 46, 84, 88, 73, 110, 112, 117, 116, 1, 255, 132, 0, 1, 255, 130, 0, 0, 64, 255, 129, 3, 1, 1, 7, 84, 88, 73, 110, 112, 117, 116, 1, 255, 130, 0, 1, 4, 1, 4, 84, 120, 105, 100, 1, 10, 0, 1, 4, 86, 111, 117, 116, 1, 4, 0, 1, 9, 83, 105, 103, 110, 97, 116, 117, 114, 101, 1, 10, 0, 1, 6, 80, 117, 98, 75, 101, 121, 1, 10, 0, 0, 0, 30, 255, 135, 2, 1, 1, 15, 91, 93, 109, 97, 105, 110, 46, 84, 88, 79, 117, 116, 112, 117, 116, 1, 255, 136, 0, 1, 255, 134, 0, 0, 47, 255, 133, 3, 1, 1, 8, 84, 88, 79, 117, 116, 112, 117, 116, 1, 255, 134, 0, 1, 2, 1, 5, 86, 97, 108, 117, 101, 1, 4, 0, 1, 10, 80, 117, 98, 75, 101, 121, 72, 97, 115, 104, 1, 10, 0, 0, 0, 82, 255, 128, 1, 32, 86, 233, 199, 225, 195, 218, 79, 205, 110, 4, 159, 13, 240, 196, 20, 100, 57, 114, 119, 2, 197, 252, 118, 88, 44, 207, 253, 236, 18, 91, 74, 200, 1, 1, 2, 1, 2, 11, 114, 97, 110, 100, 111, 109, 95, 100, 97, 116, 97, 0, 1, 1, 1, 20, 1, 20, 1, 206, 68, 44, 222, 103, 218, 131, 200, 123, 8, 124, 31, 75, 35, 242, 119, 225, 164, 18, 0, 0},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tx := app.Transaction{
// 				ID:   tt.fields.ID,
// 				Vin:  tt.fields.Vin,
// 				Vout: tt.fields.Vout,
// 			}

// 			if got := tx.Serialize(); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Transaction.Serialize() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestDeserializeTransaction(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want app.Transaction
	}{
		{
			name: "cmp",
			args: args{data: []byte{50, 127, 3, 1, 1, 11, 84, 114, 97, 110, 115, 97, 99, 116, 105, 111, 110, 1, 255, 128, 0, 1, 3, 1, 2, 73, 68, 1, 10, 0, 1, 3, 86, 105, 110, 1, 255, 132, 0, 1, 4, 86, 111, 117, 116, 1, 255, 136, 0, 0, 0, 29, 255, 131, 2, 1, 1, 14, 91, 93, 109, 97, 105, 110, 46, 84, 88, 73, 110, 112, 117, 116, 1, 255, 132, 0, 1, 255, 130, 0, 0, 64, 255, 129, 3, 1, 1, 7, 84, 88, 73, 110, 112, 117, 116, 1, 255, 130, 0, 1, 4, 1, 4, 84, 120, 105, 100, 1, 10, 0, 1, 4, 86, 111, 117, 116, 1, 4, 0, 1, 9, 83, 105, 103, 110, 97, 116, 117, 114, 101, 1, 10, 0, 1, 6, 80, 117, 98, 75, 101, 121, 1, 10, 0, 0, 0, 30, 255, 135, 2, 1, 1, 15, 91, 93, 109, 97, 105, 110, 46, 84, 88, 79, 117, 116, 112, 117, 116, 1, 255, 136, 0, 1, 255, 134, 0, 0, 47, 255, 133, 3, 1, 1, 8, 84, 88, 79, 117, 116, 112, 117, 116, 1, 255, 134, 0, 1, 2, 1, 5, 86, 97, 108, 117, 101, 1, 4, 0, 1, 10, 80, 117, 98, 75, 101, 121, 72, 97, 115, 104, 1, 10, 0, 0, 0, 82, 255, 128, 1, 32, 86, 233, 199, 225, 195, 218, 79, 205, 110, 4, 159, 13, 240, 196, 20, 100, 57, 114, 119, 2, 197, 252, 118, 88, 44, 207, 253, 236, 18, 91, 74, 200, 1, 1, 2, 1, 2, 11, 114, 97, 110, 100, 111, 109, 95, 100, 97, 116, 97, 0, 1, 1, 1, 20, 1, 20, 1, 206, 68, 44, 222, 103, 218, 131, 200, 123, 8, 124, 31, 75, 35, 242, 119, 225, 164, 18, 0, 0}},
			want: app.Transaction{
				ID: []byte{
					0x56, 0xE9, 0xC7, 0xE1, 0xC3, 0xDA, 0x4F, 0xCD,
					0x6E, 0x04, 0x9F, 0x0D, 0xF0, 0xC4, 0x14, 0x64,
					0x39, 0x72, 0x77, 0x02, 0xC5, 0xFC, 0x76, 0x58,
					0x2C, 0xCF, 0xFD, 0xEC, 0x12, 0x5B, 0x4A, 0xC8,
				},
				Vin: []app.TXInput{
					{
						Txid:      []byte{},
						Vout:      -1,
						Signature: nil,
						PubKey:    []byte("random_data")},
				},
				Vout: []app.TXOutput{*app.NewTXOutput(subsidy, "1AYmpvb95P8m7je8SZzTnLx7Z6sjRr6PK")},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := app.DeserializeTransaction(tt.args.data)
			if !reflect.DeepEqual(got.ID, tt.want.ID) {
				t.Errorf("DeserializeTransactionID() = %v, want %v", got.ID, tt.want.ID)
			}
			if !reflect.DeepEqual(got.Vin[0].PubKey, tt.want.Vin[0].PubKey) {
				t.Errorf("DeserializeTransactionVin[0]PubKey() = %v, want %v", got.Vin[0].PubKey, tt.want.Vin[0].PubKey)
			}
			if !reflect.DeepEqual(got.Vout[0], tt.want.Vout[0]) {
				t.Errorf("DeserializeTransactionVout() = %v, want %v", got.Vout[0], tt.want.Vout[0])
			}
		})
	}
}
