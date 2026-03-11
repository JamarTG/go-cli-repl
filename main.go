package main

import (
	"bufio"
	"cli-repl/structs"
	"fmt"
	"os"
	"strings"
)

func cmd2args(words []string, kv *structs.Kv) {

	cmd := words[0]
	key := words[1]
	val := words[2]

	switch cmd {
	case "SET":
		kv.Set(key, val)
	default:
		errFormat := fmt.Errorf("invalid command.")
		fmt.Println(errFormat.Error())
		os.Exit(1)
		// I'll make this return to prmpt
	}

}

func cmd1arg(words []string, kv *structs.Kv) {

	cmd := words[0]
	arg := words[1]

	switch strings.ToUpper(cmd) {

	case "GET":
		// arg represents a key
		kv.Get(arg)
	case "DELETE":
		// arg represents a key
		kv.Get(arg)
	case "COUNT":
		// arg represents a value
		kv.Get(arg)
	default:
		errFormat := fmt.Errorf("invalid command.")
		fmt.Println(errFormat.Error())
		os.Exit(1)
	}

}

func cmdnoargs(words []string, kv *structs.Kv) {
	// cmd := words[0]

	// switch strings.ToUpper(cmd) {
	// case "DELETE":
	// case "ROLLBACK":
	// case "BEGIN":
	// case "END"
	// default
	// }
}

func main() {

	kv := structs.NewKv()

	var stringInput string
	var readErr error

	var input = bufio.NewReader(os.Stdin)

	for stringInput != "END" {
		fmt.Print(" >")
		stringInput, readErr = input.ReadString('\n')
		stringInput = strings.TrimSpace(stringInput)

		if readErr != nil {
			fmt.Fprintf(os.Stderr, "Fatal error: %v\n", readErr)
			os.Exit(1)
		}

		words := strings.Split(stringInput, " ")
		wordCount := len(words)

		switch wordCount {
		case 3:
			if strings.ToUpper(words[0]) != "SET" {
				errFormat := fmt.Errorf("invalid command.Ending...")
				fmt.Println(errFormat.Error())
				os.Exit(122)
			}

			kv.Set(words[1], words[2])
		case 2:

		case 1:
			//
		default:

		}

	}

}

// | Command | Description | Output |
// |---|---|---|
// | `SET <key> <value>` | Sets key to value | Silent |
// | `GET <key>` | Prints the current value | Value or `NULL` |
// | `DELETE <key>` | Removes key if it exists | Silent |
// | `COUNT <value>` | Prints how many keys equal value | Count |
// | `BEGIN` | Starts a new transaction | Silent |
// | `ROLLBACK` | Reverts the most recent transaction | Silent or `NO TRANSACTION` |
// | `COMMIT` | Permanently applies all open transactions | Silent or `NO TRANSACTION` |
// | `END` | Exits the program | Silent |
