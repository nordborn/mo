package mo_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/nordborn/mo"
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

func TestOptionsSQL(t *testing.T) {
	var opt mo.Option[int]
	mo.TryErr(opt.Scan(3))
	assert.Equal(t, opt, mo.Some(3))
	mo.TryErr(opt.Scan(nil))
	assert.Equal(t, opt, mo.None[int]())
}

func TestOptionsSQLStr(t *testing.T) {
	var opt mo.Option[string]
	mo.TryErr(opt.Scan("test"))
	assert.Equal(t, opt, mo.Some("test"))
	mo.TryErr(opt.Scan(nil))
	assert.Equal(t, opt, mo.None[string]())
}
