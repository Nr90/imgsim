package imgsim

import "testing"

func TestDistance(t *testing.T) {
	type args struct {
		a Hash
		b Hash
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"same hash", args{18446743590535167999, 18446743590535167999}, 0},
		{"different hash", args{0, 18446743590535167999}, 51},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Distance(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Distance() = %v, want %v", got, tt.want)
			}
		})
	}
}
