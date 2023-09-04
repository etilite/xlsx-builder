#ARG BUILDER_IMAGE=golang:buster
ARG GO_VERSION=1.20.2
ARG DISTROLESS_IMAGE=gcr.io/distroless/static:nonroot
#ARG DISTROLESS_IMAGE=gcr.io/distroless/base
############################
# STEP 1 build executable binary
############################
#FROM ${BUILDER_IMAGE} as builder
FROM golang:${GO_VERSION} as builder

# Ensure ca-certficates are up to date
RUN update-ca-certificates

#WORKDIR $GOPATH/src/mypackage/myapp/
WORKDIR /app/

# use modules
#COPY go.mod .
COPY go.mod go.sum ./

ENV GO111MODULE=on
RUN go mod download && go mod verify

COPY . .

# Build the static binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
      -ldflags='-w -s -extldflags "-static"' -a \
      -o app .

#RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o app .

############################
# STEP 2 build a small image
############################
# using static nonroot image
# user:group is nobody:nobody, uid:gid = 65534:65534
FROM ${DISTROLESS_IMAGE}

# Copy our static executable
#COPY --from=builder /app/app /app
COPY --from=builder --chown=nonroot:nonroot /app/app /app

# Run the hello binary.
#ENTRYPOINT ["/go/bin/hello"]
ENTRYPOINT ["/app"]