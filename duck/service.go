package duck

import "github.com/Chanokthorn/go-public-lib-test/quack"

type DuckService struct {}

func NewDuckService() *DuckService {
	return &DuckService{}
}

func (d *DuckService) Quack() error {
	return quack.Execute("quack by duck")
}
