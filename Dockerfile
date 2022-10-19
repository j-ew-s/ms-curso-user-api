FROM golang:1.17.3

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./



RUN go build -o  /ms-curso-user-api

EXPOSE 5100

CMD ["/ms-curso-user-api"]