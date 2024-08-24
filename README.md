# xlsx-builder
[![docker pulls](https://img.shields.io/docker/pulls/etilite/xlsx-builder)](https://hub.docker.com/r/etilite/xlsx-builder)
[![docker push](https://github.com/etilite/xlsx-builder/actions/workflows/docker.yml/badge.svg)](https://github.com/etilite/xlsx-builder/actions/workflows/docker.yml)
[![go build](https://github.com/etilite/xlsx-builder/actions/workflows/go.yml/badge.svg)](https://github.com/etilite/xlsx-builder/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/etilite/xlsx-builder/graph/badge.svg?token=PYVPKWSEP1)](https://codecov.io/gh/etilite/xlsx-builder)

`xlsx-builder` is a lightweight microservice written in Go that allows you to easily generate XLSX spreadsheets from JSON requests.
This dedicated solution is perfect for projects looking to isolate their spreadsheet generation logic.
For applications that handle large datasets, xlsx-builder leverages Go's impressive performance and memory efficiency.
This makes it an optimal choice for enhancing scalability and maintainability while offloading labor-intensive tasks.

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

#### Request
- `POST /api/build`

**Request Body:**

The request body should be a JSON object with the following structure:

```json
[
   {"data": [11, "a", 2.1]},
   {"data": [22, "some-cell-data", 2, "another-cell"]}
]
```
Where each object in array presents a single row in table and `data` is array of cell values in this row.

**Request Example:**

Using `cURL`, you can make a request like this:

```sh
curl --location 'localhost:8080/api/build' \
--header 'Content-Type: application/json' \
--data '[
    {"data": [11, "a", 2.1]},
    {"data": [22, "some-cell-data", 2, "another-cell"]}
]' -o sheet.xlsx
```

#### Response

The response will be a binary XLSX file with the generated content. 
For this particular example it is a table with 2 rows and 4 cols:

|    |                |     |              |
|----|----------------|-----|--------------|
| 11 | a              | 2.1 |              |
| 22 | some-cell-data | 2   | another-cell |

### Build from source

```sh
git clone https://github.com/etilite/xlsx-builder.git
cd xlsx-builder
make run
```
This will build and run app at `http://localhost:8080`.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

If you'd like to contribute to the project, please open an issue or submit a pull request on GitHub.