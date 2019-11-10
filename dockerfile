FROM golang as builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod /app
COPY go.sum /app

RUN go mod download

COPY . /app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o vajari

FROM scratch

COPY --from=builder /app/vajari /app/

EXPOSE 3479

ENTRYPOINT ["./app/vajari"]