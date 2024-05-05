# Xlsx builder
[![codecov](https://codecov.io/gh/etilite/xlsx-builder/graph/badge.svg?token=PYVPKWSEP1)](https://codecov.io/gh/etilite/xlsx-builder)

## Usage
Request `POST http://localhost:8080/table/`

Example JSON:
```JSON
{
    "header": [
        "date",
        "id",
        "price"
    ],
    "data": [
        [
            "01.01.2023",
            1,
            10.5
        ],
        [
            "02.01.2023",
            2,
            20.3
        ],
        [
            "03.01.2023",
            3,
            "33"
        ]
    ]
}
```
Service responses with `.xlsx` file