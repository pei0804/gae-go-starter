package model

type Comments struct {
	BaseModel
	Name    string `gorm:"type:varchar(30)" json:"name"`
	Comment string `gorm:"type:text" json:"comment"`
}
