GoAPI is a RESTful API, written in Go.


# Table of Contents

- [About](#about)
- [Requirements](#requirements)
- [Recommended](#recommended)
- [Getting Started](#getting-started)
  - [Basic Authentication](#basic-authentication)
  - [Database](#database)
  - [Server](#server)
  - [Daemonize](#daemonize)
- [Licence](#licence)

# About

GoAPI was built to support the GoDo application, in providing a RESTful API.

# Requirements

- Go (compile the web app for your server)
- PostgreSQL

# Recommended

These recommendations are only due to my setup. As GoAPI can easily be setup using other servers or init systems.

- Apache2 (ProxyPass)
- Systemd (daemonize the GoAPI app)

# Getting Started

## Basic Authentication

Then change the username and password vairables, required to access the API endpoints. These variables are found at the top of the `main.go` file ("username", "password").

```go
...
var (
    username  = hasher("username")
    password  = hasher("password")
    realm     = "Please enter your username and password to gain access to this API"
)
...
```

## Database

To setup the server, make sure to edit the `dsn` variable to match the user, password and table of your postgres database. This is found in the `model/model.go` file.

```go
...
func InitDB() {
    dsn := "user=user password=password dbname=goapi port=5432"
...
```

## Server

A basic Apache configuration to relay GoAPI to your domain.

Apache Configuration Example:
```
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

## Daemonize

Daemonize GoAPI to run at all times in the background, using your init system or as a cron job.

Systemd Service Example:

```
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

# Licence

GoDo is released under the GNU General Public License v3.0. See [LICENSE](https://github.com/rsHalford/godo/LICENSE)
