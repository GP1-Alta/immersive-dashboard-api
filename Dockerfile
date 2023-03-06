
FROM golang:1.20.1-alpine

RUN mkdir /app

WORKDIR /app

COPY ./ /app

RUN go mod tidy

RUN go build -o immersivedashboard

CMD ["./immersivedashboard"]