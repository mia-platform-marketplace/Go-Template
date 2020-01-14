############################
# STEP 1 build executable binary
############################
FROM golang:1.13-alpine AS builder

RUN apk update && apk add --no-cache git

ARG COMMIT_SHA=<not-specified>

# Create appuser.
RUN adduser -D -g '' appuser

WORKDIR /app
COPY . .

RUN echo "$COMMIT_SHA" >> ./commit.sha

RUN go mod download
RUN go mod verify

RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-w -s" -o main .

############################
# STEP 2 build service image
############################

FROM scratch

LABEL maintainer="%CUSTOM_PLUGIN_CREATOR_USERNAME%" \
  name="%CUSTOM_PLUGIN_SERVICE_NAME%" \
  description="%CUSTOM_PLUGIN_SERVICE_DESCRIPTION%" \
  eu.mia-platform.url="https://www.mia-platform.eu"

# Import the user and group files from the builder.
COPY --from=builder /etc/passwd /etc/passwd

WORKDIR /app

COPY --from=builder /app/commit.sha /app/commit.sha

COPY --from=builder /app/main /app/main

# Use an unprivileged user.
USER appuser

ENTRYPOINT ["/app/main"]
