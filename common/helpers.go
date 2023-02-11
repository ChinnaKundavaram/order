package common

import "strconv"

//
func GetPageDetails(page, size string) (int, int) {
	pageNo, err := strconv.ParseInt(page, 10, 32)
	if err != nil || pageNo < 1 {
		pageNo = 0
	}
	pageSize, err := strconv.ParseInt(size, 10, 32)
	if err != nil || pageSize < 1 {
		pageSize = 10
	}
	return int(pageNo), int(pageSize)
}
