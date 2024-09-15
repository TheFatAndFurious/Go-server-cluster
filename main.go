package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func generatePorts(initialPort int, numberOfPorts int) []int {
	ports := make([]int, 5)
	for i := 0; i < numberOfPorts; i++ {
		ports[i] = initialPort + i
		fmt.Printf("port is %v\n", ports[i])
	}
	return ports
}

func createServer(port int) {
	mux := http.NewServeMux()
	mux.HandleFunc("/Hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from port %v", port)
	})
	fmt.Printf("server up and running on port %v\n", port)
	err := http.ListenAndServe(":"+strconv.Itoa(port), mux)
	if err != nil {
		fmt.Printf("server on port %v failed to start: %v\n", port, err)
		return
	}
}

func main() {
	//create a http server

	ports := generatePorts(8090, 5)
	for i, _ := range ports {
		fmt.Println(ports[i])
		go createServer(ports[i])
	}
	select {}
}
