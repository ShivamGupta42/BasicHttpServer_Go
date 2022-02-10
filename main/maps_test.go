package main

import (
	"testing"
)

func TestStore_GetValueFromStore(t *testing.T) {

	store := Store{kv: map[string]string{"val": "1"}}

	t.Run("testing value not found", func(t *testing.T) {
		got := store.GetValueFromStore("rand")
		if got != "" {
			t.Fail()
		}
	})

}
