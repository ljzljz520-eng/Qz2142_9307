FROM golang:1.24.13
WORKDIR /app
ENV CGO_ENABLED=0 GOPROXY=https://proxy.golang.org,direct
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build ./...
RUN go test -run '^$' ./...
CMD ["go", "run", "./cmd/server"]
