package repository

type CacheRepository interface {
	SetView() error
	GetView() (int, error)
	DelView() error
}
