package redis

import "strconv"

func GetPosDetailKey(id uint32) string {
	return CONSOLE_POS + "_detail_" + strconv.Itoa(int(id))
}
