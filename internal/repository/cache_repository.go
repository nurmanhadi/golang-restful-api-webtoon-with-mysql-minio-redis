package repository

type CacheRepository interface {
	SetView() error
	GetView() error
	DelView() error
}
