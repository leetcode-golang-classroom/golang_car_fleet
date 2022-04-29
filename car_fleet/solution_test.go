package car_fleet

import "testing"

func Test_carFleet(t *testing.T) {
	type args struct {
		target   int
		position []int
		speed    []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "target = 12, position = [10,8,0,5,3], speed = [2,4,1,1,3]",
			args: args{target: 12, position: []int{10, 8, 0, 5, 3}, speed: []int{2, 4, 1, 1, 3}},
			want: 3,
		},
		{
			name: "target = 10, position = [3], speed = [3]",
			args: args{target: 10, position: []int{3}, speed: []int{3}},
			want: 1,
		},
		{
			name: "target = 100, position = [0,2,4], speed = [4,2,1]",
			args: args{target: 100, position: []int{0, 2, 4}, speed: []int{4, 2, 1}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := carFleet(tt.args.target, tt.args.position, tt.args.speed); got != tt.want {
				t.Errorf("carFleet() = %v, want %v", got, tt.want)
			}
		})
	}
}
