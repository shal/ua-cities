[license]: ./LICENSE
[goreport]: https://goreportcard.com/report/github.com/shal/ua-cities

# Ukrainian cities

[![Go Report Card](https://goreportcard.com/badge/github.com/shal/ua-cities)][goreport]

## Overview

`JSON` API to for getting location of specified Ukrainian city.

## Usage

```
$ http localhost:8080/city/Kiev
```

```json
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
