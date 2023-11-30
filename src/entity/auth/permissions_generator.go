package entity

func GeneratePermissions() []EntityPermission {
	var permissions []EntityPermission
	permissions = append(permissions, generateUserPermissions()...)
	permissions = append(permissions, generateGroupPermissions()...)
	permissions = append(permissions, generatePermissionPermissions()...)
	permissions = append(permissions, generateStoreChainPermissions()...)
	permissions = append(permissions, generateStorePermissions()...)
	return permissions
}

func generateUserPermissions() []EntityPermission {
	var permissions []EntityPermission
	userCRUD := []string{"create", "read", "update", "delete"}
	for i := 0; i < len(userCRUD); i++ {
		permission, _ := CreatePermission(EntityPermission{
			Table: "entity_user",
			Name:  userCRUD[i],
		})
		permissions = append(permissions, *permission)
	}
	return permissions
}

func generateGroupPermissions() []EntityPermission {
	var permissions []EntityPermission
	groupCRUD := []string{"create", "read", "update", "delete"}
	for i := 0; i < len(groupCRUD); i++ {
		permission, _ := CreatePermission(EntityPermission{
			Table: "entity_group",
			Name:  groupCRUD[i],
		})
		permissions = append(permissions, *permission)
	}
	return permissions
}

func generatePermissionPermissions() []EntityPermission {
	var permissions []EntityPermission
	permissionCRUD := []string{"create", "read", "update", "delete"}
	for i := 0; i < len(permissionCRUD); i++ {
		permission, _ := CreatePermission(EntityPermission{
			Table: "entity_permission",
			Name:  permissionCRUD[i],
		})
		permissions = append(permissions, *permission)
	}
	return permissions
}

func generateStoreChainPermissions() []EntityPermission {
	var permissions []EntityPermission
	storeChainCRUD := []string{"create", "read", "update", "delete"}
	for i := 0; i < len(storeChainCRUD); i++ {
		permission, _ := CreatePermission(EntityPermission{
			Table: "entity_store_chain",
			Name:  storeChainCRUD[i],
		})
		permissions = append(permissions, *permission)
	}
	return permissions
}

func generateStorePermissions() []EntityPermission {
	var permissions []EntityPermission
	storeCRUD := []string{"create", "read", "update", "delete"}
	for i := 0; i < len(storeCRUD); i++ {
		permission, _ := CreatePermission(EntityPermission{
			Table: "entity_store",
			Name:  storeCRUD[i],
		})
		permissions = append(permissions, *permission)
	}
	return permissions
}
