/*
Program used by https://github.com/richardpct/aws-terraform-tuto04
It creates a web server and displays the count of requests using a redis server
*/
package main

import (
	"flag"
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"net/http"
)

var redisHost = flag.String("redishost", "", "redis server")
var redisPass = flag.String("redispass", "", "redis password")
var redisdb *redis.Client

func checkArgs() error {
	if *redisHost == "" || *redisPass == "" {
		return fmt.Errorf("Arguments redishost and/or redispass are missing")
	}
	return nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	// increment count variable
	err := redisdb.Do("incr", "count").Err()
	if err != nil {
		log.Fatal(err)
	}
	// get count variable
	v, err := redisdb.Do("get", "count").String()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "counter = %q\n", v)
}

func main() {
	flag.Parse()
	if err := checkArgs(); err != nil {
		log.Fatal(err)
	}

	redisdb = redis.NewClient(&redis.Options{
		Addr:     *redisHost + ":6379",
		Password: *redisPass,
		DB:       0,
	})
	// check if redis server is reachable
	_, err := redisdb.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("0.0.0.0:80", nil))
}
