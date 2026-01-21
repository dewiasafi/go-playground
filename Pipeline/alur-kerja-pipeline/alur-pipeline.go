package main

import (
	"context"
	"fmt"
	"time"
)

func AlurPipeline() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := generate(ctx, 1, 2, 3, 4)
	result := square(ctx, c)
	start := time.Now()

	for v := range result {
		fmt.Println("Alur pipeline:", v)
		if v == 9 {
			fmt.Println("Oke stop disini")
			cancel() // stop semua pipeline
		}
	}
	duration := time.Since(start)
	fmt.Println("done in", duration.Seconds(), "seconds")
}

func generate(ctx context.Context, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case <-ctx.Done():
				return
			case out <- n:
			}
		}
	}()
	return out
}

func square(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return
			case n, ok := <-in:
				fmt.Println(n, ok)
				if !ok {
					return
				}
				select {
				// kondisi ketika pipeline harus dihentikan paksa
				case <-ctx.Done():
					return
				case out <- n * n:
				}
			}
		}
	}()
	return out
}
