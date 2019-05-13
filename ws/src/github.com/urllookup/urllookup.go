/*
 * URL Lookup Service
 * Handles Post Requests : POST {'url': url, 'status': 'ALLOW|BLOCK'}
 * and GET Requests : GET /urlinfo/1/{hostname_and_port}/{original_path_and_query_string}
 * The lookup service keeps the status of the url that is posted to it
 * and responds with the status on receiving a lookup 'GET' request
 */

package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"net/http"
	"strings"
)

/*
 * Create a new redis client
 * Assume the redis server is running on local host
 */
func RedisNewClient() (*redis.Client,error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "172.17.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	//Verify connection
	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println(pong, err)
		fmt.Println("Failed Connection to REDIS ")

	} else {
		fmt.Println("Connection to REDIS Successful")
	}
	return client,err
}

/*
 * Set Key - Value
 * Inputs : Redis Client Handle, Key, Value
 */
func RedisClientSet(client *redis.Client, key string, value string) {
	err := client.Set(key, value, 0).Err()
	if err != nil {
		panic(err)
	}
}

/*
 * Given a Key, read the corresponding Value
 * If key is not found, return "UNKNOWN" string.
 *
 * Inputs : Redis Client handle, key
 * Return Value.
 */
func RedisClientGet(client *redis.Client, key string) string {
	val, err := client.Get(key).Result()

	if err == redis.Nil {
		//fmt.Println(key," does not exist")
		val = "NOT FOUND"
	} else if err != nil {
		panic(err)
	}
	return val
}


func main() {
	// Create Redis Client Handle
	redisClient,err := RedisNewClient()
	if err != nil {
		return
	}


	http.HandleFunc("/", func (w http.ResponseWriter, req *http.Request) {

		switch req.Method {
		case "GET":
			// Expect "/{url}/ . Extract just the url string
			lookupUrl := strings.Split(req.URL.Path,"/")[1]
			fmt.Printf("Received GET request for url %s", lookupUrl)
			fmt.Fprintf(w, "{URL : %s, Status : %s}", req.URL,
				RedisClientGet(redisClient,lookupUrl))

		case "POST":
			// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
			if err := req.ParseForm(); err != nil {
				fmt.Fprintf(w, "ParseForm() err: %v", err)
				return
			}
			url := req.FormValue("url")
			status := req.FormValue("status")
			if status != "ALLOW" && status != "BLOCK" {
				fmt.Fprintf(w, "Invalid URL status %s NOT updated", status)
			} else {
				fmt.Fprintf(w, "URL status updated = %v\n",
					req.PostForm)
				RedisClientSet(redisClient, url, status)
			}
		}
	})

	//Start http server listening on port 8080
	http.ListenAndServe("0.0.0.0:8080", nil)

}

