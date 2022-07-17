package cache

type Cache interface {
	Get(k string) (any, bool)
	Set(k string, v any)
	Remove(k string)
}
