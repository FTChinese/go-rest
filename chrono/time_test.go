package chrono

import (
	"database/sql/driver"
	"reflect"
	"testing"
	"time"
)

func TestTime_StringEN(t *testing.T) {
	type fields struct {
		Time time.Time
	}
	now := time.Now()

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "English text",
			fields: fields{Time: now},
			want:   now.In(time.UTC).Format(time.RFC1123Z),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tm := Time{
				Time: tt.fields.Time,
			}
			if got := tm.StringEN(); got != tt.want {
				t.Errorf("Time.StringEN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTime_StringCN(t *testing.T) {
	type fields struct {
		Time time.Time
	}
	now := time.Now()
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Chinese text",
			fields: fields{Time: now},
			want:   now.In(TZShanghai).Format(CST),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tm := Time{
				Time: tt.fields.Time,
			}
			if got := tm.StringCN(); got != tt.want {
				t.Errorf("Time.StringCN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTime_MarshalJSON(t *testing.T) {
	type fields struct {
		Time time.Time
	}
	now := time.Now()

	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name:    "Stringify",
			fields:  fields{Time: now},
			want:    []byte(`"` + now.In(time.UTC).Format(time.RFC3339) + `"`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tm := Time{
				Time: tt.fields.Time,
			}
			got, err := tm.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Time.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("Time.MarshalJSON() got %s", got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Time.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTime_UnmarshalJSON(t *testing.T) {
	type fields struct {
		Time time.Time
	}
	type args struct {
		data []byte
	}
	now := time.Now()
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "Parse",
			fields:  fields{Time: time.Time{}},
			args:    args{data: []byte(`"` + now.In(time.UTC).Format(time.RFC3339) + `"`)},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tm := &Time{
				Time: tt.fields.Time,
			}
			if err := tm.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Time.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			t.Logf("Time.UnmarshalJSON() %v", tm.Time)
		})
	}
}

func TestTime_Scan(t *testing.T) {
	type fields struct {
		Time time.Time
	}
	type args struct {
		value interface{}
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "Scan",
			fields:  fields{Time: time.Time{}},
			args:    args{value: time.Now().In(time.UTC).Format(SQLDateTime)},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tm := &Time{
				Time: tt.fields.Time,
			}
			if err := tm.Scan(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Time.Scan() error = %v, wantErr %v", err, tt.wantErr)
			}
			t.Logf("Time.Scan() %v", tm)
		})
	}
}

func TestTime_Value(t *testing.T) {
	type fields struct {
		Time time.Time
	}

	now := time.Now()

	tests := []struct {
		name    string
		fields  fields
		want    driver.Value
		wantErr bool
	}{
		{
			name:    "Value",
			fields:  fields{Time: now},
			want:    now.In(time.UTC).Format(SQLDateTime),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tm := Time{
				Time: tt.fields.Time,
			}
			got, err := tm.Value()
			if (err != nil) != tt.wantErr {
				t.Errorf("Time.Value() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("Time.Value() got %s", got)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Time.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}
