FROM golang:1.20.2-alpine3.17 as build

# install common tools
RUN apk update && apk upgrade && apk add ca-certificates bash git openssh gcc g++ pkgconfig build-base curl \
    && rm -rf /var/cache/apk/*

# Build App
WORKDIR /src

RUN apk add --no-cache tzdata

COPY go.mod go.sum ./

RUN go mod download


COPY /src  ./src/

# update swagger docs/
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init -g src/infrastructure/api/app.go

WORKDIR /src/src/cmd
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o /arq2-tp1-app

# final stage
FROM gcr.io/distroless/static-debian11:latest

COPY --from=build /arq2-tp1-app /app/
WORKDIR /app

CMD ["./arq2-tp1-app"]