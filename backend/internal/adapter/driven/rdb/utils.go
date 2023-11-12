package rdbadapter

func calcPaging(pageNumber int64, pageSize int64) (int64, int64) {
	start := (pageSize * (pageNumber - 1)) + 1
	end := pageSize * pageNumber
	return start, end
}
