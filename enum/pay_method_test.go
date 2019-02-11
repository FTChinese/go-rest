package enum

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestParsePayMethod(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    PayMethod
		wantErr bool
	}{
		{
			name:    "Parse Wechat Pay",
			args:    args{name: "tenpay"},
			want:    PayMethodWx,
			wantErr: false,
		},
		{
			name:    "Parse Alipay",
			args:    args{name: "alipay"},
			want:    PayMethodAli,
			wantErr: false,
		},
		{
			name:    "Parse Stripe",
			args:    args{name: "stripe"},
			want:    PayMethodStripe,
			wantErr: false,
		},
		{
			name:    "Parse Null",
			args:    args{name: `null`},
			want:    InvalidPay,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParsePayMethod(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParsePayMethod() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParsePayMethod() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPayMethod_UnmarshalJSON(t *testing.T) {
	type fields struct {
		PayMethod PayMethod `json:"payMethod"`
	}
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		args    args
		want    fields
		wantErr bool
	}{
		{
			name: "Unmarshal Alipay",
			args: args{
				b: []byte(`{"payMethod": "alipay"}`),
			},
			want: fields{
				PayMethod: PayMethodAli,
			},
			wantErr: false,
		},
		{
			name: "Unmarshal Tenpay",
			args: args{
				b: []byte(`{"payMethod": "tenpay"}`),
			},
			want: fields{
				PayMethod: PayMethodWx,
			},
			wantErr: false,
		},
		{
			name: "Unmarshal Null",
			args: args{
				b: []byte(`{"payMethod": null}`),
			},
			want: fields{
				PayMethod: InvalidPay,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got fields

			if err := json.Unmarshal(tt.args.b, &got); (err != nil) != tt.wantErr {
				t.Errorf("PayMethod.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PayMethod.UnmarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPayMethod_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		x       PayMethod
		want    []byte
		wantErr bool
	}{
		{
			name:    "Marshal Alipay",
			x:       PayMethodAli,
			want:    []byte(`"alipay"`),
			wantErr: false,
		},
		{
			name:    "Marshal Tenpay",
			x:       PayMethodWx,
			want:    []byte(`"tenpay"`),
			wantErr: false,
		},
		{
			name:    "Marshal Null",
			x:       InvalidPay,
			want:    []byte(`null`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.x.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("PayMethod.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PayMethod.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
