# Solving a real problem with pipelines and semaphores
This repository is intended to show how I solve a real problem
using the concurrent patterns **pipelines** and **semaphores**
> ⚠️ Remember: the patterns are a conceptual idea more than a strict way to do something.

## Problem
Calculate the `net profit` for each `closing` in a company's stores.

> ⚠️ Some details have been modified to avoid revealing confidential information.
###### Context
A company operates `stores`.

Each store has multiple `closings` (accounting periods, such as monthly or quarterly closings).

For every `closing`, the company tracks:

- `Sales` Revenue generated during the period.
- `Operation costs` Expenses incurred to run the store during the period.
###### Objective
Compute the `net profit` for each `closing`, where:
```text
Net profit = Sales − Operation costs
```
## Solutions
> ⚠️ I limited CPUs that can be using simultaneously to one for each solution
### Lineal ([See the code](internal/logic/logic_lineal.go))
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