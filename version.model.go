package types

import (
	baseapp "github.com/natifdevelopment/go-baseapp"
)

func (Version) TableName() string {
	return "t_version"
}

type Version struct {
	baseapp.BaseModel
}
