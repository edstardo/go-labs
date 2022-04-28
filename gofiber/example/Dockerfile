FROM golang:1.18.1-alpine3.15 as build

WORKDIR /code

COPY go.* .

RUN go mod download

COPY *.go .

RUN go build -o api *.go

FROM alpine:3.15

WORKDIR /

COPY --from=build /code/api .

EXPOSE 3000

ENTRYPOINT [ "./api" ]
