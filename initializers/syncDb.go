package initializers

import "github.com/bayunashr/gopoll/models"

func SyncDb() {
	DB.AutoMigrate(&models.User{})       // migrate tabel user
	DB.AutoMigrate(&models.Poll{})       // migrate tabel poll
	DB.AutoMigrate(&models.PollChoice{}) // migrate tabel pollchoice
	DB.AutoMigrate(&models.PollEntry{})  // migrate tabel pollentry
}
