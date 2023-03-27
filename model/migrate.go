package model

func migration() {
	DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate()
}
