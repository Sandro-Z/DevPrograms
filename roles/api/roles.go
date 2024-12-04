package api

import (
	"context"
	"roles/db"
)

type NameRole struct {
	ID   string
	Name string
	Role string
}

type RolePermissions struct {
	ID         string
	Role       string
	Permission string
}

func GetAllNames(ctx context.Context) (response []string, err error) {
	var nameRoles []NameRole
	db.DB.Model(&NameRole{}).Find(&nameRoles)
	for _, nameRole := range nameRoles {
		response = append(response, nameRole.Name)
	}
	return response, err
}

func GetRoleByName(ctx context.Context, name string) (response string, err error) {
	var nameRole NameRole
	db.DB.Model(&NameRole{}).Where("Name = ?", name).First(&nameRole)
	response = nameRole.Role
	return response, err
}

func GetPermissionsByRole(ctx context.Context, role string) (response []string, err error) {
	var rolePermissions []RolePermissions
	db.DB.Model(&RolePermissions{}).Where("Role = ?", role).Find(&rolePermissions)
	for _, rolePermission := range rolePermissions {
		response = append(response, rolePermission.Permission)
	}
	return response, err
}

func CreateName(ctx context.Context, request *NameRole) (response *NameRole, err error) {
	err = db.DB.Model(&NameRole{}).Create(request).Error
	response = request
	return response, err
}

func CreatePermission(ctx context.Context, request *RolePermissions) (response *RolePermissions, err error) {
	err = db.DB.Model(&RolePermissions{}).Create(request).Error
	response = request
	return response, err
}

func DeleteName(ctx context.Context, name string) (response string, err error) {
	var nameRole NameRole
	err = db.DB.Model(&NameRole{}).Where("Name = ?", name).Delete(nameRole).Error
	response = nameRole.Role
	return response, err
}

func DeletePermission(cts context.Context, role string, permission string) (response string, err error) {
	var rolePermission RolePermissions
	permissionToDelete := &RolePermissions{
		Role:       role,
		Permission: permission,
	}
	err = db.DB.Model(&NameRole{}).Where(permissionToDelete).Delete(rolePermission).Error
	response = rolePermission.Role
	return response, err
}

func UpdateRole(ctx context.Context, name string, role string) (response string, err error) {
	var nameRole NameRole
	db.DB.Model(&NameRole{}).Where("Name = ?", name).First(&nameRole)
	nameRole.Role = role
	err = db.DB.Model(&NameRole{}).Save(nameRole).Error
	response = role
	return response, err
}
