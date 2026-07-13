package logproc

import (
	"fmt"
	"io/fs" 
	"path/filepath"
	"sync"
)

var wg sync.WaitGroup

func RunEngine(dirPath string, target string, workers int) (int, error) {
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

		go func() {
			defer wg.Done()

			for path := range jobs { 
				fmt.Printf("Worker in path: %s\n", path) 
			}
		}() 
	}

	return 0, nil
}
