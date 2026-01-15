package structs

type UserInfoRequest struct {
	UserID string `json:"user_id"`
}

type CreateUserRequest struct {
	EmailID string `json:"email_id"`
	Name    string `json:"name"`
	Role    string `json:"role"`
}

type UserInfoResponse struct {
	UserID  string `json:"user_id"`
	EmailID string `json:"email_id"`
	Name    string `json:"name"`
	Role    string `json:"role"`
}

type User struct {
	UserID  string `gorm:"primaryKey;column:user_id" json:"user_id"`
	EmailID string `gorm:"column:email_id;uniqueIndex;not null" json:"email_id"`
	Name    string `gorm:"column:name;not null" json:"name"`
	Role    string `gorm:"column:role;not null" json:"role"`
}

// TableName specifies the table name for User model
func (User) TableName() string {
	return "users"
}
