package redis

import "strconv"

func GetRepairTypeDetailKey(id uint32) string {
	return WEIXIN_REPAIR_TYPE + "_detail_" + strconv.Itoa(int(id))
}
