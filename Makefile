.PHONY: lineal pipeline

lineal:
	go build -tags lineal -o ./build/lineal cmd/main.go

pipeline:
	go build -tags pipeline -o ./build/pipeline cmd/main.go

buffered:
	go build -tags buffered -o ./build/buffered cmd/main.go

pprof:
	go tool pprof -http=:8080 cpu.pprof