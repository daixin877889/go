package pageSize

func GetPageSize(page,size int) (int, int) {
	offset := (page - 1) * size
	return offset, size
}
