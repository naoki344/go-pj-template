package pagemodel

type Size int64

type Total int64

type Current int64

type PageResult struct {
	Size    Size
	Total   Total
	Current Current
}
