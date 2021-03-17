package webserver

import (
	"log"
	"net/http"
	"net/http/httptest"
	"time"
	"weatherapi/webserver/services"
)

// cached is middleware to cache a request into memory
func cached(duration string, isJson bool, handler func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		content := services.Storage.Get(r.RequestURI, false)
		if content != nil {
			log.Println("Cache Hit!")
			if isJson {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
			}
			w.Write(content)
		} else {
			c := httptest.NewRecorder()
			handler(c, r)

			for k, v := range c.HeaderMap {
				w.Header()[k] = v
			}

			w.WriteHeader(c.Code)
			content := c.Body.Bytes()

			if d, err := time.ParseDuration(duration); err == nil {
				log.Printf("New response cached: %s for %s\n", r.RequestURI, duration)
				services.Storage.Set(r.RequestURI, content, d)
			} else {
				log.Printf("response not cached. err: %s\n", err)
			}

			w.Write(content)
		}

	})
}
