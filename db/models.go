package db

type Movie struct {
	ID          int    `gorm:"primary key" json:"id"`
	Title       string `binding:"required" json:"title"`
	Description string `binding:"required" json:"description"`
	Year        int    `binding:"required" json:"year"`
}
