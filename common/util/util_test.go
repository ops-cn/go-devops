package util

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStructMapToStruct(t *testing.T) {
	type Foo struct {
		SBar      *string
		FBar      float64
		IBar      int64
		TBar      time.Time
		PBar      *string
		CreatedAt time.Time
	}

	type NFoo struct {
		CreatedAt time.Time
	}

	type TFoo struct {
		NFoo
		ID   int64
		SBar string
		FBar *float64
		IBar *int64
		TBar *time.Time
		PBar string
	}

	bar := "bar"
	var foo = Foo{
		SBar:      &bar,
		FBar:      1.1,
		IBar:      1,
		TBar:      time.Now(),
		CreatedAt: time.Now(),
	}

	var tfoo TFoo
	StructMapToStruct(&foo, &tfoo)

	assert.Equal(t, *foo.SBar, tfoo.SBar)
	assert.Equal(t, foo.FBar, *tfoo.FBar)
	assert.Equal(t, foo.IBar, *tfoo.IBar)
	assert.Equal(t, foo.TBar, *tfoo.TBar)
	assert.Equal(t, foo.CreatedAt, tfoo.CreatedAt)
	assert.Equal(t, tfoo.PBar, "")
}

type Role struct {
	RoleId   string
	RoleName string
}

type Src struct {
	Name string
	ID   string
	Role
}

type Dst struct {
	Name string
	ID   string
	Role
}

func TestStructCopy(t *testing.T) {
	role := Role{
		RoleId:   "000",
		RoleName: "管理员",
	}
	src := &Src{"katy", "000", role}

	var dst Dst
	StructCopy(&dst, src)
	fmt.Println(dst)
}
