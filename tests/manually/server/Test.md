execute at the project/workspace folder 

*1 terminal tab for each command*
    go run ./tests/server/server.go -port 8081
    go run ./tests/server/server.go -port 8082
    go run ./cmd/main.go -backends "http://localhost:8081 | http://localhost:8082"
