package main

import (
	"github.com/Chanokthorn/go-public-lib-test-client/duck"
	"github.com/Chanokthorn/go-public-lib-test/bar"
	"github.com/Chanokthorn/go-public-lib-test/foo"
	"github.com/Chanokthorn/go-public-lib-test/quack"
)

func main() {
	quack.Set(bar.NewQuackService("john"))

	duckService := duck.NewDuckService()

	err := duckService.Quack()
	if err != nil {
		panic(err)
	}

	quack.Set(foo.NewQuackService("jane", 24))

	err = duckService.Quack()
	if err != nil {
		panic(err)
	}

}
