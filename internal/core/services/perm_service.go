package services

import (
	"github.com/mxpadidar/letsgo/internal/core/types"
)

// RolePermsMap represents a mapping of roles to permissions.
// It is used to define the permissions for each role.
type RolePermsMap map[types.Role]types.Permission

// PermService handles access control based on roles and permissions.
// It provides methods to check if a user has a specific permission.
// Role Permissions can be loaded from a configuration file or database.
type PermService struct {
	rolePerms RolePermsMap
	logger    LogService
}

// NewPermService creates a new instance of PermService.
func NewPermService(rolePerms RolePermsMap, logger LogService) *PermService {
	return &PermService{
		rolePerms: rolePerms,
		logger:    logger,
	}
}

// CheckPerm checks if a role has a specific permission.
// It returns true if the role has the permission, false otherwise.
func (ps *PermService) CheckPerm(role types.Role, require types.Permission) bool {
	perm, ok := ps.rolePerms[role]
	if !ok {
		ps.logger.Warnf("Role %d does not exist in PermService", role)
		return false
	}
	return require&perm == require
}
