package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net"
)
func dialTimeout(network, addr string) (net.Conn, error) {
    return net.DialTimeout(network, addr, 1000)
}
func main() {
	
	for
	{
		transport := http.Transport{
                Dial:              dialTimeout,
       		}

        	client := http.Client{
                	Transport: &transport,
        	}
		resp, err := client.Get("https://jsonplaceholder.typicode.com/posts")
		if err != nil {
			log.Fatalln(err)
		}
		//We Read the response body on the line below.
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		//Convert the body to type string
		sb := string(body)
		log.Printf("%s", sb)
	}

}
