ARG GO_VERSION=1.21.1
ARG DISTROLESS_IMAGE=gcr.io/distroless/static:nonroot
############################
# STEP 1 build executable binary
############################
FROM golang:${GO_VERSION} as builder

# Ensure ca-certficates are up to date
RUN update-ca-certificates

WORKDIR /app/

# use modules
COPY go.mod go.sum ./

ENV GO111MODULE=on
RUN go mod download && go mod verify

COPY . .

# Build the static binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
      -ldflags='-w -s -extldflags "-static"' -a \
      -o app ./cmd/xlsx-builder

############################
# STEP 2 build a small image
############################
FROM ${DISTROLESS_IMAGE}

COPY --from=builder --chown=nonroot:nonroot /app/app /app

ENTRYPOINT ["/app"]