package main

import (
  "log"
  "net/http"
  "fmt"
  "github.com/garyburd/redigo/redis"
)

func get_heros(w http.ResponseWriter, req *http.Request) {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		log.Fatalf("Could not connect: %v\n", err)
	}
	defer c.Close()

	ret, _ := c.Do("SET","foo", "bar")
	fmt.Printf("%s\n", ret)

	ret, _ = c.Do("GET","foo")
	fmt.Printf("%s\n", ret)
}

func nop(w http.ResponseWriter, req *http.Request) {
	// nothing to see here
}

func build_endpoints() {
	http.HandleFunc("/", nop )
	http.HandleFunc("/get_heros", get_heros )
}

func main() {

	build_endpoints()

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}
