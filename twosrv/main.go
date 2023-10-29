package main

import "fmt"
import "time"
import "net/http"
import "github.com/bradfitz/gomemcache/memcache"

func main() {
	fmt.Println("TEST")
	memc := memcache.New("twosrv-memcache:11211")
	memc.Timeout = 5 * time.Second

	item := &memcache.Item{
		Key:        "gogogo",
		Value:      []byte("give me somebody shut"),
		Expiration: 1000,
	}

	memc.Set(item)

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte(fmt.Sprintf("call url %s", r.URL)))
		item, err := memc.Get("gogogo")

		if err != nil {
			panic(err)
		}

		w.Write([]byte(item.Value))

	}))

	http.ListenAndServe(":8080", nil)
}
