package role

const (
	roleWithPermissions = `
select r.id, r.name, p.permission_id, p.source, p.mode
from roles as r
inner join
	(
	select rp.role_id, rp.permission_id, p.source, p.mode
	from
		(
			select role_id, permission_id
			from roles_permissions
			where role_id=$1
		) as rp
	left join permissions as p
	on rp.permission_id=p.id
	) as p
on r.id=p.role_id`
)
