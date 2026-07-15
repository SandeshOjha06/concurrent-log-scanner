package logproc

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func processFile(path, target string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}

	defer file.Close()

	count := 0 

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		count += strings.Count(line, target)
	}
		
	if err := scanner.Err(); err != nil {
		return count, err
	}

	return count, nil

}

func RunEngine(dirPath string, target string, workers int) (int, error) {
	var wg sync.WaitGroup
	jobs := make(chan string, 100)
	results := make(chan int, 100)
	// take file info
	go func() {
		err := filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {
			if err != nil { 
				return err
			} 

			if !d.IsDir() {
				jobs <- path
			}

			return nil
		}) 

		if err != nil {
			fmt.Printf("Error working directory: %v\n", err)
		}

		close(jobs)
	}()

	for i := 0; i < workers; i++ {
		wg.Add(1)

		go func(workerId int) {
			defer wg.Done()
		//process file


			for path := range jobs { 
				count, err := processFile(path, target)
				if err != nil {
					fmt.Print("Worker failed")
					continue
				}

				results <- count
			}
		}(i) 
	}

	// this prevents deadlock
	// results channel would be forever filled otherwise
	go func() {
		wg.Wait()
		close(results)
	}()

	totalCount := 0

	for count := range results {
		totalCount += count
	}

	return totalCount, nil
}
