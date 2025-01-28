FROM go:alpine AS builder

WORKDIR /build

COPY . .

RUN go mod download

RUN go build -o go_ecommerce ./cmd/server

FROM scratch

COPY ./config /config

COPY --from=builder /build/go_ecommerce /

ENTRYPOINT [ "go_ecommerce", "config/local.yaml" ]