package entity

import (
	entity_auth "main/entity/auth"
)

func GeneratePermissions() []entity_auth.EntityPermission {
	var permissions []entity_auth.EntityPermission
	permissions = append(permissions, generateSalePermissions()...)
	permissions = append(permissions, generateSaleProductPermissions()...)
	return permissions
}

func generateSalePermissions() []entity_auth.EntityPermission {
	return entity_auth.GenerateGenericPermissions("entity_sale")
}

func generateSaleProductPermissions() []entity_auth.EntityPermission {
	return entity_auth.GenerateGenericPermissions("entity_sale_product")
}
