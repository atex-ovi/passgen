package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	length := flag.Int("l", 12, "Password length")
	special := flag.Bool("s", false, "Include special characters")
	numbers := flag.Bool("n", true, "Include numbers")
	help := flag.Bool("h", false, "Show help")
	flag.Parse()

	if *help {
		fmt.Println("Password Generator - Simple CLI Tool")
		fmt.Println("\nUsage:")
		fmt.Println("  passgen -l 16 -s      Generate 16 char with special chars")
		fmt.Println("  passgen -l 8          Generate 8 char (default: 12)")
		fmt.Println("  passgen -n false      Generate without numbers")
		fmt.Println("\nOptions:")
		flag.PrintDefaults()
		return
	}

	rand.Seed(time.Now().UnixNano())
	lowercase := "abcdefghijklmnopqrstuvwxyz"
	uppercase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digit := "0123456789"
	specialChars := "!@#$%^&*()_+-=[]{}"
	pool := lowercase + uppercase
	if *numbers {
		pool += digit
	}
	if *special {
		pool += specialChars
	}

	if len(pool) == 0 {
		fmt.Println("Error: No character sets selected")
		return
	}

	password := make([]byte, *length)
	for i := range password {
		password[i] = pool[rand.Intn(len(pool))]
	}

	fmt.Println("\n🔐 Generated Password:")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Printf("  %s\n", string(password))
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Printf("  Length: %d | Special: %v | Numbers: %v\n", *length, *special, *numbers)
}
