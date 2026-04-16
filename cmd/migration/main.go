package migration

import (
	"fmt"
	"github.com/alireza0/s-ui/config"
	"github.com/alireza0/s-ui/database"
	"github.com/alireza0/s-ui/database/model"
	"log"
)

func MigrateDb() {
	// void running on first install
	db := database.GetDB()

	var err error
	tx := db.Begin()
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()
	currentVersion := config.GetVersion()
	dbVersion := ""
	err = tx.Model(&model.Setting{}).
		Select("value").
		Where("`key` = ?", "version").
		Scan(&dbVersion).Error

	if err != nil {
		log.Printf("Failed to get DB version: %v", err)
	}
	fmt.Println("Current version:", currentVersion, "\nDatabase version:", dbVersion)

	if currentVersion == dbVersion {
		fmt.Println("Database is up to date, no need to migrate")
		return
	}

	fmt.Println("Start migrating database...")

	// Before 1.2
	if dbVersion == "" {
		err = to1_1(tx)
		if err != nil {
			log.Fatal("Migration to 1.1 failed: ", err)
			return
		}
		err = to1_2(tx)
		if err != nil {
			log.Fatal("Migration to 1.2 failed: ", err)
			return
		}
		dbVersion = "1.2"
	}

	// Before 1.3
	if dbVersion[0:3] == "1.2" {
		err = to1_3(tx)
		if err != nil {
			log.Fatal("Migration to 1.3 failed: ", err)
			return
		}
	}

	// Set version
	err = tx.Model(&model.Setting{}).
		Where("`key` = ?", "version").
		Update("value", currentVersion).Error
	if err != nil {
		log.Fatal("Update version failed: ", err)
		return
	}
	fmt.Println("Migration done!")
}
