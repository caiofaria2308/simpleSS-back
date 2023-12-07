package entity

func GeneratePermissions() []EntityPermission {
	var permissions []EntityPermission
	permissions = append(permissions, generateUserPermissions()...)
	permissions = append(permissions, generateGroupPermissions()...)
	permissions = append(permissions, generatePermissionPermissions()...)
	return permissions
}

func GenerateGenericPermissions(table string) []EntityPermission {
	var permissions []EntityPermission
	userCRUD := []string{"create", "list", "read", "update", "delete"}
	for i := 0; i < len(userCRUD); i++ {
		permission := EntityPermission{
			Table: table,
			Name:  userCRUD[i],
		}
		CreatePermission(&permission)
		permissions = append(permissions, permission)
	}
	return permissions
}

func generateUserPermissions() []EntityPermission {
	return GenerateGenericPermissions("entity_user")
}

func generateGroupPermissions() []EntityPermission {
	return GenerateGenericPermissions("entity_group")
}

func generatePermissionPermissions() []EntityPermission {
	return GenerateGenericPermissions("entity_permission")
}
