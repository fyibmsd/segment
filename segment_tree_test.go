package segment

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestSegmentTree(t *testing.T) {
	numbers := []int{1, 3, 5, 7, 9, 11}

	st := NewSegmentTree(numbers)

	type args struct {
		begin int
		end   int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "test query sum after init",
			args:    args{1, 4},
			want:    24,
			wantErr: false,
		},
		{
			name:    "test query sum after update",
			args:    args{1, 4},
			want:    21,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := st.QuerySum(tt.args.begin, tt.args.end)

			if (err != nil) != tt.wantErr {
				t.Errorf("QuerySum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("QuerySum() got = %v, want %v", got, tt.want)
			}

			if err = st.Update(4, 6); err != nil {
				t.Errorf("Update() error = %v", err)
			}
		})
	}
}

func BenchmarkQuerySum(b *testing.B) {
	// setup
	var data []int
	for i := 0; i < 1000; i++ {
		data = append(data, int(rand.Int31n(100000)))
	}

	st := NewSegmentTree(data)

	begin, end := 345, 678

	b.Run("query sum by segment tree", func(b *testing.B) {
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			_, _ = st.QuerySum(begin, end)
		}
	})

	b.Run("query sum by range", func(b *testing.B) {
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			sum := 0
			for i := begin; i <= end; i++ {
				sum += data[i]
			}
		}
	})
}

func ExampleSegmentTree() {
	numbers := []int{1, 3, 5, 7, 9, 11}

	st := NewSegmentTree(numbers)

	before, _ := st.QuerySum(1, 4)

	// set numbers[4] = 6
	_ = st.Update(4, 6)

	after, _ := st.QuerySum(1, 4)

	fmt.Printf("before = %d, after = %d\n", before, after)
	// Output: before = 24, after = 21
}
