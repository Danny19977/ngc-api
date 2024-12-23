package database

import (
    "gorm.io/drm"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
    "project/pkg/config"
)

func Initialize() (*gorm.DB, error) {
    dbURL := config.GetConfig().DatabaseURL
    db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info),
    })
    if err != nil {
        return nil, err
    }
    return db, nil
}
