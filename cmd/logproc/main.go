package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/SandeshOjha06/go-systems-practice/internal/logproc"
)

func main(){
	
	dirPath := flag.String("dir", "./logs", "dirrctory containing log files")

	target := flag.String("word", "ERROR", "Counter of specific words")

	workers := flag.Int("workers", 4, "Number of concurrent goroutinrs")

	flag.Parse()

	if *dirPath == "" || *target == "" {
		log.Fatalf("Both -dir and -word flags are required")
	}

	

	totalCount, err := logproc.RunEngine(*dirPath, *target, *workers) 

	if err != nil {
		log.Fatal("Execution failed: %v", err)
	}

	fmt.Printf("Scan complete... Total occurences found: %d\n", totalCount)

}
