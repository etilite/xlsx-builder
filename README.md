# xlsx-builder

`xlsx-builder` is a lightweight microservice built with Go, designed to generate XLSX files from HTTP JSON requests.

## Features

- Fast and efficient XLSX file generation from http stream
- Easy-to-use API for creating Microsoft Excel™ spreadsheets from JSON data
- Dockerized for easy deployment

## Usage
### Quick Start with Docker

```sh
docker run --rm -p 8080:8080 -e HTTP_ADDR=:8080 etilite/xlsx-builder:latest
```

This will start the service and expose its API on port 8080.

### API

**Endpoint:**

- `POST /api/build`

**Request Body:**

The request body should be a JSON object with the following structure:

```json
[
   {"data": [1, "a", 2.1]},
   {"data": [2, "some-cell-data", 2, "another-cell"]}
]
``

**Request Example:**

Using `cURL`, you can make a request like this:

```sh
curl --location 'localhost:8080/api/build' \
--header 'Content-Type: application/json' \
--data '[
    {"data": [1, "a", 2.1]},
    {"data": [2, "some-cell-data", 2, "another-cell"]}
]' -o sheet.xlsx
```

**Response:**

The response will be a binary XLSX file with the generated content. 

### Build from source

```sh
git clone https://github.com/your-repo/xlsx-builder.git
cd xlsx-builder
make run
```
This will build and run app at `http://localhost:8080`.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

If you'd like to contribute to the project, please open an issue or submit a pull request on GitHub.

## Badges

[![docker pulls](https://img.shields.io/docker/pulls/etilite/xlsx-builder)](https://hub.docker.com/r/etilite/xlsx-builder)
[![docker push](https://github.com/etilite/xlsx-builder/actions/workflows/docker.yml/badge.svg)](https://github.com/etilite/xlsx-builder/actions/workflows/docker.yml)
[![go build](https://github.com/etilite/xlsx-builder/actions/workflows/go.yml/badge.svg)](https://github.com/etilite/xlsx-builder/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/etilite/xlsx-builder/graph/badge.svg?token=PYVPKWSEP1)](https://codecov.io/gh/etilite/xlsx-builder)