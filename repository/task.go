package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type TaskRepository interface {
	Store(task *model.Task) error
	Update(id int, task *model.Task) error
	Delete(id int) error
	GetByID(id int) (*model.Task, error)
	GetList() ([]model.Task, error)
	GetTaskCategory(id int) ([]model.TaskCategory, error)
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepo(db *gorm.DB) *taskRepository {
	return &taskRepository{db}
}

func (t *taskRepository) Store(task *model.Task) error {
	err := t.db.Create(task).Error
	if err != nil {
		return err
	}

	return nil
}

func (t *taskRepository) Update(id int, task *model.Task) error {
	if err := t.db.Save(task).Error; err != nil {
		return err
	}

	return nil // TODO: replace this
}

func (t *taskRepository) Delete(id int) error {
	result := t.db.Delete(&model.Task{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil // TODO: replace this
}

func (t *taskRepository) GetByID(id int) (*model.Task, error) {
	var task model.Task
	err := t.db.First(&task, id).Error
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (t *taskRepository) GetList() ([]model.Task, error) {
	var tasks []model.Task
	if err := t.db.Find(&tasks).Error; err != nil {
		return []model.Task{}, err
	}
	return tasks, nil // TODO: replace this
}

func (t *taskRepository) GetTaskCategory(id int) ([]model.TaskCategory, error) {
	var taskCategories []model.TaskCategory
	if err := t.db.Table("tasks").Select("tasks.id, tasks.title, categories.name AS category").
			  Joins("JOIN categories ON tasks.category_id = categories.id").Where("tasks.id = ?", id).
			  Scan(&taskCategories).Error; err!= nil {
				
			return nil, err
	}
	
	return taskCategories, nil // TODO: replace this
}
