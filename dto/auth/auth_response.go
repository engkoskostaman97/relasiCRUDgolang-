package authdto

type LoginResponse struct {
	ID    int    `json:"id"`
	Name  string `gorm:"type: varchar(255)" json:"name" from:"name"`
	Email string `gorm:"type: varchar(255)" json:"email" from:"email"`
	Token string `gorm:"type: varchar(255)" json:"token"`
}
