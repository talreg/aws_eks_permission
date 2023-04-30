package main

import (
	"com.talreg.awstests/programs"
	"fmt"
	"log"
	"time"
)

func main() {
	program := programs.NewProgram()
	for {
		if err := program.Initialize(); err != nil {
			log.Print("failed to initialize", err)
		} else {
			log.Println("initialized")
			if data, err := program.Execute(); err != nil {
				log.Print("error executing:", err)
			} else {
				log.Println("retrieval was successful:")
				time.Sleep(10 * time.Millisecond)
				fmt.Println()
				fmt.Print(data)
				fmt.Println()
			}
		}
		time.Sleep(100 * time.Millisecond)
		log.Println("cycle completed, waiting 20 seconds...")
		fmt.Println()
		fmt.Println()
		time.Sleep(20 * time.Second)
	}

}
