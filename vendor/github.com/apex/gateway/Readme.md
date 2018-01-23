<img src="http://tjholowaychuk.com:6000/svg/title/APEX/GATEWAY">

Package gateway provides a drop-in replacement for net/http's `ListenAndServe` for use in AWS Lambda & API Gateway, simply swap it out for `gateway.ListenAndServe`. Extracted from [Up](https://github.com/apex/up) which provides additional middleware features and operational functionality.

```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/apex/gateway"
)

func main() {
	addr := ":" + os.Getenv("PORT")
	http.HandleFunc("/", hello)
	log.Fatal(gateway.ListenAndServe(addr, nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World from Go")
}
```

---

[![GoDoc](https://godoc.org/github.com/apex/up-go?status.svg)](https://godoc.org/github.com/apex/gateway)
![](https://img.shields.io/badge/license-MIT-blue.svg)
![](https://img.shields.io/badge/status-stable-green.svg)

<a href="https://apex.sh"><img src="http://tjholowaychuk.com:6000/svg/sponsor"></a>
