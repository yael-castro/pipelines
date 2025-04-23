# Solving a real problem with pipelines and semaphores
In this repository I have **recreated/simulated** a problem I faced at work some time ago.

Here I show three different ways to solve the problem:
1. Linear approach (Less optimal solution)
2. Using [pipelines](https://go.dev/blog/pipelines)
3. Using [pipelines](https://go.dev/blog/pipelines) and semaphores (Most optimal solution)
> ⚠️ Remember: the patterns are a conceptual idea more than a strict way to do something.
## Problem
Calculate the `net profit` for each `closing` in a company's stores.
###### Context
A company operates `stores`

Each store has multiple `closings` (accounting periods, such as monthly or quarterly)

For every `closing`, the company tracks:

- `sales` Revenue generated during the period.
- `costs` Expenses incurred to run the store during the period.
###### Objective
Develop a process to calculate and store the `net profit` of each `close`, where `net profit = sales - costs`
> ⚠️ Some details have been changed to avoid revealing confidential information but the problem remains the same.
## Solutions
### Linear ([See the code](internal/logic/logic_lineal.go))
###### How to run
```shell
make lineal
time ./build/lineal
```
###### Benchmarking
```shell
make benchmark_lineal
```
###### Diagram 
![Problem - Flow diagram](./docs/flow.svg)
### Pipelines ([See the code](internal/logic/logic_pipeline.go))
###### How to run
```shell
make pipeline
time ./build/pipeline
```
###### Benchmarking
```shell
make benchmark_pipeline
```
### Pipelines + Semaphores ([See the code](internal/logic/logic_semaphore.go))
###### How to run
```shell
make semaphore
time ./build/semaphore
```
###### Benchmarking
```shell
make benchmark_semaphore
```
## Resumen
###### Linear solution vs. Pipelines + Semaphores
```text
goos: darwin
goarch: arm64
pkg: github.com/yael-castro/pipelines/internal/logic
cpu: Apple M4
                      │   lineal.out   │            semaphore.out            │
                      │     sec/op     │   sec/op     vs base                │
Logic_CalculateProfit   328497.6m ± 3%   272.8m ± 4%  -99.92% (p=0.000 n=10)

                      │  lineal.out  │            semaphore.out             │
                      │     B/op     │     B/op      vs base                │
Logic_CalculateProfit   41.31Mi ± 0%   47.29Mi ± 0%  +14.48% (p=0.000 n=10)

                      │ lineal.out  │            semaphore.out            │
                      │  allocs/op  │  allocs/op   vs base                │
Logic_CalculateProfit   129.9k ± 0%   188.0k ± 1%  +44.70% (p=0.000 n=10)
```
> ⚠️ The benchmarks were performed using only one logic CPU at a time.
> ```go
> _ = runtime.GOMAXPROCS(1)
> ```