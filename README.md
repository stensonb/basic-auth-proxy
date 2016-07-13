# basic-auth-proxy

This project is built as a simple proxy tool, offering basic auth protection for the services defined in the config.

# Usage

## Install

```
go get github.com/revinate/basic-auth-proxy
```

## Start

```
basic-auth-proxy --help
```

# Features

1. logs all requests
2. reads creds and service config from file
3. return 401 if basicAuth is wrong
4. prevent against basicAuth timing attacks
5. proxy request to configured backend if basicauth correct

# Roadmap

1. reload config on os.Signal
2. graceful shutdown
3. improve security by increasing request time for repeat offenders (ip based, user based, both?)
