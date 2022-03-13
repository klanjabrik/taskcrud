// models/task.go

package models

import (
	"time"
)

type Task struct {
	ID         uint      `json:"id" gorm:"primary_key"`
	AssingedTo string    `json:"assignedTo"`
	Task       string    `json:"task"`
	UserId     uint      `gorm:"index:" sql:"DEFAULT:0" json:"user_id" `
	Deadline   time.Time `json:"deadline"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func GetAllTask(t *[]Task, m map[string]interface{}) (err error) {
	if err := SetupDB().Where(m).Find(t).Error; err != nil {
		return err
	}
	return nil
}

func AddNewTask(t *Task) (err error) {
	if err = SetupDB().Create(t).Error; err != nil {
		return err
	}
	return nil
}

func GetOneTaskId(t *Task, id string) (err error) {
	m := map[string]interface{}{"id": id}
	if err := GetOneTask(t, m); err != nil {
		return err
	}
	return nil
}

func GetOneTask(t *Task, m map[string]interface{}) (err error) {
	if err := SetupDB().Where(m).First(t).Error; err != nil {
		return err
	}
	return nil
}

func PutOneTask(t *Task, tupdate *Task) (err error) {
	if err := SetupDB().Model(&t).Updates(tupdate).Error; err != nil {
		return err
	}
	return nil
}

func DeleteTask(t *Task) (err error) {
	if err := SetupDB().Delete(&t).Error; err != nil {
		return err
	}
	return nil
}
