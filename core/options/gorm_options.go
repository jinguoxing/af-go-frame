package options

import (
    "github.com/jinguoxing/af-go-frame/core/store/orm"
    "gorm.io/gorm"
    "time"
)

// MySQLOptions defines options for mysql database.
type MySQLOptions struct {
    Host                  string        `json:"host,omitempty"                     mapstructure:"host"`
    Username              string        `json:"username,omitempty"                 mapstructure:"username"`
    Password              string        `json:"-"                                  mapstructure:"password"`
    Database              string        `json:"database"                           mapstructure:"database"`
    MaxIdleConnections    int           `json:"max-idle-connections,omitempty"     mapstructure:"max-idle-connections"`
    MaxOpenConnections    int           `json:"max-open-connections,omitempty"     mapstructure:"max-open-connections"`
    MaxConnectionLifeTime time.Duration `json:"max-connection-life-time,omitempty" mapstructure:"max-connection-life-time"`
    LogLevel              int           `json:"log-level"                          mapstructure:"log-level"`
}



// NewClient create mysql store with the given config.
func (o *MySQLOptions) NewClient() (*gorm.DB, error) {
    opts := &orm.Options{
        Host:                  o.Host,
        Username:              o.Username,
        Password:              o.Password,
        Database:              o.Database,
        MaxIdleConnections:    o.MaxIdleConnections,
        MaxOpenConnections:    o.MaxOpenConnections,
        MaxConnectionLifeTime: o.MaxConnectionLifeTime,
        LogLevel:              o.LogLevel,
    }

    return orm.New(opts)
}