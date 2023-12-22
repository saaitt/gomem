package domain

type UseCase interface {
	Get(bucket, key string) (interface{}, error)
}
