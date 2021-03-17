# Stage 1
FROM golang:1.16.0-alpine3.13 AS dependency_builder

WORKDIR /go/src

RUN apk update
RUN apk add --no-cache bash ca-certificates git

COPY go.mod .
COPY go.sum .

RUN go mod download

# Stage 2
FROM dependency_builder AS service_builder

WORKDIR /usr/app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o bin

# Stage 3
FROM alpine:latest  

RUN apk --no-cache add ca-certificates tzdata
WORKDIR /root/
COPY --from=service_builder /usr/app/bin .
COPY --from=service_builder /usr/app/config.json .

ENTRYPOINT ["./bin"]
