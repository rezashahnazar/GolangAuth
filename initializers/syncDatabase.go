package initializers

import "github.com/rezashahnazar/GolangAuth/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
