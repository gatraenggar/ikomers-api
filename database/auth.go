package database

type Auth struct {
	RefreshToken string `gorm:"primaryKey, size:512"`
}
