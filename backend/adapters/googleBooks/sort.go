package googleBooks

import "github.com/timsofteng/xyz-home-task/service"

type byTitle []service.Book

func (b byTitle) Len() int {
	return len(b)
}

func (b byTitle) Less(i, j int) bool {
	return b[i].Title < b[j].Title
}

func (b byTitle) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
