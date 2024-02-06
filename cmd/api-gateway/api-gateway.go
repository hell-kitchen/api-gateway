package main

import "go.uber.org/fx"

func main() {
	fx.New(
		NewOptions(),
	).Run()
}

func NewOptions() fx.Option {
	return fx.Options()
}
