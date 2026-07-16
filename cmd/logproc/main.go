package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/SandeshOjha06/go-systems-practice/internal/logproc"
)

func main() {
	dirPath := flag.String("dir", "./logs", "directory containing log files")
	target := flag.String("word", "ERROR", "Counter of specific words")
	workers := flag.Int("workers", 4, "Number of concurrent goroutines")

	flag.Parse()

	if *dirPath == "" || *target == "" {
		log.Fatalf("Both -dir and -word flags are required")
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	totalCount, err := logproc.RunEngine(ctx, *dirPath, *target, *workers)

	if err != nil {
		// 2. FIXED: Changed Fatal to Fatalf
		log.Fatalf("Execution failed: %v", err)
	}

	fmt.Printf("Scan complete... Total occurences found: %d\n", totalCount)
}
