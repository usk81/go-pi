package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/usk81/go-pi/calc"

	"github.com/spf13/cobra"
)

var (
	// RootCmd sets task command config
	RootCmd = &cobra.Command{
		Use: "gopi",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 3 {
				Exit(fmt.Errorf("arguments are not enough"), 1)
			}
			h, err := strconv.ParseFloat(args[0], 64)
			if err != nil {
				Exit(err, 1)
			}
			w, err := strconv.ParseFloat(args[1], 64)
			if err != nil {
				Exit(err, 1)
			}
			age, err := strconv.ParseFloat(args[2], 64)
			if err != nil {
				Exit(err, 1)
			}
			r, err := calc.Run(w, h, age)
			if err != nil {
				Exit(err, 1)
			}
			fmt.Printf("Classification: %s\nIndex: %f\nIndexType: %s\nStatus: %s\n", r.Classification, r.Index, r.IndexType, r.Status)
		},
	}
)

func main() {
	if err := RootCmd.Execute(); err != nil {
		Exit(fmt.Errorf("failed to run: %s", err.Error()), 1)
	}
}

// Exit finishs requests
func Exit(err error, codes ...int) {
	var code int
	if len(codes) > 0 {
		code = codes[0]
	} else {
		code = 2
	}
	fmt.Println(err)
	os.Exit(code)
}
