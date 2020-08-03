FROM golang

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o webchat *.go

EXPOSE 8080

ENTRYPOINT ["/app/webchat"]