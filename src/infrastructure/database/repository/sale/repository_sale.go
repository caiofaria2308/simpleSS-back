package repository

import (
	entity "main/entity/sale"

	"gorm.io/gorm"
)

type RepositorySale struct {
	DB *gorm.DB
}

func NewRepositorySale(db *gorm.DB) *RepositorySale {
	return &RepositorySale{DB: db}
}

func (r *RepositorySale) Create(sale entity.EntitySale) (*entity.EntitySale, error) {
	err := r.DB.Create(&sale).Error
	if err != nil {
		return nil, err
	}
	return &sale, nil
}

func (r *RepositorySale) GetAll(store_id string) ([]*entity.EntitySale, error) {
	var sales []*entity.EntitySale
	err := r.DB.Where("store_id =?", store_id).Find(&sales).Error
	if err != nil {
		return nil, err
	}
	return sales, nil
}

func (r *RepositorySale) GetByID(store_id string, id string) (*entity.EntitySale, error) {
	var sale entity.EntitySale
	sales, err := r.GetAll(store_id)
	if err != nil {
		return nil, err
	}
	err = r.DB.Where("id =?", store_id, id).Find(&sales).First(&sale).Error
	if err != nil {
		return nil, err
	}
	return &sale, nil
}

func (r *RepositorySale) FilterByNumber(store_id string, number string) (sales []*entity.EntitySale, err error) {
	sales, err = r.GetAll(store_id)
	if err != nil {
		return nil, err
	}
	number = "%" + number + "%"
	err = r.DB.Where("number like ?", number).Find(&sales).Error
	if err != nil {
		return nil, err
	}
	return sales, nil
}

func (r *RepositorySale) FilterByDate(store_id string, initialDate string, endDate string) (sales []*entity.EntitySale, err error) {
	sales, err = r.GetAll(store_id)
	if err != nil {
		return nil, err
	}
	endDate = "%" + endDate + "%"
	err = r.DB.Where("date between %s and %s", initialDate, endDate).Find(&sales).Error
	if err != nil {
		return nil, err
	}
	return sales, nil
}

func (r *RepositorySale) FilterByClient(store_id string, client_id string) (sales []*entity.EntitySale, err error) {
	sales, err = r.GetAll(store_id)
	if err != nil {
		return nil, err
	}
	err = r.DB.Where("client_id =?", client_id).Find(&sales).Error
	if err != nil {
		return nil, err
	}
	return sales, nil
}

func (r *RepositorySale) FilterByEmployee(store_id string, employee_id string) (sales []*entity.EntitySale, err error) {
	sales, err = r.GetAll(store_id)
	if err != nil {
		return nil, err
	}
	err = r.DB.Where("employee_id =?", employee_id).Find(&sales).Error
	if err != nil {
		return nil, err
	}
	return sales, nil
}

func (r *RepositorySale) Update(sale entity.EntitySale) (*entity.EntitySale, error) {
	err := r.DB.Save(&sale).Error
	if err != nil {
		return nil, err
	}
	return &sale, nil
}

func (r *RepositorySale) Delete(store_id string, id string) error {
	var sale entity.EntitySale
	sales, err := r.GetAll(store_id)
	if err != nil {
		return err
	}
	err = r.DB.Where("id =?", store_id, id).Find(&sales).Delete(&sale).Error
	if err != nil {
		return err
	}
	return nil
}
