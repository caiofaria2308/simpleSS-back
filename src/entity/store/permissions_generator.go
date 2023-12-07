package entity

import (
	entity_auth "main/entity/auth"
)

func GeneratePermissions() []entity_auth.EntityPermission {
	var permissions []entity_auth.EntityPermission
	permissions = append(permissions, generateStorePermissions()...)
	permissions = append(permissions, generateStoreChainPermissions()...)
	permissions = append(permissions, generateStoreEmployeePermissions()...)

	return permissions
}

func generateStorePermissions() []entity_auth.EntityPermission {
	return entity_auth.GenerateGenericPermissions("entity_store")
}

func generateStoreChainPermissions() []entity_auth.EntityPermission {
	return entity_auth.GenerateGenericPermissions("entity_store_chain")
}

func generateStoreEmployeePermissions() []entity_auth.EntityPermission {
	return entity_auth.GenerateGenericPermissions("entity_store_employee")
}
