package options

import (
    "github.com/jinguoxing/af-go-frame/core/store/gormx"
    "github.com/jinguoxing/af-go-frame/core/store/orm"
    "gorm.io/gorm"
    "time"
)

// DBOptions defines options for  database.
type DBOptions struct {
    DBType                string        `json:"dbtype,omitempty"                   mapstructure:"db-type"`
    Host                  string        `json:"host,omitempty"                     mapstructure:"host"`
    Username              string        `json:"username,omitempty"                 mapstructure:"username"`
    Password              string        `json:"-"                                  mapstructure:"password"`
    Database              string        `json:"database"                           mapstructure:"database"`
    MaxIdleConnections    int           `json:"max-idle-connections,omitempty"     mapstructure:"max-idle-connections"`
    MaxOpenConnections    int           `json:"max-open-connections,omitempty"     mapstructure:"max-open-connections"`
    MaxConnectionLifeTime time.Duration `json:"max-connection-life-time,omitempty" mapstructure:"max-connection-life-time"`
    LogLevel              int           `json:"log-level"                          mapstructure:"log-level"`
    IsDebug               bool          `json:"isdebug"                           mapstructure:"is-debug"`
    TablePrefix           string        `json:"tableprefix"                       mapstructure:"table-prefix"`

}


// NewClient create mysql store with the given config.
func (o *DBOptions) NewMySqlClient() (*gorm.DB, error) {
    opts := &gormx.Options{
        DBType:     "mysql",
        Host:                  o.Host,
        Username:              o.Username,
        Password:              o.Password,
        Database:              o.Database,
        MaxIdleConnections:    o.MaxIdleConnections,
        MaxOpenConnections:    o.MaxOpenConnections,
        MaxConnectionLifeTime: o.MaxConnectionLifeTime,
        LogLevel:              o.LogLevel,
    }

    return gormx.New(opts)
}