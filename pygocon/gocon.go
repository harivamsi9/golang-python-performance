package main

/*
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"net/http"
	"sync"
	"time"
	"unsafe"
)

//export bar
func bar(a string, b *C.char) *C.char {
	c := a
	d := C.GoString(b) // Convert C string to Go string
	go func() {
		for {
			// Just printing for demonstration
			fmt.Println("a:", a)
			fmt.Println("b:", d) // Correctly used Go string from C char pointer
			fmt.Println("c:", c)
			fmt.Println("d:", d)
		}
	}()
	return C.CString("hello from Go") // Allocate new C string
}

//export freeme
func freeme(addr *C.char) {
	C.free(unsafe.Pointer(addr)) // Free memory allocated by C.CString
}

func makeRequest(url *C.char) {
	// Convert C string to Go string
	goURL := C.GoString(url)
	fmt.Printf("URL: %s\n", goURL)

	// // Create a request with a User-Agent header
	// req, err := http.NewRequest("GET", goURL, nil)
	// if err != nil {
	// 	fmt.Println("Error creating request:", err)
	// 	return
	// }

	// // Set custom headers to avoid blocking
	// req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36")

	// // Send HTTP request
	// resp, err := client.Do(req)
	// if err != nil {
	// 	// fmt.Println("Error sending request to", url, ":", err)
	// 	return
	// }
	// defer resp.Body.Close()
	// fmt.Printf("URL: %s, Status Code: %d\n", url, resp.StatusCode)

	// Check if the URL is empty
	// if goURL == "" {
	// 	fmt.Println("Error: URL is empty!")
	// 	return
	// }

	// Make the HTTP request
	resp, err := http.Get(goURL)
	if err != nil {
		// fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	// Print the response status code
	fmt.Println("Response", goURL, "status code:", resp.StatusCode)
	// fmt.Println("Request successful, status:", resp.Status)
}

//export mput
func mput(_keys []*C.char) int {
	var wg sync.WaitGroup
	// Process the URLs one by one (no large byte array)
	for _, _u := range _keys {
		url := C.GoString(_u)
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			makeRequest(C.CString(url))
		}(url)
	}

	wg.Wait() // Wait for all goroutines to finish

	return 0
}

//export gobuffchan
func gobuffchan(_keys []*C.char) int {
	requester := insrequester.NewRequester().Load()
	// var wg sync.WaitGroup
	ch := make(chan string, len(_keys))
	start := time.Now()
	defer func() {
		fmt.Printf("Execution time: %s\n", time.Since(start))
	}()

	// Process the URLs one by one (no large byte array)
	for _, _u := range _keys {
		url := C.GoString(_u)
		go func() {
			res, _ := requester.Get(insrequester.RequestEntity{Endpoint: url})
			if res == nil {
				ch <- fmt.Sprintf("%s: %s", url, "Error")
				return
			}
			ch <- fmt.Sprintf("%s: %d", url, res.StatusCode)
		}()
	}

	for range _keys {
		response := <-ch
		fmt.Println(response)
	}

	return 0
}

// //export mput
// func mput(_keys []*C.char, _val []*C.char) int {
// 	keys := make([]string, 0)
// 	for _, _k := range _keys {
// 		bs := C.GoString(_k) // Convert each C string to Go string (copy)
// 		keys = append(keys, bs)
// 	}
// 	fmt.Println("Keys:", keys)

// 	vals := make([]string, 0)
// 	for _, _v := range _val {
// 		bs := C.GoString(_v) // Convert each C string to Go string (copy)
// 		vals = append(vals, bs)
// 	}
// 	fmt.Println("Vals:", vals)

// 	return 0
// }

// func main() {}
