package entity

type Users struct {
	Username string `json:"username" bson:"Username" binding:"required"`
	Password string `json:"password" bson:"Password" binding:"required"`
	Role     string `json:"role" bson:"Role"`
}
