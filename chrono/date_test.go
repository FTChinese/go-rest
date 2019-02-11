package chrono

import (
	"reflect"
	"testing"
	"time"
)

func TestDate_MarshalJSON(t *testing.T) {
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
			name:    "Marshal Today",
			fields:  fields{Time: now},
			want:    []byte(`"` + now.In(time.UTC).Format(SQLDate) + `"`),
			wantErr: false,
		},
		{
			name:    "Marshal Zero",
			fields:  fields{Time: time.Time{}},
			want:    []byte(`null`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Date{
				Time: tt.fields.Time,
			}
			got, err := d.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Date.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Date.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
