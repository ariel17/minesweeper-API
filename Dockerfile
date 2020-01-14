# build stage
FROM golang as builder

ENV GO111MODULE=on
ENV PUBLIC_URL=http://localhost:8080

WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

# final stage
FROM scratch
COPY --from=builder /app/minesweeper-API /app/
EXPOSE 8080
ENTRYPOINT ["/app/minesweeper-API"]
