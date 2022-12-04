# The Swiss knife for your golang app

## Usage

### Mysql

```go
import (
    "example/internal/config"
    "github.com/avkosme/swiss-knife-go/mysql"

    "github.com/getsentry/sentry-go"
)

err := sentry.Init(sentry.ClientOptions{
		Dsn: config.SentryDsn,
})
db := mysql.NewMysql(
		config.MysqlUser,
		config.MysqlPassword,
		config.MysqlAddr,
		config.MysqlDatabase,
)
db.Connect()

// []map[string]any []map[<feild name>]<value>
result := mysql.FindIn(db.Db, query, err)
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate if exist.


## License

[MIT](https://choosealicense.com/licenses/mit/)