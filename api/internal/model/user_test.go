package model

import (
	"encoding/json"
	"testing"
)

func TestName(t *testing.T) {
	user := NewUser()
	b, err := json.Marshal(user)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(b))
}
