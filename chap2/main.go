package main

import (
	"log"
	"os"

	_ "./matchers"
	"./search"
	"fmt"
)

func init()  {
	log.SetOutput(os.Stdout)
}

func main()  {
	fmt.Println(os.Getenv("GOPATH"))
	search.Run("president")
}
