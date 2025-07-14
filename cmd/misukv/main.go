package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/misudb/misudb/storage/kv"
)

func main() {
	kvStore := kv.New()
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to MisuKV! Type 'exit' to quit.")
	fmt.Print("> ")

	for scanner.Scan() {
		input := scanner.Text()
		if input == "exit" || input == "quit" {
			break
		}

		parts := strings.SplitN(input, " ", 3)
		if len(parts) < 2 {
			fmt.Println("Invalid command. Use 'set <key> <value>' or 'get <key>' or 'delete <key>'.")
			continue
		}
		command := parts[0]
		key := parts[1]
		switch command {
		case "set":
			if len(parts) != 3 {
				fmt.Println("Invalid command. Use 'set <key> <value>' or 'get <key>' or 'delete <key>'.")
				continue
			}
			value := parts[2]
			kvStore.Set(key, value)
			fmt.Printf("Set %s = %s\n", key, value)
		case "get":
			value, ok := kvStore.Get(key)
			if ok {
				fmt.Printf("Get %s = %s\n", key, value)
			} else {
				fmt.Printf("Key `%s` not found\n", key)
			}
		case "delete":
			kvStore.Delete(key)
			fmt.Printf("Deleted `%s`\n", key)
		default:
			fmt.Println("Invalid command. Use 'set <key> <value>' or 'get <key>' or 'delete <key>'.")
		}
		fmt.Print("> ")
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}
}
