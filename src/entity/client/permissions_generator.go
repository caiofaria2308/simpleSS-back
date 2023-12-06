package entity

import (
	entity_auth "main/entity/auth"
)

func GeneratePermissions() []entity_auth.EntityPermission {
	var permissions []entity_auth.EntityPermission
	permissions = append(permissions, generateClientPermissions()...)
	return permissions
}

func generateClientPermissions() []entity_auth.EntityPermission {
	return entity_auth.GenerateGenericPermissions("entity_client")
}
