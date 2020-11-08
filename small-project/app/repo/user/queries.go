package user

const (
	userByLoginWithRole = `
select u.id, u.login, u.password, ur.role_id
from (
	select * 
	from users
	where login=$1) as u
left join users_roles as ur
on u.id=ur.user_id;`

	userByIDWithRole = `
select u.id, u.login, u.password, ur.role_id
from (
	select * 
	from users
	where id=$1) as u
left join users_roles as ur
on u.id=ur.user_id;`
)
