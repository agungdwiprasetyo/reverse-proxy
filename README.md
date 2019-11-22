# Pengendali API (Reverse proxy for API Gateway)

[![Build Status](https://travis-ci.org/agungdwiprasetyo/reverse-proxy.svg?branch=master)](https://travis-ci.org/agungdwiprasetyo/reverse-proxy)
[![codecov](https://codecov.io/gh/agungdwiprasetyo/reverse-proxy/branch/master/graph/badge.svg)](https://codecov.io/gh/agungdwiprasetyo/reverse-proxy)
[![golang](https://img.shields.io/badge/GoLang-v1.12-green.svg?logo=google)](https://golang.org/doc/devel/release.html#go1.12)

[![Coverage Graph](https://codecov.io/gh/agungdwiprasetyo/reverse-proxy/branch/master/graphs/commits.svg)](https://codecov.io/gh/agungdwiprasetyo/reverse-proxy)

A http tool for reverse proxy like nginx **(NO Framework, just write in PURE Go)**

### Use
* Install Golang
```sh
$ brew install golang
```

* Build app
```sh
$ make build
```

* Copy `config.example.json` to `config.json`
```sh
$ cp config.example.json config.json
```

* Add your services in `config.json`
```json
{
    "services": [
        {
            "root": "[root for service, example: /myapp]",
            "host": "[host service, example: http://localhost:8000]"
        }
    ]
}
```

* Set app port
```json
{
    "gateway_port": 3000,
}
```

* Run binary
```sh
$ ./bin
[GATEWAY] :3000/myapp/ |===> http://localhost:8000
Server running on port :3000
```

In nginx, configuration for adding proxy like this:
```
location /myapp/ {
    proxy_pass http://127.0.0.1:8000;
    proxy_set_header Host $host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_redirect    off;
}
```

In this app, for adding a proxy just add in `config.json`