package mysql

import (
    "github.com/go-sql-driver/mysql"
    "net"
    "strconv"
    "time"
)

// 关系型数据库的配置
type DBConf struct {

    // 以下配置关于dsn
    WriteTimeout string `json:"write_timeout"` // 写超时时间
    Loc          string `json:"loc"`           // 时区
    Port         int    `json:"port"`          // 端口
    ReadTimeout  string `json:"read_timeout"`  // 读超时时间
    Charset      string `json:"charset"`       // 字符集
    ParseTime    bool   `json:"parse_time"`    // 是否解析时间
    Protocol     string `json:"protocol"`      // 传输协议
    Dsn          string `json:"dsn"`           // 直接传递dsn，如果传递了，其他关于dsn的配置均无效
    Database     string `json:"database"`      // 数据库
    Collation    string `json:"collation"`     // 字符序
    Timeout      string `json:"timeout"`       // 连接超时时间
    Username     string `json:"username"`      // 用户名
    Password     string `json:"password"`      // 密码
    Driver       string `json:"driver"`        // 驱动
    Host         string `json:"host"`          // 数据库地址
    AllowNativePasswords bool `json:"allow_native_passwords"` // 是否允许nativePassword

    // 以下配置关于连接池
    MaxIdleConnections     int    `json:"max-idle-connections,omitempty"`     // 最大空闲连接数
    MaxOpenConnections     int    `json:"max-open-connections,omitempty"`     // 最大连接数
    MaxConnectionLifeTime time.Duration `json:"max-connection-life-time,omitempty"` // 连接最大生命周期
    MaxConnectionIdletime time.Duration `json:"conn_max_idletime"` // 空闲最大生命周期

}

// NewMySQLOptions create a `zero` value instance.
func NewMySQLOptions() *DBConf{
    return &DBConf{
        Host:                  "127.0.0.1",
        Port:                   3306,
        Username:              "",
        Password:              "",
        Database:              "",
        MaxIdleConnections:    30,
        MaxOpenConnections:    30,
        MaxConnectionLifeTime: time.Duration(10) * time.Second,
    }
}



func (conf *DBConf) FormatDSN() (string, error) {

    port := strconv.Itoa(conf.Port)

    timeout ,err := time.ParseDuration(conf.Timeout)

    readTimeout, err := time.ParseDuration(conf.ReadTimeout)
    if err != nil {
        return "", err
    }
    writeTimeout, err := time.ParseDuration(conf.WriteTimeout)
    if err != nil {
        return "", err
    }
    location, err := time.LoadLocation(conf.Loc)
    if err != nil {
        return "", err
    }

    driverConf := &mysql.Config{
        User:                 conf.Username,
        Passwd:               conf.Password,
        Net:                  conf.Protocol,
        Addr:                 net.JoinHostPort(conf.Host, port),
        DBName:               conf.Database,
        Collation:            conf.Collation,
        Loc:                  location,
        Timeout:              timeout,
        ReadTimeout:          readTimeout,
        WriteTimeout:         writeTimeout,
        ParseTime:            conf.ParseTime,
        AllowNativePasswords: conf.AllowNativePasswords,
    }
    return driverConf.FormatDSN(),nil

}





