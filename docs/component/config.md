# 配置
配置模块'参考了'哔哩哔哩的Go-Kratos的配置模块, 原始参考文档 > [Go-Kratos](https://go-kratos.dev/docs/component/config/)


## 特性
## 1.支持多种配置源
框架定义了Soruce接口和Watcher接口来适配各种源, 框架内置了`本地文件`和`环境变量`实现

## 2.支持多种配置格式
配置组件复用了encoding中的反序列化逻辑作为配置解析使用。默认支持以下格式的解析
* json
* proto
* xml
* yaml

框架将根据配置文件类型匹配对应的Codec，进行配置文件的解析。您也可以通过实现Codec并用encoding.RegisterCodec方法，将它注册进去，来解析其它格式的配置文件。

## 3.热更新
模块默认支持热更新，当源发生变更时， 模块缓存的值将会变化。 同时提供了回调方法`Observer`， 供用户自定义使用

## 4.配置合并
在config组件中，所有的配置源中的配置（文件）将被逐个读出，分别解析成map，并合并到一个map中去。因此在加载完毕后，不需要再理会配置的文件名，不用文件名来进行查找，而是用内容中的结构来对配置的值进行索引即可。设计和编写配置文件时，请注意各个配置文件中，根层级的key不要重复，否则可能会被覆盖。

举例：

有如下两个配置文件：
```yaml
# 文件1
foo:
  baz: "2"
  biu: "example"
hello:
  a: b
```
```yaml
# 文件2
foo:
  bar: 3
  baz: aaaa
hey:
  good: bad
  qux: quux
```
.Load后,将被合并为以下结构
```yaml
{
  "foo": {
    "baz": "aaaa",
    "bar": 3,
    "biu": "example"
  },
  "hey": {
    "good": "bad",
    "qux": "quux"
  },
  "hello": {
    "a": "b"
  }
}
```
我们可以发现，配置文件的各层级将分别合并，在key冲突时会发生覆盖，而具体的覆盖顺序，会由配置源实现中的读取顺序决定，因此这里重新提醒一下，各个配置文件中，根层级的key不要重复，也不要依赖这个覆盖的特性，从根本上避免不同配置文件的内容互相覆盖造成问题。

在使用时，可以用.Value("foo.bar")直接获取某个字段的值，也可以用.Scan方法来将整个map读进某个结构体中，具体使用方式请看下文。

## 5.支持环境变量
可以环境变量`PROJECT_ENV` 定义配置文件后缀, 比如输入的配置文件是`config.yaml`:
```shell
PROJECT_ENV: dev       实际读取--->   config_dev.yaml
PROJECT_ENV: release   实际读取--->   config_release.yaml
PROJECT_ENV:           实际读取--->   config.yaml
```

 # 使用
我们在kratos的基础上包装了一个struct, 在此,推荐使用新的方法

配置文件
```yaml
logs:
  - name: ago
    cores:
      - destination: stdout
        log_level: info
```

初始化方法
```go
package main

import (
	"fmt"
	"github.com/jinguoxing/af-go-frame/core/config"
	"github.com/jinguoxing/af-go-frame/core/config/env"
	"github.com/jinguoxing/af-go-frame/core/config/file"
	"github.com/jinguoxing/af-go-frame/core/logx/zapx"
	"os"
	"path"
	"runtime"
)

var pwd string = "."

func init() {
	os.Setenv(config.ProjectEnvKey, "DEV")

	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
		pwd = abPath
	}
}

type CoreConfig struct {
	Destination string `json:"destination"` //log output destination
	LogLevel    string `json:"log_level"`   //lowest log level in this core
}

type Options struct {
	Name        string       `json:"name"`
	CoreConfigs []CoreConfig `json:"cores"`
}

type LogConfigs struct {
	Logs []Options `json:"logs"`
}

func InitSources(paths ...string) {
	sources := make([]config.Source, len(paths)+1)
	sources[0] = env.NewSource()
	for i, path := range paths {
		sources[i+1] = file.NewSource(path)
	}
	config.Init(sources...)
}

func main() {
	InitSources("config.yaml")
	logs := config.Scan[LogConfigs]()

	fmt.Printf("logs config %v", logs)

}
```

## 读取环境变量
如果想要自定义环境的前准:
```go
os.Setenv(config.ProjectPrefix,"AFG_")
```
模块会根据该环境变量, 取出















