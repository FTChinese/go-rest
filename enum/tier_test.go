package enum

import (
	"database/sql/driver"
	"encoding/json"
	"reflect"
	"testing"
)

func TestTierMarshal(t *testing.T) {
	data := struct {
		Invalid  Tier
		Standard Tier
		Premium  Tier
	}{
		Invalid:  InvalidTier,
		Standard: TierStandard,
		Premium:  TierPremium,
	}

	b, err := json.Marshal(data)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%s", b)
}

func TestParseTier(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    Tier
		wantErr bool
	}{
		{
			name:    "Test Standard Tier",
			args:    args{"standard"},
			want:    TierStandard,
			wantErr: false,
		},
		{
			name:    "Test Premium Tier",
			args:    args{"premium"},
			want:    TierPremium,
			wantErr: false,
		},
		{
			name:    "Test Unknown Tier",
			args:    args{"advanced"},
			want:    InvalidTier,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseTier(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseTier() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseTier() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTier_String(t *testing.T) {
	tests := []struct {
		name string
		x    Tier
		want string
	}{
		{
			name: "Standard Tier",
			x:    TierStandard,
			want: "standard",
		},
		{
			name: "Premium Tier",
			x:    TierPremium,
			want: "premium",
		},
		{
			name: "Invalid Tier",
			x:    InvalidTier,
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.x.String(); got != tt.want {
				t.Errorf("Tier.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTier_StringCN(t *testing.T) {
	tests := []struct {
		name string
		x    Tier
		want string
	}{
		{
			name: "Standard Tier",
			x:    TierStandard,
			want: "标准会员",
		},
		{
			name: "Premium Tier",
			x:    TierPremium,
			want: "高级会员",
		},
		{
			name: "Invalid Tier",
			x:    InvalidTier,
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.x.StringCN(); got != tt.want {
				t.Errorf("Tier.StringCN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTier_StringEN(t *testing.T) {
	tests := []struct {
		name string
		x    Tier
		want string
	}{
		{
			name: "Standard Tier",
			x:    TierStandard,
			want: "Standard",
		},
		{
			name: "Premium Tier",
			x:    TierPremium,
			want: "Premium",
		},
		{
			name: "Invalid Tier",
			x:    InvalidTier,
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.x.StringEN(); got != tt.want {
				t.Errorf("Tier.StringEN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTier_UnmarshalJSON(t *testing.T) {
	type args struct {
		b []byte
	}
	var tier Tier
	tests := []struct {
		name    string
		x       *Tier
		args    args
		wantErr bool
	}{
		{
			name:    "Standard Tier",
			x:       &tier,
			args:    args{[]byte(`"standard"`)},
			wantErr: false,
		},
		{
			name:    "Premium Tier",
			x:       &tier,
			args:    args{[]byte(`"premium"`)},
			wantErr: false,
		},
		{
			name:    "Invalid Tier",
			x:       &tier,
			args:    args{[]byte(`"invalid"`)},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.x.UnmarshalJSON(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("Tier.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTier_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		x       Tier
		want    []byte
		wantErr bool
	}{
		{
			name:    "Standard",
			x:       TierStandard,
			want:    []byte(`"standard"`),
			wantErr: false,
		},
		{
			name:    "Premium",
			x:       TierPremium,
			want:    []byte(`"premium"`),
			wantErr: false,
		},
		{
			name:    "Invalid",
			x:       InvalidTier,
			want:    []byte("null"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.x.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Tier.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tier.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTier_Scan(t *testing.T) {
	type args struct {
		src interface{}
	}
	var tier Tier
	tests := []struct {
		name    string
		x       *Tier
		args    args
		wantErr bool
	}{
		{
			name:    "Standard",
			x:       &tier,
			args:    args{[]byte("standard")},
			wantErr: false,
		},
		{
			name:    "Premium",
			x:       &tier,
			args:    args{[]byte("premium")},
			wantErr: false,
		},
		{
			name:    "NULL",
			x:       &tier,
			args:    args{nil},
			wantErr: false,
		},
		{
			name:    "Invalid",
			x:       &tier,
			args:    args{[]byte("invalid")},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.x.Scan(tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("Tier.Scan() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTier_Value(t *testing.T) {
	tests := []struct {
		name    string
		x       Tier
		want    driver.Value
		wantErr bool
	}{
		{
			name:    "Standard",
			x:       TierStandard,
			want:    "standard",
			wantErr: false,
		},
		{
			name:    "Premium",
			x:       TierPremium,
			want:    "premium",
			wantErr: false,
		},
		{
			name:    "null",
			x:       InvalidTier,
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.x.Value()
			if (err != nil) != tt.wantErr {
				t.Errorf("Tier.Value() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tier.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}
