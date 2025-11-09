package models

type GirlModel struct {
	ID      int
	Name    string     `gorm:"size:32; not null"`
	BoyList []BoyModel `gorm:"foreignKey:GirlId"`
}

type BoyModel struct {
	ID        int
	Name      string `gorm:"size:32; not null"`
	GirlId    int
	GirlModel GirlModel `gorm:"foreignKey:GirlId"`
}
