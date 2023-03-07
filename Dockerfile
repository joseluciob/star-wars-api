FROM golang:alpine AS base
    RUN apk update && apk add --no-cache git && apk add --no-cache bash && apk add build-base

    RUN mkdir /app
    WORKDIR /app

    COPY . .
    COPY .env .

    RUN go get -d -v ./...
    RUN go install -v ./...

FROM base AS compiler_importer
    RUN go build -o /bin/importer ./cmd/importer

FROM base AS swapi_app

    RUN go install -mod=mod github.com/githubnemo/CompileDaemon
    RUN go get -v golang.org/x/tools/gopls

    COPY --from=compiler_importer --chown=us /bin/importer /bin/importer

    ENTRYPOINT CompileDaemon --build="go build -a -installsuffix cgo -o /bin/server ./cmd/server" --command=/bin/server