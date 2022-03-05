FROM 'golang:1.17.8-alpine3.15'

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY ./ ./

RUN go build -o /consumer cmd/messageConsumer/main.go

CMD [ "/consumer" ]