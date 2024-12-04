package dao

import (
	"errors"

	"git.ana/xjtuana/api-micro-mail/model"
)

func (d *DAO) FindMailByID(id string) (*model.Mail, error) {
	data := &model.Mail{}
	if err := d.db.Model(data).Where("id = ?", id).First(data).Error; err != nil {
		return nil, err
	}
	data.Template = new(model.Template)
	if err := d.db.Model(data.Template).Where("id = ?", data.TmplID).First(data.Template).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (d *DAO) InsertMail(data *model.Mail) error {
	return d.db.Create(data).Error
}

func (d *DAO) UpdateMail(data *model.Mail) error {
	tx := d.db.Where("id = ?", data.ID).Omit("pk").Updates(data)
	if err := tx.Error; err != nil {
		return err
	}
	if tx.RowsAffected == 0 {
		return errors.New("nothing updated")
	}
	return nil
}

func (d *DAO) DeleteMail(id string) error {
	return d.db.Model(&model.Mail{}).Delete("id = ?", id).Error
}
