package geom

import (
	"testing"
)

func TestDistance(t *testing.T) {
	type args struct {
		X1 float64
		Y1 float64
		X2 float64
		Y2 float64
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name:    "#1",
			args:    args{X1: 1, Y1: 1, X2: 4, Y2: 5},
			want:    5,
			wantErr: false,
		},
		{
			name:    "#2",
			args:    args{X1: -1, Y1: 1, X2: 4, Y2: 5},
			want:    -1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Distance(tt.args.X1, tt.args.Y1, tt.args.X2, tt.args.Y2)
			if (err != nil) != tt.wantErr {
				t.Errorf("Distance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Distance() = %v, want %v", got, tt.want)
			}
		})
	}
}
