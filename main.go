package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"time"
	// "github.com/go-redis/redis"
)

const message = "This is new golang docker and kubernetes example"

// var redisdb *redis.Client

func main() {
	// initRedis()
	// initRedisSets()
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-type", "text/plain; charset=utf-8")
		w.Write([]byte(message))
	})

	// mux.HandleFunc("/increaseVisit", increaseVisit)
	srv := NewServer(mux)

	fmt.Println("Listening to Port  :8081")
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}

}

// func initRedis() {
// 	redisdb = redis.NewClient(&redis.Options{
// 		Addr:         "redis-server:6379",
// 		DialTimeout:  10 * time.Second,
// 		ReadTimeout:  30 * time.Second,
// 		WriteTimeout: 30 * time.Second,
// 		PoolSize:     10,
// 		PoolTimeout:  30 * time.Second,
// 	})
// }

// func initRedisSets() {
// 	if err := redisdb.Set("visits", 0, 0).Err(); err != nil {
// 		panic(err)
// 	}

// }

// func increaseVisit(w http.ResponseWriter, r *http.Request) {
// 	client := redis.NewClient(&redis.Options{
// 		Addr:     "redis-server:6379",
// 		Password: "",
// 		DB:       0,
// 	})

// 	_, err := client.Ping().Result()
// 	if err != nil {
// 		panic(err)
// 	}

// 	// fmt.Fprintf(w, "Hello, %q", pong)

// 	visit, err := client.Get("visits").Result()
// 	if err != nil {
// 		fmt.Println("Error Fetching the Visits", err)
// 	}

// 	w.Write([]byte("Number of visit: " + visit))

// 	visNumber, err := strconv.Atoi(visit)
// 	if err != nil {
// 		fmt.Println("Error fetching the visit number")
// 	}

// 	if err := client.Set("visits", visNumber+1, 0).Err(); err != nil {
// 		fmt.Println("Error Updateing visit number")
// 	}

// }

// NewServer Create a new server
func NewServer(mux *http.ServeMux) *http.Server {
	tlsConfig := &tls.Config{
		PreferServerCipherSuites: true,
		CurvePreferences: []tls.CurveID{
			tls.CurveP256,
			tls.X25519, //Go 1.8 only
		},
		MinVersion: tls.VersionTLS12,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		},
	}

	srv := &http.Server{
		Addr:         ":8081",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		TLSConfig:    tlsConfig,
		Handler:      mux,
	}

	return srv
}
