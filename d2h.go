package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/big"
	"os"
)

func main() {
	// Define flags
	encode := flag.Bool("e", false, "Encode from decimal to hexadecimal")
	decode := flag.Bool("d", false, "Decode from hexadecimal to decimal")
	flag.Parse()

	// Check whether exactly one flag is set
	if (*encode && *decode) || (!*encode && !*decode) {
		fmt.Println("Please enter either -e for encoding or -d for decoding.")
		os.Exit(1)
	}

	// Read input from stdin
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	// Create a new big.Int instance
	num := new(big.Int)

	if *decode {
		// Decode: Hexadecimal to decimal
		num.SetString(input, 16)
		fmt.Println(num.String())
	} else if *encode {
		// Encode: Decimal to hexadecimal
		num.SetString(input, 10)
		fmt.Println(num.Text(16))
	}
}

