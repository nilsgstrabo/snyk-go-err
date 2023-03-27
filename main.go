package main

import (
	"fmt"

	"github.com/urfave/negroni/v2"
)

func main() {
	negroni.New()
	fmt.Println("done")
}
