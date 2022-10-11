FROM golang:alpine as builder

WORKDIR /go/src

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /go/bin/main .


FROM scratch

COPY --from=builder /go/bin/main .

EXPOSE 80

ENTRYPOINT ["/main"]