run:
	@go run src/2024/11/main.go
profile:
	@go run src/2024/11/main.go
	@go tool pprof -http 127.0.0.1:8080 cpu_profile.prof
