package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"
)

type myData struct {
	Name string
	ID   int
}

func main() {

	filePath := "./myData.json"
	workers := 2
	var wg sync.WaitGroup

	myChan := make(chan myData)

	// spawn json reader
	wg.Add(1)
	go func() {
		err := reader(myChan, filePath)
		if err != nil {
			fmt.Printf("reader error: %s", err)
		}
		wg.Done()
	}()

	for id := 1; id <= workers; id++ {
		wg.Add(1)
		go func(id int) {
			worker(myChan, id)
			wg.Done()
		}(id)
	}

	wg.Wait()

}

func reader(myChan chan myData, fp string) error {
	fmt.Println("reader starting")
	defer fmt.Println("reader stopping")

	var allJson []myData

	b, err := os.ReadFile(fp)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, &allJson)
	if err != nil {
		return err
	}

	for _, json := range allJson {
		myChan <- json
	}
	close(myChan)

	return nil
}

func worker(myChan chan myData, workerID int) {
	fmt.Printf("worker %d starting\n", workerID)
	defer fmt.Printf("worker %d exiting\n", workerID)
	for data := range myChan {
		fmt.Printf("worker %d: processing\n\tName: %s\n\tID: %d\n", workerID, data.Name, data.ID)
		time.Sleep(1)
	}

}
