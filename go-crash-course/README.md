# Go Crash Course

This mirrors your Rust crash course with idiomatic Go examples. Each example is a separate Go file and is invoked from `main.go`. Every example prints its name first.

## Run

```bash
./run.sh
```

## Topics

- if / else (`ex_if_else.go`)
  - Branch on conditions to control flow. Conditions must be boolean.
  - Example mirrors parity/divisibility checks.

  Code:
  ```go
  x := 7
  if x%2 == 0 {
	  fmt.Println("even")
  } else if x%3 == 0 {
	  fmt.Println("divisible by three")
  } else {
	  fmt.Println("odd")
  }
  ```

- Option unwrap (`ex_option_unwrap.go`)
  - Go uses pointers and zero-values in place of `Option<T>`. Check for nil for presence, use defaults otherwise.

  Code:
  ```go
  var a *int
  v := 10
  a = &v 
  var b *int

  if a != nil { fmt.Println(*a) }
  y := 0
  if b != nil { y = *b }
  fmt.Println("unwrap_or:", y)
  if a != nil { fmt.Println((*a) * 2) }
  ```

- for loops (`ex_for_loops.go`)
  - Classic `for` and `range` over slices. Reverse iteration done by indexing.

  Code:
  ```go
  sum := 0
  for i := 0; i < 5; i++ { sum += i }
  v := []int{1,2,3}
  for _, x := range v { fmt.Println(x) }
  for i := len(v)-1; i >= 0; i-- { fmt.Println(v[i]) }
  ```

- traits (`ex_traits.go`)
  - Rust traits map to Go interfaces. Implement methods on structs and accept interfaces.

  Code:
  ```go
  type Greeter interface{ Greet() string }
  type Person struct{ Name string }
  func (p Person) Greet() string { return "hello "+p.Name }
  func greetAll(items []Greeter){ for _,it := range items { fmt.Println(it.Greet()) } }
  ```

- HTTP call (`ex_http.go`)
  - Use `net/http` with context for timeout. Print status and body length.

  Code:
  ```go
  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
  defer cancel()
  req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "https://httpbin.org/get", nil)
  resp, _ := http.DefaultClient.Do(req)
  defer resp.Body.Close()
  b, _ := io.ReadAll(resp.Body)
  fmt.Println(resp.Status, len(b))
  ```

- JSON serde (`ex_json_serde.go`)
  - Use `encoding/json` to marshal/unmarshal structs with tags.

  Code:
  ```go
  type User struct{ ID uint32 `json:"id"`; Name string `json:"name"`; Active bool `json:"active"` }
  s, _ := json.Marshal(User{ID:1, Name:"neo", Active:true})
  var d User
  _ = json.Unmarshal(s, &d)
  ```

- Shell command (`ex_shell_cmd.go`)
  - Run commands with `os/exec` and split output by lines.

  Code:
  ```go
  out, _ := exec.Command("sh", "-c", "echo go && uname -s").Output()
  for _, l := range strings.Split(strings.TrimSpace(string(out)), "\n") { fmt.Println(l) }
  ```

- Borrow checker (`ex_borrow_checker.go`)
  - Go has no borrow checker; pointers/values illustrate the concept differences.

  Code:
  ```go
  s := "abc"
  _ = lenOf(&s)
  r1, r2 := s, s
  s = s + "d"
  fmt.Println(r1, r2, s)
  ```

- Unit tests (`ex_unit_tests.go`, `ex_unit_tests_test.go`)
  - Place tests in `_test.go` files and run with `go test ./...`.

  Code:
  ```go
  func add(a,b int) int { return a+b }
  // ex_unit_tests_test.go:
  func TestAdd(t *testing.T){ 
    if add(2,3)!=5 { 
        t.Fatal("bad") 
        } 
  }
  ```

- Collections (`ex_collections.go`)
  - Slices and maps; use `slices`/`maps` (Go 1.21+) helpers and sorted key iteration.

  Code:
  ```go
  v := []int{3,1,2}; slices.Sort(v)
  h := map[string]int{"a":1}
  h2 := maps.Clone(h)
  keys := make([]string,0,len(h)); for k := range h { keys = append(keys,k) }
  slices.Sort(keys); for _, k := range keys { fmt.Println(k, h[k]) }
  ```

- Error handling (`ex_error_handling.go`)
  - Use `(T, error)` returns and handle with `if err != nil { ... }`.

  Code:
  ```go
  data, err := os.ReadFile("file.txt")
  if err != nil { /* handle */ }
  _ = string(data)
  ```

- Functional programming (`ex_functional.go`)
  - Map/filter/reduce patterns via loops; closures via function literals.

  Code:
  ```go
  v := []int{1,2,3,4,5}
  tmp := make([]int,0,len(v)); for _,x := range v { tmp = append(tmp, x*2) }
  f := make([]int,0,len(tmp)); for _,x := range tmp { if x%3==0 { f = append(f,x) } }
  sum := 0; for _,x := range f { sum += x }
  f1 := func(x int) int { return x+1 }
  _ = f1(9)
  ```

- Coroutines (`ex_coroutines.go`)
  - Go's goroutines and channels for concurrency.

  Code:
  ```go
  ch := make(chan string)
  go func(){ ch <- "ping" }()
  go func(){ ch <- "pong" }()
  fmt.Println(<-ch, <-ch)
  ```

## Tests

```bash
go test ./...
```
