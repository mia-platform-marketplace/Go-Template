############################
# STEP 1 build executable binary
############################
FROM golang:1.14 AS builder

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
  name="%CUSTOM_PLUGIN_SERVICE_NAME%" \
  description="%CUSTOM_PLUGIN_SERVICE_DESCRIPTION%" \
  eu.mia-platform.url="https://www.mia-platform.eu" \
  vcs.sha="$COMMIT_SHA"

WORKDIR /app

COPY --from=builder /app/build/* ./

# Use an unprivileged user.
USER 1000

CMD ["/app/main"]
