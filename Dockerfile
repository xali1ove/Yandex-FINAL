FROM golang:1.22.5-alpine

RUN apk add --no-cache git sqlite

WORKDIR /app

COPY . .

RUN go build -o todo-scheduler .

RUN mkdir -p /app/web
COPY web /app/web

ENV TODO_DBFILE=/app/scheduler.db
ENV TODO_PORT=7540

CMD ["sh", "-c", "./todo-scheduler -port ${TODO_PORT:-7540}"]

