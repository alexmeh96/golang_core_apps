
FROM golang:1.21.5-alpine AS builder

WORKDIR /work

#RUN apk --no-cache add bash

COPY ["go.mod", "go.sum", "./"]
RUN go mod download && go mod verify

# build
COPY main.go ./
RUN go build -o ./bin/app main.go

FROM alpine AS runner

COPY --from=builder /work/bin/app /

CMD ["/app"]