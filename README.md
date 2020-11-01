# Kuda Web Server

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Open Issues](https://img.shields.io/github/issues-raw/Thor-x86/kuda)](https://github.com/Thor-x86/kuda/issues)
[![Open Pull Request](https://img.shields.io/github/issues-pr-raw/Thor-x86/kuda)](https://github.com/Thor-x86/kuda/pulls)
[![Unit Test Result](https://img.shields.io/travis/Thor-x86/kuda)](https://travis-ci.org/Thor-x86/kuda)


Fast and concurrent in-memory web server. It compress static files with [gzip](https://en.wikipedia.org/wiki/Gzip), put them into RAM, and serve them as a normal web server. So Dev/Ops don't have to worry about storage speed, just focus on networking matter.

The best use case is serving Single Page Application (SPA) like [React](https://reactjs.org/), [Vue](https://vuejs.org/), and [Angular](https://angular.io/).

Special thanks to [fasthttp](https://github.com/valyala/fasthttp) and contributors for making kuda possible 🤘

## How to use

```
USAGE:
    kuda [arguments] <public_root_directory>

ARGUMENTS:
    --port=...     : TCP Port to be listened (default: "8080")
    --origins=...  : Which domains to be allowed by CORS policy (default: "")
    --port-tls=... : Use this to listen for HTTPS requests (default: "")
    --domain=...   : Required to redirect from http to https (default: "localhost")
    --cert=...     : SSL certificate file, required if "--port-tls" specified
    --key=...      : SSL secret key file, required if "--port-tls" specified
```

## Usage Example

Let's say you're now in a directory with `kuda` executable and `my-app` as your react app project.

```
cd my-app
npm run build
cd ..

./kuda my-app/build --port=8000 --origins=localhost:9000
```

After executed, it will put everything inside `my-app/build` into RAM and serve at port 8000. Beside of that, we're assuming the API backend runs on port 9000. Thus, we have to add it as allowed origins with `--origins=...` flag to prevent [CORS Origin](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS) problem.

### How about SSL?

Now we have same directory as above with `cert.pem` and `secret.key` inside of it.

```
./kuda my-app/build --port=8000 --origins=localhost:9000 --port-tls=8090 --domain=localhost --cert=cert.pem --key=secret.key
```

The `localhost:8000` is considered using HTTP protocol. So every connection coming to that port will be redirected to `https://localhost:8090`.


### Serving Production App

In production environment, assuming you have `www.my-domain.com`, it will be like:

```
sudo ./kuda my-app/build --port=80 --origins=localhost:9000 --port-tls=443 --domain=www.my-domain.com --cert=/etc/ssl/certs/my-domain.pem --key=/etc/ssl/private/my-domain.key
```

## With Docker

Mostly, we deploy web app with docker and kubernetes (for orchestration). In order to do that, use this commands.

```
docker pull kuda:latest
docker run --name kuda-demo --volume my-compiled-webapp:/srv -p 8080:8080 --rm kuda:latest
```

Where `my-compiled-webapp` is your directory with compiled app inside. **NEVER EVER** point project root directory as public directory, otherwise your machine will run out of memory very shortly!

## With Docker Compose

I would recommend you to use Docker Compose instead of plain Docker for sake of maintainability. This is an example of `docker-compose.yml`:

```yml
version: "3"

services:
    kuda:
        image: kuda:latest
        volumes:
            - ./my-compiled-webapp:/srv
        environment:
            KUDA_PUBLIC_DIR: "/srv"
            KUDA_DOMAIN: "localhost"
            KUDA_PORT: "8080"
            KUDA_ORIGINS: ""
            KUDA_PORT_TLS: ""
            KUDA_CERT: ""
            KUDA_KEY: ""
        ports:
            - 8080:8080
```

## Benchmark

I have no enough resource to benchmark Kuda Web Server myself. If you did benchmark and comparation with kuda, please let us know via **issue** section.

## Contribution

Anyone can contribute on this project. We welcome you to Pull Request or Open an Issue. To working on source code, you will require:
- Linux or Mac preferred (Use [WSL](https://docs.microsoft.com/en-us/windows/wsl/install-win10) if you are Windows 10 user or [Cygwin](https://www.cygwin.com/) for older Windows)
- Golang
- Makefile
- NPM (for demo)
- Docker + Docker Compose (optional)

### Makefile Commands

There are things you can do with makefile on this project:
- `make`       -- Compile for all platforms and store them inside "build" directory
- `make test`  -- Run this everytime you want to pull request
- `make clean` -- Removes all compiled files to reduce storage usage
- `make demo`  -- Run demonstration, go to `http://localhost:8080` on browser while running this command

### Security Report

To keep other users' security, please send email to [athaariqa@gmail.com](mailto://athaariqa@gmail.com) instead. **Do not** open issue for security report.
