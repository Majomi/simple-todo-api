package todo

type Service interface {
	DB
}

type service struct {
	DB
}

func NewService(storage DB) Service {
	return service{DB: storage}
}
