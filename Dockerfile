############################
# STEP 1 build executable binary
############################
FROM golang:1.20 AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go mod verify

COPY . .

RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-w -s" -o main .

WORKDIR /app/build

RUN cp -r /app/main /app/LICENSE .

############################
# STEP 2 build service image
############################

FROM scratch

ARG COMMIT_SHA=<not-specified>

LABEL maintainer="%CUSTOM_PLUGIN_CREATOR_USERNAME%" \
  name="mia_template_service_name_placeholder" \
  description="%CUSTOM_PLUGIN_SERVICE_DESCRIPTION%" \
  eu.mia-platform.url="https://www.mia-platform.eu" \
  vcs.sha="$COMMIT_SHA"

ENV SERVICE_VERSION="0.0.0"

# Import the user and group files from the builder.
COPY --from=builder /etc/passwd /etc/passwd
# Import the certs from the builder.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

WORKDIR /app

COPY --from=builder /app/build/* ./

# Use an unprivileged user.
USER 1000

CMD ["/app/main"]
