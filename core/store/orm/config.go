package orm

import (
    "af-go-frame/core/store/sqlConfig"
    "gorm.io/gorm"
)

type OrmDBConf struct {

    *sqlConfig.DBConf

    *gorm.Config

}

