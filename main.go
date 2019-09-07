package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	autoruns "github.com/steve-offutt/autoruns-csv/autoruns"
)

func main() {
	if len(os.Args) != 2 {
		exeName := filepath.Base(os.Args[0])
		fmt.Println("Incorrect program usage.")
		fmt.Printf("\tUsage: %s CSV_INPUT\n", exeName)
		os.Exit(1)
	}
	auto := autoruns.NewAutoRuns(os.Args[1])
	for _, a := range auto {
		jSlice, err := json.Marshal(a)
		if err != nil {
			log.Panicln(err)
		}
		fmt.Println(string(jSlice))
	}
}
