package main

// import edilen paketler yazılır.

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/berkmenguc/goscanner/internal/scanner"
)

var (
	host       = flag.String("host", "scanme.nmap.org", "Target host to scan")
	timeout    = flag.Duration("timeout", 1*time.Second, "Connection timeout")
	startPort  = flag.Int("start", 1, "Start port number")
	endPort    = flag.Int("end", 1024, "End port number")
	outputFile = flag.String("output", "", "Output file to save results (optional)")
)

func main() {
	flag.Parse()
	// Tarama yapılacak hedef IP veya domain
	// Bağlantı zaman aşımı süresi
	slice := make([]scanner.Scanner, 0)
	result := make(chan scanner.Scanner, 1024)

	var wg sync.WaitGroup

	for port := *startPort; port <= *endPort; port++ {
		wg.Add(1)
		go func(p int) {
			s := scanner.ScanPort(*host, p, *timeout)
			result <- s
			wg.Done()
		}(port)
	}
	go func() {
		wg.Wait()
		close(result)
	}()

	for s := range result {
		if s.Open {
			fmt.Printf("Port %d is open\n", s.Port)
			slice = append(slice, s)

		}
		if s.Banner != "" {
			fmt.Printf("Banner for port %d: %s\n", s.Port, s.Banner)
		}
	}

	if *outputFile != "" {
		file, err := os.Create(*outputFile)
		if err != nil {
			fmt.Printf("Error creating output file: %v\n", err)
			return
		}
		defer file.Close()
		enc := json.NewEncoder(file)
		enc.SetIndent("", "  ")

		enc.Encode(slice)
	}

}
