[license]: ./LICENSE
[goreport]: https://goreportcard.com/report/github.com/ashanaakh/ua-cities

# Ukrainian cities

[![Go Report Card](https://goreportcard.com/badge/github.com/ashanaakh/ua-cities)][goreport]

## Overview

`JSON` API to for getting location of specified Ukrainian city.

## Usage

```
$ curl localhost:8080/city/:city
```

```JSON
{
  "name": "Kiev",
  "location": {
  "lat": "50.584981",
  "lon": "30.235748"
  }
}
```

## License

Project released under the terms of the MIT [license][license].

