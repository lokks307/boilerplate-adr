module github.com/lokks307/adr-boilerplate

go 1.20

require (
	github.com/friendsofgo/errors v0.9.2
	github.com/karlseguin/ccache/v2 v2.0.8
	github.com/labstack/echo/v4 v4.6.2
	github.com/lokks307/djson/v2 v2.0.7
	github.com/mattn/go-colorable v0.1.11
	github.com/pelletier/go-toml v1.9.4
	github.com/sirupsen/logrus v1.8.1
	github.com/volatiletech/null/v8 v8.1.2
	github.com/volatiletech/sqlboiler/v4 v4.8.3
	github.com/volatiletech/strmangle v0.0.4
)

require (
	github.com/asaskevich/govalidator v0.0.0-20230301143203-a9d515a09cc2 // indirect
	github.com/denisenkom/go-mssqldb v0.10.0 // indirect
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/gofrs/uuid v4.4.0+incompatible // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/golang-sql/civil v0.0.0-20190719163853-cb61b32ac6fe // indirect
	github.com/labstack/gommon v0.3.1 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/spf13/cast v1.4.1 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.1 // indirect
	github.com/volatiletech/inflect v0.0.1 // indirect
	github.com/volatiletech/randomize v0.0.1 // indirect
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5 // indirect
	golang.org/x/net v0.0.0-20210913180222-943fd674d43e // indirect
	golang.org/x/sys v0.0.0-20211103235746-7861aae1554b // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/time v0.0.0-20201208040808-7e3f01d25324 // indirect
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
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
