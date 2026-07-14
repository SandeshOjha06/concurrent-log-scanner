package logproc

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
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
			

			for path := range jobs { 
				fmt.Printf("Worker %d in path: %s\n", workerId, path) 
			}
		}() 
	}

	wg.Wait()
	fmt.Println("All workers have finished")

	return 0, nil
}
