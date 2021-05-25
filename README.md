# GoAPI

## GoAPI was built to support the GoDo application, in providing a RESTful API.

[![Go Reference](https://pkg.go.dev/badge/github.com/rsHalford/goapi.svg)](https://pkg.go.dev/github.com/rsHalford/goapi)

---

# Table of Contents

- [Getting Started](#getting-started)
  - [Requirements](#requirements)
  - [Setting Up](#setting-up)
    - [Authentication](#authentication)
    - [Server](#server)
    - [Daemonize](#daemonize)
- [Licence](#licence)

---

# Getting Started

## Requirements

- Go (to compile applications)
- PostgreSQL

To do this, all you need is to have Go - [the programming language](https://golang.org/doc/install) - installed on your computer.

To edit the necessary variables to hook up GoAPI with a database. You will need to clone the GoAPI repository.

```sh
$ git clone github.com/rsHalford/goapi
```

Then after making the [necessary changes](#setting-up) to the source code. Build the GoAPI binary, for the operating server it will be ran on.

```sh
$ env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build github.com/rsHalford/goapi
```

> This example command will build GoAPI to be executable on Debian 10.

Then you would need to have a web server installed and a way to run GoAPI as a daemon.

- Apache2 (ProxyPass to your web address)
- Systemd (daemonize GoAPI to run in the background)

> These recommendations are based on what I found to work best for my setup.

## Setting Up

To setup the server application, there a three changes that need to be made.

### Authentication

To secure your todo list online, you will need to change the `api_username` and `api_password` variables, required to access the API endpoints. These variables are found in the `config.yaml` file.

These values will be what you send to the API with each request, using Basic Authentication.

> An example of the `config.yaml` can be found in the [project's repository](https://github.com/rsHalford/goapi/blob/main/config.yaml)

```yaml
api_username: "username"
api_password: "password"
```

Currently GoAPI only supports PostgreSQL. To link up the server to a database, make sure to edit the `config.yaml`. Providing the `db_username`, `db_password`, `name` and `port` - relating to your database address.

```yaml
db_username: "username"
db_password: "password"
name: "goapi"
port: "5432"
```

### Server

GoAPI by default serves on port `:8080`. As an example, below is a basic Apache configuration file to relay GoAPI to your domain.

Apache Configuration Example:

```apache
<VirtualHost *:80>
        ServerName goapi.example.com
        ServerAdmin webmaster@example.com
        ProxyRequests Off
        <Proxy *>
                Require all granted
        </Proxy>
        ProxyPass / http://127.0.0.1:8080/
        ProxyPassReverse / http:127.0.0.1:8080/
        ErrorLog ${APACHE_LOG_DIR}/goapi-error.log
        CustomLog ${APACHE_LOG_DIR}/goapi-access.log combined
</VirtualHost>
```

### Daemonize

To have GoAPI run at all times in the background. You will need to make it run as a daemon. This is possible by creating one as a service with your init system - such as with systemd.

Systemd Service Example:

```toml
[Unit]
Description=GoAPI RESTful API

[Service]
ExecStart=/path/to/goapi/excutable
WorkingDirectory=/path/to/goapi
StandardOutput=journal
StandardError=inherit
SyslogIdentifier=goapi
User=user
Group=group
Type=simple
Restart=on-failure

[Install]
WantedBy=multi-user.target
```

---

# Licence

GoAPI is released under the GNU General Public License v3.0.

👉 See [LICENSE](https://github.com/rsHalford/goapi/blob/main/LICENSE).
