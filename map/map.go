package map

type Mapa[K comparable, V any] interface {
	Get(K) (V, bool)
	Set(K, V)
	Delete(K)
}

func New[K comparable, V any]() Storage[K, V] {
	return Storage[K, V](store[K, V]{ptrRef: make(map[K]V)})
}

type store[K comparable, V any] struct {
	ptrRef map[K]V
}

func (p store[K, V]) Set(key K, val V) {
	p.ptrRef[key] = val
}
func (p store[K, V]) Delete(key K) {
	delete(p.ptrRef, key)
}

func (p store[K, V]) Get(key K) (V, bool) {
	val, ok := p.ptrRef[key]
	return val, ok
}
