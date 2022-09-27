package orm

import (
    "af-go-frame/core/store/mysql"
    "gorm.io/gorm"
)

type (
    OrmDBConf struct {



        // DB的基础配置
        *mysql.DBConf
        // gorm的配置
        *gorm.Config
    }

    Option func(conf *OrmDBConf)


)

type ORMOption func(config *OrmDBConf) error

type ORMService interface {
    GetDB(option ...ORMOption)
}


