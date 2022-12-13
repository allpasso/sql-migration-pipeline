FROM golang:1.19.4-alpine3.17 as builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o ./exec

FROM alpine

WORKDIR /app

COPY --from=builder /app/exec /app/exec

RUN pwd && ls

CMD ["/app/exec"]

# docker build -t sql-migration-pipeline .

# docker run --network=mysql-8 -v $(pwd)/tmp:/app/tmp -e ENVIRONMENT=test -e DB_HOST=mysql-8 -e DB_PORT=3306 -e DB_USER=root -e DB_PASS=root -e DB_DATABASE=api-teste -e DIR_SQL_FILES=./tmp sql-migration-pipeline
