FROM golang:1.20-bullseye as builder

WORKDIR /app/src

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 go build -o /app/bin/server

FROM gcr.io/distroless/static-debian11

COPY --from=builder /app/bin/ /app/bin/

EXPOSE 8090

ENTRYPOINT ["/app/bin/server"]
