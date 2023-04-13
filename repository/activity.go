package repository

import (
	"todo-app/models"

	"gorm.io/gorm"
)

type ActivityRepository interface {
	CreateActivity(ActivityGroup models.ActivityGroup) (models.ActivityGroup, error)
	GetActivity(ActivityGroupId int) (models.ActivityGroup, error)
	UpdateActivity(ActivityGroup models.ActivityGroup) (models.ActivityGroup, error)
	DeleteActivity(activity models.ActivityGroup) (models.ActivityGroup, error)
	FindActivity() ([]models.ActivityGroup, error)
}

func RepositoryActivity(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindActivity() ([]models.ActivityGroup, error) {
	var Activity []models.ActivityGroup
	err := r.db.Raw("SELECT * FROM activity_groups").Scan(&Activity).Error

	return Activity, err
}

func (r *repository) DeleteActivity(activity models.ActivityGroup) (models.ActivityGroup, error) {
	err := r.db.Delete(&activity).Error // Using Delete method

	return activity, err
}

func (r *repository) CreateActivity(ActivityGroup models.ActivityGroup) (models.ActivityGroup, error) {
	err := r.db.Create(&ActivityGroup).Error
	return ActivityGroup, err
}
func (r *repository) UpdateActivity(ActivityGroup models.ActivityGroup) (models.ActivityGroup, error) {
	err := r.db.Save(&ActivityGroup).Error
	return ActivityGroup, err
}

func (r *repository) GetActivity(ActivityGroupId int) (models.ActivityGroup, error) {
	var ActivityGroup models.ActivityGroup
	err := r.db.Preload("Todo").First(&ActivityGroup, ActivityGroupId).Error
	return ActivityGroup, err
}
