package gormx

import (
    "fmt"
    "github.com/go-sql-driver/mysql"
    "gorm.io/driver/mysql"
    "gorm.io/driver/postgres"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "gorm.io/gorm/schema"
    "strings"
    "time"
)


const (
    defaultModelSafe        = false
    defaultCharset          = `utf8mb4`
    defaultMaxIdleConnCount = 10               // Max idle connection count in pool.
    defaultMaxOpenConnCount = 0                // Max open connection count in pool. Default is no limit.
    defaultMaxConnLifeTime  = 30 * time.Second // Max lifetime for per connection in pool in seconds.
)


// New create a new gorm db instance with the given options.
func New(opts *Options) (*gorm.DB, error) {

    var dialector gorm.Dialector

    dsn := fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=%s`,
        opts.Username,
        opts.Password,
        opts.Host,
        opts.Database,
        defaultCharset,
        true,
        "Local")

    switch strings.ToLower(opts.DBType) {
    case "mysql":
        dialector = mysql.Open(dsn)
    case "postgres":
        dialector = postgres.Open(dsn)
    default:
        dialector = sqlite.Open(dsn)

    }

    gconfig := &gorm.Config{
        Logger: opts.Logger,
        NamingStrategy: schema.NamingStrategy{
            TablePrefix:   opts.TablePrefix,
            SingularTable: true,
        },
    }

    db, err := gorm.Open(dialector, gconfig)
    if err != nil {
        return nil, err
    }

    if opts.IsDebug {
        db.Debug()
    }

    sqlDB, err := db.DB()
    if err != nil {
        return nil, err
    }

    // SetMaxOpenConns sets the maximum number of open connections to the database.
    if opts.MaxOpenConnections > 0 {
        sqlDB.SetMaxOpenConns(opts.MaxOpenConnections)
    }else{
        sqlDB.SetMaxOpenConns(defaultMaxOpenConnCount)
    }

    // SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
    if opts.MaxConnectionLifeTime > 0 {
        sqlDB.SetConnMaxLifetime(opts.MaxConnectionLifeTime)
    }else{
        sqlDB.SetConnMaxLifetime(defaultMaxConnLifeTime)
    }
    // SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
   if opts.MaxIdleConnections > 0 {
       sqlDB.SetMaxIdleConns(opts.MaxIdleConnections)
   }else {
       sqlDB.SetMaxIdleConns(defaultMaxIdleConnCount)
   }

    return db, nil
}