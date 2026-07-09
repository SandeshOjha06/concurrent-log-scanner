package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/SandeshOjha06/go-systems-practice"
)

func main(){
	
	dirPath := flag.String("dir", "./logs", "dirrctory containing log files")

	target := flag.String("word", "ERROR", "Counter of specific words")

	workers := flag.String("workers", 4, "Number of concurrent goroutinrs")

	flag.Parse()

	if *dirPath == "" || *target == "" {
		log.Fatal("Both -dir and -word flags are required")
	}

	startTime := time.Now()

	totalCount, err := logproc.RunEngine(*dirPath, *target, *workers) 

	if err != nil {
		log.Fatal("Execution failed: %v", err)
	}



}
