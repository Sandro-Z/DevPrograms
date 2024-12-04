package dao

import (
	"errors"

	"git.ana/xjtuana/api-micro-mail/model"
)

func (d *DAO) FindTemplateByID(id string) (*model.Template, error) {
	data := &model.Template{}
	if err := d.db.Model(data).Where("id = ?", id).First(data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (d *DAO) InsertTemplate(data *model.Template) error {
	return d.db.Create(data).Error
}

func (d *DAO) UpdateTemplate(data *model.Template) error {
	tx := d.db.Where("id = ?", data.ID).Omit("pk").Updates(data)
	if err := tx.Error; err != nil {
		return err
	}
	if tx.RowsAffected == 0 {
		return errors.New("nothing updated")
	}
	return nil
}

func (d *DAO) DeleteTemplate(id string) error {
	return d.db.Model(&model.Template{}).Delete("id = ?", id).Error
}
