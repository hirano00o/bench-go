package malloc

import (
	"testing"
)

type args struct {
	count int
}

var tests = []struct {
	name string
	args args
}{
	{
		name: "Try 10 times",
		args: args{
			count: 10,
		},
	},
	{
		name: "Try 100 times",
		args: args{
			count: 100,
		},
	},
	{
		name: "Try 1,000 times",
		args: args{
			count: 1000,
		},
	},
	{
		name: "Try 100,000 times",
		args: args{
			count: 100000,
		},
	},
	{
		name: "Try 10,000,000 times",
		args: args{
			count: 10000000,
		},
	},
}

func BenchmarkEveryStringAllocation(b *testing.B) {
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			b.ResetTimer()
			c := 0
			for i := 0; i < b.N; i++ {
				got := everyStringAllocation(tt.args.count)
				c += got
			}
			b.Logf("result of %d times: cap avg = %d", tt.args.count, c/b.N)
		})
	}
}

func BenchmarkOnceStringAllocation(b *testing.B) {
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			b.ResetTimer()
			c := 0
			for i := 0; i < b.N; i++ {
				got := onceStringAllocation(tt.args.count)
				c += got
			}
			b.Logf("result of %d times: cap avg = %d", tt.args.count, c/b.N)
		})
	}
}

func BenchmarkEveryStructAllocation(b *testing.B) {
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			b.ResetTimer()
			c := 0
			for i := 0; i < b.N; i++ {
				got := everyStructAllocation(tt.args.count)
				c += got
			}
			b.Logf("result of %d times: cap avg = %d", tt.args.count, c/b.N)
		})
	}
}

func BenchmarkOnceStructAllocation(b *testing.B) {
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			b.ResetTimer()
			c := 0
			for i := 0; i < b.N; i++ {
				got := onceStructAllocation(tt.args.count)
				c += got
			}
			b.Logf("result of %d times: cap avg = %d", tt.args.count, c/b.N)
		})
	}
}
