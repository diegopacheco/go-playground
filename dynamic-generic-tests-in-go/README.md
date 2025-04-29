### Run

```bash
./run.sh
```

### Tests

```bash
❯ ./tests.sh
=== RUN   TestAdd
Add(1, 2) = 3; expected 3
Add(2, 3) = 5; expected 5
Add(-3, 4) = 1; expected 1
Add(0, 0) = 0; expected 0
--- PASS: TestAdd (0.00s)
=== RUN   TestDivision
Divide(1, 2) = 0; expected 0
Divide(2, 3) = 0; expected 0
Divide(-3, 4) = 0; expected 0
Divide(0, 0) = 0; expected 0
--- PASS: TestDivision (0.00s)
=== RUN   TestMultiplyAndSub
* (1, 2) = 2; expected 2
* (2, 3) = 6; expected 6
* (-3, 4) = -12; expected -12
* (0, 0) = 0; expected 0
- (1, 2) = -1; expected -1
- (2, 3) = -1; expected -1
- (-3, 4) = -7; expected -7
- (0, 0) = 0; expected 0
--- PASS: TestMultiplyAndSub (0.00s)
PASS
ok  	github.com/diegopacheco/dynamic-generic-tests-in-go/math	0.002s
```