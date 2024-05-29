FROM golang:1.22.3-alpine3.19

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o ./githubaction

CMD ["/app/githubaction"]