package main

import (
	"fmt"

	"github.com/equinor/radix-common/utils/slice"
	// "github.com/urfave/negroni/v3"
	// secretproviderfake "sigs.k8s.io/secrets-store-csi-driver/pkg/client/clientset/versioned/fake"
)

func main() {
	// negroni.New()
	fmt.Println("done")
	slice.All(nil, func(s string) bool { return true })
	// secretproviderfake.NewSimpleClientset()
}
