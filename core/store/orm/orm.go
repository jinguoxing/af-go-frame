package orm

import (
    "gorm.io/gorm"
    "sync"
)


type AFGorm struct {

    dbs map[string]*gorm.DB
    lock *sync.RWMutex
}



func (s *AFGorm) GetDB(option ...ORMOption) (*gorm.DB, error) {

    //dsn :=
    //
    //for _, opt := range option {
    //
    //    if err := opt(config)
    //
    //}

}


