### Rationale

* Minimalistic rest service
* HTTP + Json
* No External Dependencies
* Small and simple
* Only 12 LoC

### Run

```bash
./run.sh
```

### Code

```Go
package main

import (
	"net/http"
)

func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "Hello, World!"}`))
	}))
}
```
