module github.com/lokks307/adr-boilerplate

go 1.16

require (
	github.com/friendsofgo/errors v0.9.2
	github.com/karlseguin/ccache/v2 v2.0.8
	github.com/labstack/echo/v4 v4.6.2
	github.com/lokks307/djson/v2 v2.0.4 // indirect
	github.com/lokks307/go-util v1.1.4
	github.com/mattn/go-colorable v0.1.11
	github.com/mattn/go-sqlite3 v1.14.18 // indirect
	github.com/pelletier/go-toml v1.9.4 // indirect
	github.com/sirupsen/logrus v1.8.1
	github.com/volatiletech/null/v8 v8.1.2
	github.com/volatiletech/sqlboiler/v4 v4.8.3
	github.com/volatiletech/strmangle v0.0.4
)

replace (
	github.com/lokks307/adr-boilerplate/action => ./action
	github.com/lokks307/adr-boilerplate/cache => ./cache
	github.com/lokks307/adr-boilerplate/domain => ./domain
	github.com/lokks307/adr-boilerplate/env => ./env
	github.com/lokks307/adr-boilerplate/middleware => ./middleware
	github.com/lokks307/adr-boilerplate/models => ./models
	github.com/lokks307/adr-boilerplate/responder => ./responder
	github.com/lokks307/adr-boilerplate/types/e => ./types/e
)
