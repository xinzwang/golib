package redis

import "strconv"

func GetCompanyDetailKey(id uint32) string {
	return CONSOLE_COMPANY + "_detail_" + strconv.Itoa(int(id))
}
