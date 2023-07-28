##STAGE : 1
FROM golang:latest as builder
LABEL maintainer="Anisur Rahman <sunny.cse7575@gmail.com>"
WORKDIR /app
COPY . .
RUN go mod download
RUN go build
##STAGE : 2
RUN CGO_ENABLED=0 go build -o main -a
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/.env .
EXPOSE 3030
CMD ["./main","startServer","-a=false"]
