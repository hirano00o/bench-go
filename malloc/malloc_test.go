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
			for i := 0; i < b.N; i++ {
				everyStringAllocation(tt.args.count)
			}
		})
	}
}

func BenchmarkOnceStringAllocation(b *testing.B) {
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				onceStringAllocation(tt.args.count)
			}
		})
	}
}

func BenchmarkEveryStructAllocation(b *testing.B) {
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				everyStructAllocation(tt.args.count)
			}
		})
	}
}

func BenchmarkOnceStructAllocation(b *testing.B) {
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				onceStructAllocation(tt.args.count)
			}
		})
	}
}
