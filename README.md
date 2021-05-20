# GoAPI

## GoAPI was built to support the GoDo application, in providing a RESTful API.

---

# Table of Contents

- [Getting Started](#getting-started)
  - [Requirements](#requirements)
  - [Recommendations](#recommendations)
    - [Basic Authentication](#basic-authentication)
    - [Database](#database)
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

Then after making the [necessary changes](#basic-authentication) to the source code. Build the GoAPI binary, for the operating server it will be ran on.

```sh
$ env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build github.com/rsHalford/goapi
```

> This example command will build GoAPI to be executable on Debian 10.

Then you would need to have a web server installed and a way to run GoAPI as a daemon.

- Apache2 (ProxyPass to your web address)
- Systemd (daemonize GoAPI to run in the background)

> These recommendations are based on what I found to work best for my setup.

## Recommendations

To setup the server application, there a four changes that need to be made.

### Basic Authentication

To secure your todo list online, you will need to change the `username` and `password` variables, required to access the API endpoints. These variables are found at the top of the `main.go` file.

```go
...
var (
    username  = hasher("username")
    password  = hasher("password")
    realm     = "Please enter your username and password to gain access to this API"
)
...
```

### Database

Currently GoAPI only supports PostgreSQL. To link up the server to a database, make sure to edit the `dsn` variable to provide the `user`, `password` and `dbname` relating to your database address. This is found in the `model/model.go` file.

```go
...
func InitDB() {
    dsn := "user=user password=password dbname=goapi port=5432"
...
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

👉 See [LICENSE](https://github.com/rsHalford/godo/blob/main/LICENSE).
