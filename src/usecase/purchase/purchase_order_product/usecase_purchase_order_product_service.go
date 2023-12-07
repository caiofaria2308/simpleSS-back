package usecase_purchase_order_product

import entity "main/entity/purchase"

type UseCasePurchaseOrderProduct struct {
	repo IRepositoryPurchaseOrderProduct
}

func NewService(repo IRepositoryPurchaseOrderProduct) *UseCasePurchaseOrderProduct {
	return &UseCasePurchaseOrderProduct{repo: repo}
}

func (s *UseCasePurchaseOrderProduct) GetAll(store_id string) ([]*entity.EntityPurchaseOrderProduct, error) {
	return s.repo.GetAll(store_id)
}

func (s *UseCasePurchaseOrderProduct) GetByID(store_id string, purchaseOrderID string, id string) (*entity.EntityPurchaseOrderProduct, error) {
	return s.repo.GetByID(store_id, purchaseOrderID, id)
}

func (s *UseCasePurchaseOrderProduct) Create(purchaseOrderProduct entity.EntityPurchaseOrderProduct) (*entity.EntityPurchaseOrderProduct, error) {
	err := purchaseOrderProduct.Validate()
	err = entity.CreatePurchaseOrderProduct(&purchaseOrderProduct)
	if err != nil {
		return nil, err
	}
	return s.repo.Create(purchaseOrderProduct)
}

func (s *UseCasePurchaseOrderProduct) Update(purchaseOrderProduct entity.EntityPurchaseOrderProduct) (*entity.EntityPurchaseOrderProduct, error) {
	err := purchaseOrderProduct.Validate()
	err = entity.UpdatePurchaseOrderProduct(&purchaseOrderProduct)
	if err != nil {
		return nil, err
	}
	return s.repo.Update(purchaseOrderProduct)
}

func (s *UseCasePurchaseOrderProduct) Delete(store_id string, purchaseOrderID string, id string) error {
	return s.repo.Delete(store_id, purchaseOrderID, id)
}
