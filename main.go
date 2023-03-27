package main

import (
	"fmt"

	"github.com/urfave/negroni/v3"
)

func main() {
	negroni.New()
	fmt.Println("done")
}
