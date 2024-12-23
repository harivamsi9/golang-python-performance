package main

// Pure GO with Insrequester + channels + goroutines

import (
	"fmt"
	"time"

	"github.com/useinsider/go-pkg/insrequester"
)

func main() {
	requester := insrequester.NewRequester().Load()

	// requester := insrequester.NewRequester().Load()

	urls := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.youtube.com",
		"https://www.amazon.com",
		"https://www.wikipedia.org",
		"https://www.reddit.com",
		"https://www.twitter.com",
		"https://www.instagram.com",
		"https://www.linkedin.com",
		"https://www.netflix.com",
		"https://www.github.com",
		"https://www.apple.com",
		"https://www.microsoft.com",
		"https://www.yahoo.com",
		"https://www.bing.com",
		"https://www.stackoverflow.com",
		"https://www.quora.com",
		"https://www.nytimes.com",
		"https://www.cnn.com",
		"https://www.pinterest.com",
		"https://www.tumblr.com",
		"https://www.ebay.com",
		"https://www.imdb.com",
		"https://www.spotify.com",
		"https://www.slack.com",
		"https://www.paypal.com",
		"https://www.dropbox.com",
		"https://www.cloudflare.com",
		"https://www.airbnb.com",
		"https://www.medium.com",
		"https://www.whatsapp.com",
		"https://www.tiktok.com",
		"https://www.foxnews.com",
		"https://www.nbcnews.com",
		"https://www.forbes.com",
		"https://www.huffpost.com",
		"https://www.bbc.com",
		"https://www.nike.com",
		"https://www.uber.com",
		"https://www.lyft.com",
		"https://www.weather.com",
		"https://www.walmart.com",
		"https://www.target.com",
		"https://www.bestbuy.com",
		"https://www.cnet.com",
		"https://www.reuters.com",
		"https://www.telegraph.co.uk",
		"https://www.theguardian.com",
		"https://www.businessinsider.com",
		"https://www.usatoday.com",
		"https://www.buzzfeed.com",
		"https://www.cbsnews.com",
		"https://www.squarespace.com",
		"https://www.stripe.com",
		"https://www.godaddy.com",
		"https://www.wix.com",
		"https://www.gizmodo.com",
		"https://www.vogue.com",
		"https://www.merriam-webster.com",
		"https://www.flickr.com",
		"https://www.behance.net",
		"https://www.ted.com",
		"https://www.buzzfeednews.com",
		"https://www.dailymotion.com",
		"https://www.giphy.com",
		"https://www.vimeo.com",
		"https://www.vice.com",
		"https://www.theverge.com",
		"https://www.npr.org",
		"https://www.la.com",
		"https://www.thechive.com",
		"https://www.wework.com",
		"https://www.lazada.com",
		"https://www.mercadolivre.com",
		"https://www.aliexpress.com",
		"https://www.shopify.com",
		"https://www.samsung.com",
		"https://www.ford.com",
		"https://www.chevrolet.com",
		"https://www.toyota.com",
		"https://www.bmw.com",
		"https://www.honda.com",
		"https://www.hyundai.com",
		"https://www.rolex.com",
		"https://www.prada.com",
		"https://www.dior.com",
	}

	ch := make(chan string, len(urls))
	fmt.Println("len(urls):", len(urls))
	start := time.Now()
	defer func() {
		fmt.Printf("Execution time: %s\n", time.Since(start))
	}()
	for _, url := range urls {
		go func() {
			res, _ := requester.Get(insrequester.RequestEntity{Endpoint: url})
			if res == nil {
				ch <- fmt.Sprintf("%s: %s", url, "Error")
				return
			}
			ch <- fmt.Sprintf("%s: %d", url, res.StatusCode)
		}()
	}

	for range urls {
		response := <-ch
		fmt.Println(response)
	}

}
