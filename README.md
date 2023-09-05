# Xlsx builder
## Usage
`make dev-up`

`POST http://localhost:8080/invoice/`

```JSON
{
    "id": "1234",
    "date": "2023-07-29",
    "amount": "10",
    "client": {
        "fullName": "John Smith",
        "accountId": "US0001",
        "email": "random-name@gmail.com"
    },
    "header": [
        "date",
        "id",
        "price"
    ],
    "data": [
        [
            "01.01.2023",
            "1",
            "10"
        ],
        [
            "02.01.2023",
            "2",
            "20"
        ],
        [
            "03.01.2023",
            "3",
            "33"
        ]
    ]
}
```
service responses with `.xlsx` file