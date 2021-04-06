FROM golang:1.14 AS build
RUN mkdir -p /golang/wizard-enhance
WORKDIR /golang/wizard-enhance
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-s -w -X main.Version=latest -X main.GitCommit=24130b9704a9cd398932c3f0d2262b8568e02e65' -o wizard-enhance cmd/main.go

FROM ubuntu:20.10
WORKDIR /root
COPY --from=build /golang/wizard-enhance/wizard-enhance .
EXPOSE 19921
CMD ["./wizard-enhance", "--listen", "127.0.0.1:19921"]