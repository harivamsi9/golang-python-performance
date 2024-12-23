package main

/*
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"time"

	"github.com/useinsider/go-pkg/insrequester"
)

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

func main() {}
