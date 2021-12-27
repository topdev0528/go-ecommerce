package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Activity struct {
	ent.Schema
}

func (Activity) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "activity"},
	}
}
func (Activity) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Comment(""),
		field.String("description").Comment(""),
		field.Time("start_time").Comment(""),
		field.Time("end_time").Comment(""),

		field.String("remark").Comment(""),
		field.Int8("online").Comment(""),
		field.String("entrance_img").Comment(""),
		field.String("internal_top_img").Comment(""),
		field.String("name").Comment(""),
	}
}
func (Activity) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
func (Activity) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("spu", Spu.Type).StorageKey(
			edge.Table("activity_spu"), edge.Columns("activity_id", "spu_id")),
	}
}
