package domain

type UserGroupMember struct {
	UserGroup UserGroup
	User      User
	Role      UserGroupRole
}
