package entity

import (
	entity_auth "main/entity/auth"
)

func GeneratePermissions() []entity_auth.EntityPermission {
	var permissions []entity_auth.EntityPermission
	permissions = append(permissions, generateProductGroupPermissions()...)
	permissions = append(permissions, generateProductPromotionPermissions()...)
	permissions = append(permissions, generateProductPermissions()...)
	permissions = append(permissions, generateProductStockPermissions()...)
	return permissions
}

func generateProductGroupPermissions() []entity_auth.EntityPermission {
	return entity_auth.GenerateGenericPermissions("entity_product_group")
}

func generateProductPromotionPermissions() []entity_auth.EntityPermission {
	return entity_auth.GenerateGenericPermissions("entity_product_promotion")
}

func generateProductPermissions() []entity_auth.EntityPermission {
	return entity_auth.GenerateGenericPermissions("entity_product")
}

func generateProductStockPermissions() []entity_auth.EntityPermission {
	return entity_auth.GenerateGenericPermissions("entity_product_stock")
}
