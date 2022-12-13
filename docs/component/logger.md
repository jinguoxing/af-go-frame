# 日志
日志模块借鉴了[marmotedu/log](https://github.com/marmotedu/iam/tree/master/pkg/log), 在此基础上改造成了我们想要样子。目录结构、功能参考了[gogf/gf](https://github.com/gogf/gf)



## 设计理念
1. zap日志库为核心
2. 使用配置文件读取配置初始化
3. 日志分级输出, 支持多个等级输出

## 基本使用
基本使用和绝大部分日志库类似,

### 默认Logger
模块有默认的Logger, 用户可以不用自己创建实例. `zapx.DefaultLogger()`可以获取默认日志实例, 默认日志实例只会将日志打印到标准输出。

### 初始化
有时候想要将日志输出到文件, 需要读取项目的配置文件

#### 配置文件
```yaml
logs:
  - name: ago                                   #日志名称
    default: true                               #默认日志实例
    enable_caller: true                         #是否输出日志调用方
    stacktrace_level:                           #输出堆栈的最低日志等级
    stack_filter:                               #堆栈过滤字段
      - runtime
    development: true                           #是否开启development模式
    cores:                                      #日志输出的流
      - destination: infos.log                  #输出的文件名
        rotate_size: 10MB                       #文件分割的大小
        core_type: file                         #日志流的类型, file, stdout, 其他类型用户可以自己注册
        enable_color: false                     #是否开启输出颜色
        output_format: json                     #日志输出的格式, json, console
        filename_format: "-%Y-%m-%d.log"        #文件分割的名称，日期格式，下面的是默认格式
        log_level: info                         #该日志流的格式
```


#### 配置结构体
```go
type LogConfigs struct {
    Logs []Options `json:"logs"`
}

//Options single logger config
type Options struct {
    Name            string       `json:"name"`               
    local           boo
    Default         bool         `json:"default"`            
    EnableCaller    bool         `json:"enable_caller"`      
    StacktraceLevel string       `json:"stacktrace_level"`   
    Development     bool         `json:"development"`        
    StackFilter     []string     `json:"stack_filter"`  		
    CoreConfigs     []CoreConfig `json:"cores"`              
}

//CoreConfig zap log core config
type CoreConfig struct {
    RotateSize     string `json:"rotate_size"       mapstructure:"rotate_size"`         //file rotate size
    Destination    string `json:"destination"       mapstructure:"destination"`         //log output destination
    CoreType       string `json:"core_type"         mapstructure:"core_type"`             //core writeSyncer type, 'file', 'stdout'
    EnableColor    bool   `json:"enable_color"      mapstructure:"enable_color"`       //enable color
    OutputFormat   string `json:"output_format"     mapstructure:"output_format"`     //log line output format: 'line', 'json'
    FileNameFormat string `json:"filename_format"   mapstructure:"filename_format"` //name format in file rotate
    LogLevel       string `json:"log_level"         mapstructure:"log_level"`             //lowest log level in this core
}
```

#### 实例化
初始化使用泛型特性, 更加的直观
```go
logs := config.Scan[zapx.LogConfigs]()
zapx.Loads(logs)
```

### 打印日志
```go
zapx.Debug("This is a debug message")
zapx.Info("This is a info message")
zapx.Warn("This is a formatted %s message")
zapx.Error("Message printed with Errorw")
zapx.Panic("This is a panic message")
```

在日志中添加字段或者格式化
```go
zapx.Debugf("This is a %s message", "debug")
zapx.Infof("This is a %s message", "info")
zapx.Warn("This is a formatted %s message", zapx.Int32("int_key", 10))
zapx.Errorw("Message printed with Errorw", "X-Request-ID", "fbf54504-64da-4088-9b86-67824a7fb508")
zapx.Panic("This is a panic message")
```

## 功能配置

### 绑定context
设置context，下面的logger实例绑定context后, 将会一直有requestID字段
```go
logger := zapx.GetLogger('ln')
ctx := context.Background()
context.WithValue(ctx, "requestID", "RID:123456789156487")
logger.SetLevel(zapx.InfoLevel).Info("This is a V level message")
```

### 日志级别
```go
logger.SetLevel(zapx.InfoLevel).Info("This is a V level message")
```

### 日志配置
```go
logger.Caller(true)  //是否输出caller
ln.Development() // 是否开启development模式
```

### 打印堆栈
```go
logger.PrintStack()
zapx.PrintStack()
```

## 高级特性

### 自定义WriteSyncer
如果想要将日志发送到云或者第三方分析平台,可以自定义`WriteSyncer`
```go
//WriteSyncer 自定义的WriteSyncer, 如果需要其他的参数,可以在配置中增加
func WriteSyncer() zapcore.WriteSyncer {
	// check file path, create file if not exists
	destination := pwd + "/error.log"
	// set filename format
	fileNameFormat := zapx.DefaultFileNameFormat

	hook, _ := rotatelogs.New(
		strings.Replace(destination, ".log", "", -1)+fileNameFormat,
		rotatelogs.WithLinkName(destination),
		rotatelogs.WithRotationSize(10*zapx.MB),
	)
	return zapcore.AddSync(hook)
}

func main(){
	// 注册方法必须在Loads之前
    zapx.RegisterWriteSyncer("test", WriteSyncer())
    InitSources(fmt.Sprintf("%s/config.yaml", pwd))
    // 初始化全局logger
    defer zapx.Flush()
    
    logs := config.Scan[zapx.LogConfigs]()
    zapx.Loads(logs)
}


```

### 注册hook
添加钩子函数, 适合遇到特定的日志发送邮件,统计等操作
```go
zapx.GetLogger().RegisterHook(zapx.SampleHook)
```


