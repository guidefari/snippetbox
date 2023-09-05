FROM golang:1.21 as base

FROM base as dev

RUN go install github.com/cosmtrek/air@latest

WORKDIR /app
CMD ["air"]
