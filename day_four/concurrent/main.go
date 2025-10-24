package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

// Product represents part of the JSON response structure
type Product struct {
	Title string `json:"title"`
}

// fetchProduct runs as a goroutine to fetch and send product data
func fetchProduct(id int, wg *sync.WaitGroup, resultsChan chan<- string, errorsChan chan<- error) {
	defer wg.Done() // mark goroutine as complete

	url := fmt.Sprintf("https://dummyjson.com/products/%d", id)
	resp, err := http.Get(url)
	if err != nil {
		errorsChan <- fmt.Errorf("product %d: network error: %v", id, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errorsChan <- fmt.Errorf("product %d: received HTTP %d", id, resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		errorsChan <- fmt.Errorf("product %d: read error: %v", id, err)
		return
	}

	var product Product
	if err := json.Unmarshal(body, &product); err != nil {
		errorsChan <- fmt.Errorf("product %d: JSON decode error: %v", id, err)
		return
	}

	resultsChan <- product.Title
}

func main() {
	productIDs := []int{1, 2, 3, 4, 5, 6, 7, 8}

	resultsChan := make(chan string, len(productIDs)) // for successful titles
	errorsChan := make(chan error, len(productIDs))   // for any errors

	var wg sync.WaitGroup

	for _, id := range productIDs {
		wg.Add(1)
		go fetchProduct(id, &wg, resultsChan, errorsChan)
	}

	go func() {
		wg.Wait()
		close(resultsChan)
		close(errorsChan)
	}()

	var titles []string
	var errs []error

	for title := range resultsChan {
		titles = append(titles, title)
	}
	for err := range errorsChan {
		errs = append(errs, err)
	}

	fmt.Println("✅ Product Titles:")
	for _, title := range titles {
		fmt.Println("-", title)
	}

	if len(errs) > 0 {
		fmt.Println("\n Errors:")
		for _, e := range errs {
			fmt.Println("-", e)
		}
	}
}
