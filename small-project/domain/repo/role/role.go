package role

import "chaelub/thinksurance/small-project/domain/models"

type RoleRepoI interface {
	GetRoleById(id int64) (models.Role, error)
	// Expand the interface with methods like SaveRole(models.Role) and so on
}
