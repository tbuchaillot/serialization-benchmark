package analytics_bebop

//
/*
package main

import (
	"fmt"
	"os"

	"github.com/200sc/bebop"
)

func main() {
	f, _ := os.Open("analytics.bop")
	defer f.Close()
	bopf, _, err := bebop.ReadFile(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	out, _ := os.Create("analytics.go")
	defer out.Close()
	settings := bebop.GenerateSettings{
		PackageName: "analytics_bebop",
	}
	bopf.Generate(out, settings)
}
*/
