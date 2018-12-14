package idcard

import "testing"

func TestGetAge(t *testing.T) {
	type args struct {
		idNo string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"normal", args{"400000199112134569"}, 27},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAge(tt.args.idNo); got != tt.want {
				t.Errorf("GetAge() = %v, want %v", got, tt.want)
			}
		})
	}
}
