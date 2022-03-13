// models/user.go

package models

import (
	"time"
)

type User struct {
	ID         uint      `json:"id" gorm:"primary_key"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Email      string    `json:"email"`
	EmailToken string    `json:"email_token"`
	Token      string    `json:"-"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func GetAllUser(u *[]User, m map[string]interface{}) (err error) {
	if err = SetupDB().Where(m).Find(u).Error; err != nil {
		return err
	}
	return nil
}

func AddNewUser(u *User) (err error) {
	if err = SetupDB().Create(u).Error; err != nil {
		return err
	}
	return nil
}

func GetOneUserId(u *User, id string) (err error) {
	m := map[string]interface{}{"id": id}
	if err := GetOneUser(u, m); err != nil {
		return err
	}
	return nil
}

func GetOneUser(u *User, m map[string]interface{}) (err error) {
	if err := SetupDB().Where(m).First(u).Error; err != nil {
		return err
	}
	return nil
}

func PutOneUser(u *User, uupdate map[string]interface{}) (err error) {
	if err := SetupDB().Model(&u).Updates(uupdate).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUser(u *User) (err error) {
	if err := SetupDB().Delete(&u).Error; err != nil {
		return err
	}
	return nil
}
