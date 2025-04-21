.PHONY: lineal pipeline semaphore benchmark benchmark_lineal benchmark_pipeline benchmark_semaphore

lineal:
	go build -tags lineal -o ./build/lineal cmd/main.go

pipeline:
	go build -tags pipeline -o ./build/pipeline cmd/main.go

semaphore:
	go build -tags semaphore -o ./build/semaphore cmd/main.go

benchmark:
	go test ./internal/logic -o "$(t).test" -run ^# -tags "$(t)" -bench . -benchmem -count="$(c)" -timeout 120m -v

benchmark_lineal:
	make benchmark t=lineal c=10

benchmark_pipeline:
	make benchmark t=pipeline c=10

benchmark_semaphore:
	make benchmark t=semaphore c=10