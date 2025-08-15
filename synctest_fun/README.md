# Synctest

Synctest is a testing utility for Go that allows you to simulate time and test concurrent code more easily.

### Run

```bash
./run.sh
```

### Results
```bash
❯ ./tests.sh
=== RUN   TestSynctest

=== Testing/Synctest Package ===
Synctest demo function created: func(*testing.T)
Virtual elapsed time: 1h0m0s
Real elapsed time: 113.703µs
--- PASS: TestSynctest (0.00s)

```bash
❯ ./tests.sh
=== RUN   TestSynctest

=== Testing/Synctest Package ===
Synctest demo function created: func(*testing.T)
Virtual elapsed time: 1h0m0s
Real elapsed time: 113.703µs
--- PASS: TestSynctest (0.00s)
PASS
ok  	github.com/diegopacheco/synctest_fun	(cached)
```