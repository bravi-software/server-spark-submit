package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

// version is definded during the build
var version string

func getServiceAddress() string {
	if env := os.Getenv("VIRTUAL_PORT"); env != "" {
		return ":" + env
	}

	if env := os.Getenv("HTTP_PORT"); env != "" {
		return ":" + env
	}

	return ":3000"
}

func handleSubmit(w http.ResponseWriter, r *http.Request, defaults []string) {
	var d struct{ Args []string }

	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	if len(d.Args) == 0 {
		http.Error(w, "Missing submit args", 400)
		return
	}

	args := append(defaults, d.Args...)

	out, err := exec.Command("spark-submit", args...).CombinedOutput()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	log.Println(out)

	w.WriteHeader(200)
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "version" {
		fmt.Println("server-spark-submit version " + version)
		return
	}

	defaults := strings.Split(os.Getenv("SPARK_SUBMIT_DEFAULT_ARGS"), " ")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			handleSubmit(w, r, defaults)
			return
		}

		http.NotFound(w, r)
	})

	addrs := getServiceAddress()
	log.Printf("Starting server at %s\n", addrs)
	log.Fatal(http.ListenAndServe(addrs, nil))
}
