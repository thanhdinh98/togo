package service

import "gtodo/src"

type Context struct {
}

func (c *Context) CheckPermission() (bool, error) {
	return true, nil
}

func NewContext() src.IContextService {
	return &Context{}
}
