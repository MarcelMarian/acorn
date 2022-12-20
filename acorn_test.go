package main

import (
	"context"
	"testing"
	"time"
)

func Test_jiggling(t *testing.T) {
	type args struct {
		ctx      context.Context
		interval time.Duration
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	tests := []struct {
		name string
		args args
	}{
		{
			name: "normal",
			args: args{
				ctx:      ctx,
				interval: time.Second,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jiggling(tt.args.ctx, tt.args.interval)
		})
	}
}
