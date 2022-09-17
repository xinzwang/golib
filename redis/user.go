package redis

import "strconv"

func GetUserDetailKey(id uint32) string {
	return CONSOLE_USER + "_detail_" + strconv.Itoa(int(id))
}
