FROM golang:1.21
LABEL authors="michail"

COPY . /app
COPY .env /app/

WORKDIR /app

RUN go mod download && go mod tidy && \
    go install github.com/pressly/goose/v3/cmd/goose@latest

WORKDIR ./cmd

RUN CGO_ENABLED=0 GOOS=linux go build -o ../bin/main
WORKDIR /app

EXPOSE 3333

CMD [ "cd ./migrations | goose postgres $DATABASE_URL up" ]
CMD [ "/app/bin/main" ]
