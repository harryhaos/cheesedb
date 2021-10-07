package cheesedb


type db struct{

}

type Db interface {
}

func NewDb() (Db, error) {
	return &db{}, nil
}

func (d *db) Get() {

}

func (d *db) Set() {

}
