from ctypes import *
import time

# Define a structure for Go string
class go_string(Structure):
    _fields_ = [
        ("p", c_char_p),
        ("n", c_int)
    ]

# Define the GoSlice structure to represent a Go slice
class GoSlice(Structure):
    _fields_ = [("data", POINTER(c_void_p)),
                ("len", c_longlong), ("cap", c_longlong)]




# Load the shared library built from Go
# lib = cdll.LoadLibrary('./librequests.so')  # Make sure the shared library is in the same directory
lib = cdll.LoadLibrary('./librequestsbuff.so')  # Make sure the shared library is in the same directory

def bar(str):
    b = go_string(c_char_p(str.encode('utf-8')), len(str))
    lib.bar.restype = c_char_p
    a = lib.bar(b, c_char_p(str.encode('utf-8')))
    print(a.decode('utf-8'))  # Output the string returned from Go
    lib.freeme(a)  # Free the memory allocated by C.CString

# def pass_array_to_go(keys, values):
#     # Convert string list to a list of c_void_p (for pointer to char)
#     a0 = [cast(c_char_p(s.encode('utf-8')), c_void_p) for s in keys]
#     a1 = (c_void_p * len(a0))(*a0)
#     _keys = GoSlice(a1, len(a1), len(a1))
    
#     b0 = [cast(c_char_p(s.encode('utf-8')), c_void_p) for s in values]
#     b1 = (c_void_p * len(b0))(*b0)
#     _values = GoSlice(b1, len(b1), len(b1))

#     # Call Go function with the GoSlice
#     lib.mput(_keys, _values)


def pass_array_to_go(urls):
    # Convert list of strings to a list of c_void_p (for pointer to char)
    url_pointers = [cast(c_char_p(url.encode('utf-8')), c_void_p) for url in urls]
    url_array = (c_void_p * len(url_pointers))(*url_pointers)
    
    # Create a GoSlice to represent the array of URLs
    go_urls = GoSlice(url_array, len(url_pointers), len(url_pointers))
    
    # Call the Go function to send requests using goroutines
    # lib.mput(go_urls, len(urls))
    lib.gobuffchan(go_urls, len(urls))


# pass_array_to_go(keys, values)
# urls = ["https://www.google.com"]*100
urls = [
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
    "https://www.reddit.com",
    "https://www.paypal.com",
    "https://www.dropbox.com",
    "https://www.cloudflare.com",
    "https://www.airbnb.com",
    "https://www.spotify.com",
    "https://www.medium.com",
    "https://www.whatsapp.com",
    "https://www.tiktok.com",
    "https://www.foxnews.com",
    "https://www.nbcnews.com",
    "https://www.forbes.com",
    "https://www.huffpost.com",
    "https://www.bbc.com",
    "https://www.nike.com",
    "https://www.adidas.com",
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
    "https://www.linkedin.com",
    "https://www.microsoft.com",
    "https://www.airbnb.com",
    "https://www.gizmodo.com",
    "https://www.vogue.com",
    "https://www.merriam-webster.com",
    "https://www.imdb.com",
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
    "https://www.bestbuy.com",
    "https://www.target.com",
    "https://www.shopify.com",
    "https://www.samsung.com",
    "https://www.ford.com",
    "https://www.chevrolet.com",
    "https://www.toyota.com",
    "https://www.cars.com",
    "https://www.bmw.com",
    "https://www.honda.com",
    "https://www.hyundai.com",
    "https://www.rolex.com",
    "https://www.louisvuitton.com",
    "https://www.gucci.com",
    "https://www.prada.com",
    "https://www.chanel.com",
    "https://www.dior.com"
]

urls = set(urls)  # Remove duplicates
# urls.remove("https://www.adidas.com")
# urls.remove("https://www.cars.com")
# urls.remove("https://www.chanel.com")

urls.remove("https://www.adidas.com")
urls.remove("https://www.cars.com")
urls.remove("https://www.chanel.com")
urls.remove("https://www.gucci.com")
urls.remove("https://www.louisvuitton.com")


urls = sorted(urls)  # Sort URLs for consistent output
print("len(urls): ", len(urls))

start_time = time.time()
# pass_array_to_go(urls[:25])
pass_array_to_go(urls)

print("--- %s seconds ---" % (time.time() - start_time))

# time.sleep(100)
