FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o /docker-miniproject

EXPOSE 8080

CMD [ "/docker-miniproject" ]