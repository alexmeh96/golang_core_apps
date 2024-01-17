FROM golang:1.21.5-alpine AS builder

WORKDIR /work

#RUN apk --no-cache add bash

COPY ["https_app/go.mod", "https_app/go.sum", "./"]
RUN go mod download

# build
COPY https_app ./
RUN go build -o ./bin/app main.go

FROM alpine AS runner

COPY --from=builder /work/bin/app /

CMD ["/app"]