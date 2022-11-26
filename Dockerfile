FROM golang:1.19.3-bullseye
ADD . /app
WORKDIR /app
RUN go build -o main .

EXPOSE 8080
CMD ["/app/main"]