package mo_test

import (
	"encoding/json"
	"fmt"
	"mo"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOptionsJSON(t *testing.T) {

	type User struct {
		Name mo.Option[string] `json:"name"`
		Age  mo.Option[int]    `json:"age"`
	}

	json1 := mo.ResultFrom(json.Marshal(User{
		Name: mo.Some("Alice"),
		Age:  mo.Some(30),
	})).Try()
	assert.Equal(t, string(json1), `{"name":"Alice","age":30}`)

	json2 := mo.ResultFrom(json.Marshal(User{
		Name: mo.None[string](),
		Age:  mo.Some(25),
	})).Try()
	assert.Equal(t, string(json2), `{"name":null,"age":25}`)

	// Deserialization example
	var user3 User
	json.Unmarshal([]byte(`{"name": "Bob", "age": null}`), &user3)
	fmt.Printf("%+v\n", user3)
	assert.Equal(t, user3, User{
		Name: mo.Some("Bob"),
		Age:  mo.None[int](),
	}, "user3")
}
