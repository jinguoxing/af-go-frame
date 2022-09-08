package sqlConfig


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
    ConnMaxIdle     int    `json:"conn_max_idle"`     // 最大空闲连接数
    ConnMaxOpen     int    `json:"conn_max_open"`     // 最大连接数
    ConnMaxLifetime string `json:"conn_max_lifetime"` // 连接最大生命周期
    ConnMaxIdletime string `json:"conn_max_idletime"` // 空闲最大生命周期

}
