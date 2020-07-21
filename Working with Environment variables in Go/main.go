package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Reading Environment Variable\n")
	var databasePass string
	databasePass = os.Getenv("DATABASE_PASS")
	fmt.Printf("Database secret is:%s\n\n", databasePass)

	err := os.Setenv("DATABASE_PASS", "PennyKnowsAboutIt")
	if err != nil {
		fmt.Println(err)
	}
	databasePass = os.Getenv("DATABASE_PASS")
	fmt.Printf("Database new secret is:%s\n\n", databasePass)
}
