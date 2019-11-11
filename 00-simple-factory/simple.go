/*
@Time : 2019/11/11 10:14
@Author : Tux
@File : simple
@Description :
*/

package simple

import (
	"fmt"
)

// API is interface
type API interface {
	Say(name string) string
}

type hiAPI struct {
}

func (h *hiAPI) Say(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

type helloAPI struct {

}

func (h helloAPI) Say(name string) string {
	return fmt.Sprintf("Hello, %s", name)

}

func NewAPI(t int) API {
	if t == 1 {
		return &hiAPI{}
	} else if t == 2 {
		return &helloAPI{}
	}

	return nil
}
