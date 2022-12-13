
翻译: [English](README.md) | [简体中文](README_zh.md)

# Af-Go-Frame 
af-go-frame是基于Gin的一套轻量级 Go 微服务框架，包含大量微服务相关功能及工具。

## 目标
我们致力于提供完整的微服务研发体验，整合相关框架及工具后，微服务治理相关部分可对整体业务开发周期无感，从而更加聚焦于业务交付。


### 设计原则

* 简单：不过度设计，代码平实简单；
* 通用：通用业务开发所需要的基础库的功能；
* 高效：提高业务迭代的效率；
* 稳定：基础库可测试性高，覆盖率高，有线上实践安全可靠；
* 健壮：通过良好的基础库设计，减少错用；
* 高性能：性能高, 比如为了日志库使用了性能最优的zap日志库
* 扩展性：良好的接口设计，来扩展实现，或者通过新增基础库目录来扩展功能；
* 容错性：为失败设计，大量引入对 SRE 的理解，鲁棒性高；
* 工具链：包含大量工具链


## 特性
* [Config](docs/component/config.md) ：支持多数据源方式，进行配置合并铺平，通过 Atomic 方式支持动态配置；
* [Logger](docs/component/logger.md) ：标准日志接口，可方便集成三方 log 库，并可通过 fluentd 收集日志；


## 快速开始

### 依赖
- [Go>=1.18](https://golang.org/dl/)

我们提供了demo项目如下:
* [af-goframe-demo](https://github.com/jinguoxing/af-goframe-demo)


## 提交
提交信息的结构应该如下所示:
```text
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

提交信息应按照下面的格式:
- fix: simply describe the problem that has been fixed
- feat(log): simple describe of new features
- deps(examples): simple describe the change of the dependency
- break(http): simple describe the reasons for breaking change


## 许可证
Af-Go-frame is MIT licensed. See the [LICENSE](./LICENSE) file for details.