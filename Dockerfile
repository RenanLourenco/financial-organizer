FROM golang:1.21 as builder

WORKDIR /app

COPY . ./

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o server

EXPOSE 5000

CMD ["./server"]

FROM alpine:3

COPY --from=builder /app/server /server

ENV GIN_MODE release
EXPOSE 5000

CMD ["./server"] 




# FROM golang:alpine AS builder

# RUN apk update && apk add --no-cache 'git=~2'
# ENV GO111MODULE=on

# WORKDIR /app

# COPY . ./

# RUN go get -d -v

# RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o server

# FROM alpine:3

# COPY --from=builder /app/server /server

# ENV GIN_MODE release
# EXPOSE 5000

# CMD ["./server"] 
