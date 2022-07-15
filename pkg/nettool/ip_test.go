package nettool

import (
	"reflect"
	"testing"
)

func TestGetAllIPs(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    []string
		wantErr bool
	}{
		{
			name:    "valied cidr inpuy",
			input:   "192.168.0.0/30",
			want:    []string{"192.168.0.1/30", "192.168.0.2/30"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			allips, err := GetAllIPs(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllIPs error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(allips, tt.want) {
				t.Errorf("wanted:\n%v\ngot:\n%v", tt.want, allips)
			}
		})
	}
}
