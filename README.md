# Fizz-Buzz REST Server

This is a simple Fizz-Buzz REST server implemented in Golang using the standard Go HTTP server and router.

It provides an API endpoint for generating Fizz-Buzz sequences and retrieving statistics about the most common requests.

## Features

- Accepts parameters for generating Fizz-Buzz sequences.
- Exposes an API endpoint for generating Fizz-Buzz sequences.
- Provides statistics about the most common requests.
- Ready for production: Graceful Shutdown, Containerization.
- Easy to maintain: lot of comments, unit tests.

## Getting Started

1. Clone this repository:

```bash
git clone https://github.com/yourusername/fizzbuzz-server.git
cd fizzbuzz-server
```

2. Build the Fizz-Buzz server:

```bash
make run
```

3. Test endpoint fizzbuzz:

```bash
$ curl -s -d '{"int1":3, "int2":5, "limit":15, "str1":"fizz", "str2":"buzz"}' -H "Content-Type: application/json" -X POST http://localhost:8000/fizzbuzz | jq
[
  "1",
  "2",
  "fizz",
  "4",
  "buzz",
  "fizz",
  "7",
  "8",
  "fizz",
  "buzz",
  "11",
  "fizz",
  "13",
  "14",
  "fizzbuzz"
]
```

4. Test stats endpoint

```bash
$ curl -s  http://localhost:8000/stats | jq
{
  "hits": 1,
  "most_common_request": {
    "int1": 3,
    "int2": 5,
    "limit": 15,
    "str1": "fizz",
    "str2": "buzz"
  }
}
```