FROM golang:1.23

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

# собираем исполняемый файл
RUN go build -o brandscout-api ./cmd/main.go

# REST API
EXPOSE 8080
# используем не-рут юзера (по умолчанию рут). 
USER 1001

CMD ["./brandscout-api"]