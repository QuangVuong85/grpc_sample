# Executor

![version](https://img.shields.io/badge/version-0.0.1-red) [![contributors](https://img.shields.io/badge/contributors-1-blue)]()

Executor is a simple worker pool implemented for Golang.

Features:

- Execute concurrent job by Goroutine
- Rate limiter
- Dynamic handler by reflect

### Usage

```go
executor, err := executor.NewExecutor(executor.DefaultExecutorConfig())

if err != nil {
  logrus.Error(err)
}

// close resource before quit
defer executor.Close()

for i := 0; i < 10; i++ {
  executor.Publish(func(input int) {
    fmt.Printf("2 * %d = %d \n", input, 2*input)
  }, i)

  executor.Publish(func(input int) {
    fmt.Printf("2 ^ %d = %d \n", input, input^2)
  }, i)

  executor.Publish(func(a int, b int) {
    fmt.Printf("%d + %d = %d \n", a, b, a+b)
  }, i, i+1)
}

=== RUN   TestExecute
2 * 0 = 0 
2 ^ 0 = 2 
0 + 1 = 1 
2 * 1 = 2 
2 ^ 1 = 3 
1 + 2 = 3 
2 * 2 = 4 
2 ^ 2 = 0 
2 + 3 = 5 
--- PASS: TestExecute (0.00s)
```
