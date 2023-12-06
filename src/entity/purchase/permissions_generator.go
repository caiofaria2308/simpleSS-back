package entity

import (
	entity_auth "main/entity/auth"
)

func GeneratePermissions() []entity_auth.EntityPermission {
	var permissions []entity_auth.EntityPermission
	permissions = append(permissions, generatePurchaseOrderPermissions()...)
	permissions = append(permissions, generatePurchaseOrderProductPermissions()...)
	return permissions
}

func generatePurchaseOrderPermissions() []entity_auth.EntityPermission {
	return entity_auth.GenerateGenericPermissions("entity_purchase_order")
}

func generatePurchaseOrderProductPermissions() []entity_auth.EntityPermission {
	return entity_auth.GenerateGenericPermissions("entity_purchase_order_product")
}
