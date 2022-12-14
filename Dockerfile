FROM golang:1.19.4-alpine3.17 as builder

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./exec

RUN mkdir sql_empty

FROM scratch

WORKDIR /

COPY --from=builder /app/exec /exec

COPY --from=builder /app/sql_empty /sql

CMD ["/exec"]

# docker build -t sql-migration-pipeline .

# docker run --network=mysql-8 -v $(pwd)/tmp:/app/tmp -e ENVIRONMENT=test -e DB_HOST=mysql-8 -e DB_PORT=3306 -e DB_USER=root -e DB_PASS=root -e DB_DATABASE=api-teste sql-migration-pipeline
