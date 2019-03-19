# Pengendali API (Reverse proxy for API Gateway)

[![Build Status](https://travis-ci.org/agungdwiprasetyo/reverse-proxy.svg?branch=master)](https://travis-ci.org/agungdwiprasetyo/reverse-proxy)
[![codecov](https://codecov.io/gh/agungdwiprasetyo/reverse-proxy/branch/master/graph/badge.svg)](https://codecov.io/gh/agungdwiprasetyo/reverse-proxy)
[![golang](https://img.shields.io/badge/GoLang-v1.11-green.svg?logo=google)](https://golang.org/doc/devel/release.html#go1.11)

[![Coverage Graph](https://codecov.io/gh/agungdwiprasetyo/reverse-proxy/branch/master/graphs/commits.svg)](https://codecov.io/gh/agungdwiprasetyo/reverse-proxy)

A http tool for reverse proxy like nginx

### Use
* Install Golang and get dependencies
```sh
$ brew install golang
$ glide install
```

* Build app
```sh
$ make build
```

* Copy `config.example.json` to `config.json`
```sh
$ cp config.example.json config.json
```

* Add your proxy in `config.json`
```json
{
    "proxy": [
        {
            "root": "[root for proxy, example: /myapp/]",
            "host": "[host proxy, example: http://localhost:8000]"
        }
    ]
}
```

* Set app port
```json
{
    "gatewayPort": 3000,
}
```

* Run binary
```sh
$ ./bin
[GATEWAY] :3000/myapp/ |===> http://localhost:8000
Server running on port :3000
```

In nginx, the configuration for adding proxy is like this:
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