package main

type Store struct {
	kv map[string]string
}

func (s *Store) GetValueFromStore(key string) string {
	return s.kv[key]
}
