FROM golang:1.14.4-alpine3.12 as builder
WORKDIR /go/src/app
COPY . .
 
RUN CGO_ENABLED=0 GOOS=linux go build -v -o main
 
FROM gcr.io/cloud-builders/gcloud:latest
 
WORKDIR /usr/local/bin/
 
COPY --from=builder /go/src/app/main /main
 
ENTRYPOINT ["/main"]