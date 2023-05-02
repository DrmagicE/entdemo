package schema

import (
	"context"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"entdemo/ent/hook"
)

// Car holds the schema definition for the Car entity.
type Car struct {
	ent.Schema
}

func (c Car) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
					return next.Mutate(ctx, m)
				})
			},
			ent.OpUpdate|ent.OpUpdateOne,
		),
	}
}

func (Car) Mixin() []ent.Mixin {
	return []ent.Mixin{
		SoftDeleteMixin{},
	}
}

// Fields of the Car.
func (Car) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the Car.
func (Car) Edges() []ent.Edge {
	return []ent.Edge{
		// 创建一个指向User类型的反向关联关系”owner“。
		// 通过Ref方法，显示的将其与在User中定义的“cars”关联关系关联。
		edge.From("owner", User.Type).
			Ref("cars").
			// 通过Unique表达Car只能属于一个User。
			// (如果不加Unique，那表达的就是多对多的关系了)
			Unique(),
	}
}
