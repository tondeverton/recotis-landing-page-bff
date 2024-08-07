# Renata Cotis's Landing Page BFF

This project is a REST API for [Renata Cotis's landing page](https://recotis.com). Nowadays this project provides two routes for the landing page which are `/signed-urls` and `/email`.

## Run

This is a Go project.

You can run using the current go version (^1.22) installed with this command:

```shell
go run ./cmd/recotis-landing-page-bff
```

Otherwise you can build and run a docker image with these commands:

* Build

    ```shell
    docker build -t recotis/bff .
    ```

* Run

    ```shell
    docker run -p 8080:8080 recotis/bff
    ```

## Test

There're unit tests available and you can run using this command:

```shell
go test ./...
```
