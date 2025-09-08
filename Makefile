run:
	# @go run src/2024/11/main.go
	@watchexec -r -e odin,txt "odin run ./2019/01 -out:build/main -debug"

profile:
	@go run src/2024/11/main.go
	@go tool pprof -http 127.0.0.1:8080 cpu_profile.prof
