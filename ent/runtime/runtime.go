// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"entdemo/ent/car"
	"entdemo/ent/group"
	"entdemo/ent/schema"
	"entdemo/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	carMixin := schema.Car{}.Mixin()
	carMixinHooks0 := carMixin[0].Hooks()
	car.Hooks[0] = carMixinHooks0[0]
	carMixinInters0 := carMixin[0].Interceptors()
	car.Interceptors[0] = carMixinInters0[0]
	groupMixin := schema.Group{}.Mixin()
	groupMixinHooks0 := groupMixin[0].Hooks()
	group.Hooks[0] = groupMixinHooks0[0]
	groupMixinInters0 := groupMixin[0].Interceptors()
	group.Interceptors[0] = groupMixinInters0[0]
	userMixin := schema.User{}.Mixin()
	userMixinHooks0 := userMixin[0].Hooks()
	user.Hooks[0] = userMixinHooks0[0]
	userMixinInters0 := userMixin[0].Interceptors()
	user.Interceptors[0] = userMixinInters0[0]
}

const (
	Version = "v0.12.2"                                         // Version of ent codegen.
	Sum     = "h1:Ndl/JvCX76xCtUDlrUfMnOKBRodAtxE5yfGYxjbOxmM=" // Sum of ent codegen.
)