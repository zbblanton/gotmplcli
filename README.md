# Description

Simple CLI for templating in shell scripts or whatever else you want using go templates. Inspired by [j2cli](https://pypi.org/project/j2cli/).

# Usage

``` bash
gotmplcli --values <VALUES FILE> --template <TEMPLATE FILE> --output <OUTPUT FILE>
```

`--output` is optional. If not provided, the CLI prints to stdout.

# Example

example-values.yaml
``` yaml
---
user: "www-data"
error_log: "logs/error.log"
address: "192.168.0.100"
port: "8080"
```

example.tmpl
```
daemon            off;
worker_processes  2;
user              {{ .user }};

error_log         {{ .error_log }} info;

http {
    server_tokens off;
    include       mime.types;
    charset       utf-8;

    access_log    logs/access.log  combined;

    server {
        server_name   localhost;
        listen        {{ .address }}:{{ .port }};

        error_page    500 502 503 504  /50x.html;

        location      / {
            root      html;
        }

    }
}
```

``` bash
ubuntu@:~$ gotmplcli --values example-values.yaml --template example.tmpl
daemon            off;
worker_processes  2;
user              www-data;

error_log         logs/error.log info;

http {
    server_tokens off;
    include       mime.types;
    charset       utf-8;

    access_log    logs/access.log  combined;

    server {
        server_name   localhost;
        listen        192.168.0.100:8080;

        error_page    500 502 503 504  /50x.html;

        location      / {
            root      html;
        }

    }
}
```

# Installing

``` bash
wget https://github.com/zbblanton/gotmplcli/releases/download/v0.1.0/gotmplcli-linux-amd64.tar.gz
tar xvzf gotmplcli-linux-amd64.tar.gz
mv gotmplcli /usr/local/bin
```

# Compiling

``` bash
go build
```