# Math Rush Scoring Service

## Build

CGO_ENABLED=0 go build ./...

## Run

CGO_ENABLED=0 go run ./cmd/server

## Test

go test -count=1 ./...
