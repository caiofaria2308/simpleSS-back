package usecase_purchaseorder

import entity "main/entity/purchase"

type UseCasePurchaseOrder struct {
	repo IRepositoryPurchaseOrder
}

func NewService(repo IRepositoryPurchaseOrder) *UseCasePurchaseOrder {
	return &UseCasePurchaseOrder{repo: repo}
}

func (s *UseCasePurchaseOrder) GetAll(store_id string) ([]*entity.EntityPurchaseOrder, error) {
	return s.repo.GetAll(store_id)
}

func (s *UseCasePurchaseOrder) GetByID(store_id string, id string) (*entity.EntityPurchaseOrder, error) {
	return s.repo.GetByID(store_id, id)
}

func (s *UseCasePurchaseOrder) FilterbyNumber(store_id string, number string) (purchases []*entity.EntityPurchaseOrder, err error) {
	return s.repo.FilterbyNumber(store_id, number)
}

func (s *UseCasePurchaseOrder) FilterByDate(store_id string, initialDate string, endDate string) (purchases []*entity.EntityPurchaseOrder, err error) {
	return s.repo.FilterByDate(store_id, initialDate, endDate)
}

func (s *UseCasePurchaseOrder) FilterByProvider(store_id string, provider string) (purchases []*entity.EntityPurchaseOrder, err error) {
	return s.repo.FilterByProvider(store_id, provider)
}

func (s *UseCasePurchaseOrder) FilterByEmployee(store_id string, employee string) (purchases []*entity.EntityPurchaseOrder, err error) {
	return s.repo.FilterByEmployee(store_id, employee)
}

func (s *UseCasePurchaseOrder) Create(purchaseOrder entity.EntityPurchaseOrder) (*entity.EntityPurchaseOrder, error) {
	err := purchaseOrder.Validate()
	if err != nil {
		return nil, err
	}
	return s.repo.Create(purchaseOrder)
}

func (s *UseCasePurchaseOrder) Update(purchaseOrder entity.EntityPurchaseOrder) (*entity.EntityPurchaseOrder, error) {
	err := purchaseOrder.Validate()
	if err != nil {
		return nil, err
	}
	return s.repo.Update(purchaseOrder)
}

func (s *UseCasePurchaseOrder) Delete(store_id string, id string) error {
	return s.repo.Delete(store_id, id)
}
