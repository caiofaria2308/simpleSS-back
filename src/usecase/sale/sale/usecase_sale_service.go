package usecase_sale

import entity "main/entity/sale"

type UseCaseSale struct {
	repo IRepositorySale
}

func NewService(repo IRepositorySale) *UseCaseSale {
	return &UseCaseSale{repo: repo}
}

func (s *UseCaseSale) GetAll(store_id string) ([]*entity.EntitySale, error) {
	return s.repo.GetAll(store_id)
}

func (s *UseCaseSale) GetByID(store_id string, id string) (*entity.EntitySale, error) {
	return s.repo.GetByID(store_id, id)
}

func (s *UseCaseSale) FilterByNumber(store_id string, number string) (sales []*entity.EntitySale, err error) {
	return s.repo.FilterByNumber(store_id, number)
}

func (s *UseCaseSale) FilterByDate(store_id string, initialDate string, endDate string) (sales []*entity.EntitySale, err error) {
	return s.repo.FilterByDate(store_id, initialDate, endDate)
}

func (s *UseCaseSale) FilterByEmployee(store_id string, employee_id string) (sales []*entity.EntitySale, err error) {
	return s.repo.FilterByEmployee(store_id, employee_id)
}

func (s *UseCaseSale) FilterByClient(store_id string, client_id string) (sales []*entity.EntitySale, err error) {
	return s.repo.FilterByClient(store_id, client_id)
}

func (s *UseCaseSale) Create(sale entity.EntitySale) (*entity.EntitySale, error) {
	err := sale.Validate()
	err = entity.CreateSale(&sale)
	if err != nil {
		return nil, err
	}
	return s.repo.Create(sale)
}

func (s *UseCaseSale) Update(sale entity.EntitySale) (*entity.EntitySale, error) {
	err := sale.Validate()
	err = entity.UpdateSale(&sale)
	if err != nil {
		return nil, err
	}
	return s.repo.Update(sale)
}

func (s *UseCaseSale) Delete(store_id string, id string) error {
	return s.repo.Delete(store_id, id)
}
