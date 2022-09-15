package redis

import "strconv"

func GetDepDetailKey(id uint32) string {
	return CONSOLE_DEP + "_detail_" + strconv.Itoa(int(id))
}
