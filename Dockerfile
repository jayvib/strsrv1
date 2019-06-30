FROM golang:latest AS builder
ENV GO111MODULE=on
WORKDIR /go/src/github.com/jayvib/app/
COPY go.mod .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/src/github.com/jayvib/app/main .
RUN chmod +x ./main
EXPOSE 8080 
ENTRYPOINT ["./main"]
