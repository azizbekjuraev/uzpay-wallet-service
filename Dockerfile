FROM golang:1.27

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping ./cmd/main.go

CMD ["/docker-gs-ping"]
