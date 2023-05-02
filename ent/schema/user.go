package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		SoftDeleteMixin{},
	}
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		// 定义一条名为"cars"的，指向Car类型的边，在car表中使用"user_id"保存关联关系
		edge.To("cars", Car.Type).StorageKey(edge.Column("user_id")),
		// 创建一个指向Group类型的反向关联关系”groups“。
		// 通过Ref方法，显示的将其与在Group中定义的“users”关联关系关联。
		edge.From("groups", Group.Type).
			Ref("users"),
	}
}
