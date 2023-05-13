package postgres

import (
	"os"
	"reflect"
	"self-payrol/config/seed"
	"self-payrol/model"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitGorm initializes a gorm DB object with a postgres connection
func InitGorm() *gorm.DB {
	connection := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(connection))
	if err != nil {
		log.Error().Msgf("cant connect to database %s", err)
	}

	err = db.AutoMigrate(&model.Position{}, &model.User{}, &model.Company{}, &model.Transaction{})
	if err != nil {
		log.Error().Msgf("failed to migrate %s", err)
	}

	SeedData(db,
		seed.CompanySeed,
		seed.PositionSeed,
		seed.UserSeed)

	return db
}

// SeedData inserts initial data into the database
func SeedData(DB *gorm.DB, seeds ...interface{}) {
	for _, seed := range seeds {
		var count int64
		firstData := reflect.ValueOf(seed).Index(0).Interface()
		if DB.Migrator().HasTable(firstData) {
			DB.Model(firstData).Count(&count)
			if count < 1 {
				DB.CreateInBatches(seed, reflect.ValueOf(seed).Len())
			}
		}
	}
}
