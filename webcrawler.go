package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"
)

var (
	// Regular expression to extract links from HTML
	linkRegex = regexp.MustCompile(`href=["'](http[s]?://[^"']+)["']`)
	// Mutex to ensure thread-safe writes to the visited map
	mutex = &sync.Mutex{}
	// Map to track visited URLs
	visited = make(map[string]bool)
	// WaitGroup to manage goroutines
	wg sync.WaitGroup
)

func main() {
	// Starting URL
	startURL := "https://google.com"

	// Maximum depth to crawl
	maxDepth := 2

	fmt.Printf("Starting web crawl at %s with a depth of %d\n", startURL, maxDepth)
	crawl(startURL, maxDepth)
}

// crawl is the main function for crawling a URL recursively
func crawl(url string, depth int) {
	// Use a buffered channel to limit concurrency
	concurrency := 5
	semaphore := make(chan struct{}, concurrency)

	// Launch the crawling process
	wg.Add(1)
	go func() {
		defer wg.Done()
		crawlHelper(url, depth, semaphore)
	}()

	// Wait for all goroutines to finish
	wg.Wait()
	fmt.Println("Crawling completed.")
}

// crawlHelper handles individual URL crawling
func crawlHelper(url string, depth int, semaphore chan struct{}) {
	if depth == 0 {
		return
	}

	// Check if the URL has already been visited
	mutex.Lock()
	if visited[url] {
		mutex.Unlock()
		return
	}
	visited[url] = true
	mutex.Unlock()

	// Acquire a slot in the semaphore
	semaphore <- struct{}{}
	defer func() { <-semaphore }()

	fmt.Printf("Crawling: %s (Depth: %d)\n", url, depth)

	// Fetch the page content
	body, err := fetchURL(url)
	if err != nil {
		fmt.Printf("Failed to fetch URL: %s, Error: %s\n", url, err)
		return
	}

	// Extract links from the page
	links := extractLinks(body)

	// Recursively crawl the discovered links
	for _, link := range links {
		wg.Add(1)
		go func(link string) {
			defer wg.Done()
			crawlHelper(link, depth-1, semaphore)
		}(link)
	}
}

// fetchURL fetches the HTML content of a URL
func fetchURL(url string) (string, error) {
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP error: %d", resp.StatusCode)
	}

	// Read the response body using io.Copy
	var buf strings.Builder
	_, err = io.Copy(&buf, resp.Body)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

// extractLinks parses the HTML and returns all found links
func extractLinks(html string) []string {
	matches := linkRegex.FindAllStringSubmatch(html, -1)
	var links []string
	for _, match := range matches {
		if len(match) > 1 {
			links = append(links, match[1])
		}
	}
	return links
}
