package requests

import "errors"

var (
	Roles = map[string]int{
		"Client": 1,
		"Driver": 2,
		"Admin":  3,
	}
)

func GetRoleId(role string) (int, error) {
	roleId, ok := Roles[role]
	if !ok {
		return 0, errors.New("role does not exist")
	}

	return roleId, nil
}

func IsAdmin(roleId int) bool {
	return Roles["Admin"] == roleId
}

func IsDriver(roleId int) bool {
	return Roles["Driver"] == roleId
}

func IsClient(roleId int) bool {
	return Roles["Client"] == roleId
}
