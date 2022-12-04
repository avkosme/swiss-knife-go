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

userList := user.FindByQuery(rep.Db, gift.SqlQuery.String, limit, offset)    
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate if exist.


## License

[MIT](https://choosealicense.com/licenses/mit/)