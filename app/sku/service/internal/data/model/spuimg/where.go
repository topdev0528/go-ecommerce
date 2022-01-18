// Code generated by entc, DO NOT EDIT.

package spuimg

import (
	"mall-go/app/sku/service/internal/data/model/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.SpuImg {
	return predicate.SpuImg(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.SpuImg {
	return predicate.SpuImg(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.SpuImg {
	return predicate.SpuImg(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.SpuImg {
	return predicate.SpuImg(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.SpuImg {
	return predicate.SpuImg(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.SpuImg {
	return predicate.SpuImg(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.SpuImg {
	return predicate.SpuImg(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.SpuImg {
	return predicate.SpuImg(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.SpuImg {
	return predicate.SpuImg(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.SpuImg {
	return predicate.SpuImg(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.SpuImg {
	return predicate.SpuImg(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// DeleteTime applies equality check predicate on the "delete_time" field. It's identical to DeleteTimeEQ.
func DeleteTime(v time.Time) predicate.SpuImg {
	return predicate.SpuImg(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeleteTime), v))
	})
}

// Img applies equality check predicate on the "img" field. It's identical to ImgEQ.
func Img(v string) predicate.SpuImg {
	return predicate.SpuImg(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldImg), v))
	})
}

// SpuID applies equality check predicate on the "spu_id" field. It's identical to SpuIDEQ.
func SpuID(v int64) predicate.SpuImg {
	return predicate.SpuImg(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSpuID), v))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.SpuImg {
	return predicate.SpuImg(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.SpuImg {
	return predicate.SpuImg(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.SpuImg {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SpuImg(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.SpuImg {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SpuImg(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.SpuImg {
	return predicate.SpuImg(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.SpuImg {
	return predicate.SpuImg(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.SpuImg {
	return predicate.SpuImg(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.SpuImg {
	return predicate.SpuImg(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	})
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.SpuImg {
	return predicate.SpuImg(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.SpuImg {
	return predicate.SpuImg(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.SpuImg {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SpuImg(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.SpuImg {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SpuImg(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.SpuImg {
	return predicate.SpuImg(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.SpuImg {
	return predicate.SpuImg(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.SpuImg {
	return predicate.SpuImg(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.SpuImg {
	return predicate.SpuImg(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateTime), v))
	})
}

// DeleteTimeEQ applies the EQ predicate on the "delete_time" field.
func DeleteTimeEQ(v time.Time) predicate.SpuImg {
	return predicate.SpuImg(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeleteTime), v))
	})
}

// DeleteTimeNEQ applies the NEQ predicate on the "delete_time" field.
func DeleteTimeNEQ(v time.Time) predicate.SpuImg {
	return predicate.SpuImg(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeleteTime), v))
	})
}

// DeleteTimeIn applies the In predicate on the "delete_time" field.
func DeleteTimeIn(vs ...time.Time) predicate.SpuImg {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SpuImg(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDeleteTime), v...))
	})
}

// DeleteTimeNotIn applies the NotIn predicate on the "delete_time" field.
func DeleteTimeNotIn(vs ...time.Time) predicate.SpuImg {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SpuImg(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDeleteTime), v...))
	})
}

// DeleteTimeGT applies the GT predicate on the "delete_time" field.
func DeleteTimeGT(v time.Time) predicate.SpuImg {
	return predicate.SpuImg(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeleteTime), v))
	})
}

// DeleteTimeGTE applies the GTE predicate on the "delete_time" field.
func DeleteTimeGTE(v time.Time) predicate.SpuImg {
	return predicate.SpuImg(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeleteTime), v))
	})
}

// DeleteTimeLT applies the LT predicate on the "delete_time" field.
func DeleteTimeLT(v time.Time) predicate.SpuImg {
	return predicate.SpuImg(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeleteTime), v))
	})
}

// DeleteTimeLTE applies the LTE predicate on the "delete_time" field.
func DeleteTimeLTE(v time.Time) predicate.SpuImg {
	return predicate.SpuImg(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeleteTime), v))
	})
}

// DeleteTimeIsNil applies the IsNil predicate on the "delete_time" field.
func DeleteTimeIsNil() predicate.SpuImg {
	return predicate.SpuImg(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeleteTime)))
	})
}

// DeleteTimeNotNil applies the NotNil predicate on the "delete_time" field.
func DeleteTimeNotNil() predicate.SpuImg {
	return predicate.SpuImg(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeleteTime)))
	})
}

// ImgEQ applies the EQ predicate on the "img" field.
func ImgEQ(v string) predicate.SpuImg {
	return predicate.SpuImg(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldImg), v))
	})
}

// ImgNEQ applies the NEQ predicate on the "img" field.
func ImgNEQ(v string) predicate.SpuImg {
	return predicate.SpuImg(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldImg), v))
	})
}

// ImgIn applies the In predicate on the "img" field.
func ImgIn(vs ...string) predicate.SpuImg {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SpuImg(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldImg), v...))
	})
}

// ImgNotIn applies the NotIn predicate on the "img" field.
func ImgNotIn(vs ...string) predicate.SpuImg {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SpuImg(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldImg), v...))
	})
}

// ImgGT applies the GT predicate on the "img" field.
func ImgGT(v string) predicate.SpuImg {
	return predicate.SpuImg(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldImg), v))
	})
}

// ImgGTE applies the GTE predicate on the "img" field.
func ImgGTE(v string) predicate.SpuImg {
	return predicate.SpuImg(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldImg), v))
	})
}

// ImgLT applies the LT predicate on the "img" field.
func ImgLT(v string) predicate.SpuImg {
	return predicate.SpuImg(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldImg), v))
	})
}

// ImgLTE applies the LTE predicate on the "img" field.
func ImgLTE(v string) predicate.SpuImg {
	return predicate.SpuImg(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldImg), v))
	})
}

// ImgContains applies the Contains predicate on the "img" field.
func ImgContains(v string) predicate.SpuImg {
	return predicate.SpuImg(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldImg), v))
	})
}

// ImgHasPrefix applies the HasPrefix predicate on the "img" field.
func ImgHasPrefix(v string) predicate.SpuImg {
	return predicate.SpuImg(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldImg), v))
	})
}

// ImgHasSuffix applies the HasSuffix predicate on the "img" field.
func ImgHasSuffix(v string) predicate.SpuImg {
	return predicate.SpuImg(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldImg), v))
	})
}

// ImgEqualFold applies the EqualFold predicate on the "img" field.
func ImgEqualFold(v string) predicate.SpuImg {
	return predicate.SpuImg(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldImg), v))
	})
}

// ImgContainsFold applies the ContainsFold predicate on the "img" field.
func ImgContainsFold(v string) predicate.SpuImg {
	return predicate.SpuImg(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldImg), v))
	})
}

// SpuIDEQ applies the EQ predicate on the "spu_id" field.
func SpuIDEQ(v int64) predicate.SpuImg {
	return predicate.SpuImg(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSpuID), v))
	})
}

// SpuIDNEQ applies the NEQ predicate on the "spu_id" field.
func SpuIDNEQ(v int64) predicate.SpuImg {
	return predicate.SpuImg(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSpuID), v))
	})
}

// SpuIDIn applies the In predicate on the "spu_id" field.
func SpuIDIn(vs ...int64) predicate.SpuImg {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SpuImg(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSpuID), v...))
	})
}

// SpuIDNotIn applies the NotIn predicate on the "spu_id" field.
func SpuIDNotIn(vs ...int64) predicate.SpuImg {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SpuImg(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSpuID), v...))
	})
}

// SpuIDIsNil applies the IsNil predicate on the "spu_id" field.
func SpuIDIsNil() predicate.SpuImg {
	return predicate.SpuImg(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSpuID)))
	})
}

// SpuIDNotNil applies the NotNil predicate on the "spu_id" field.
func SpuIDNotNil() predicate.SpuImg {
	return predicate.SpuImg(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSpuID)))
	})
}

// HasSpu applies the HasEdge predicate on the "spu" edge.
func HasSpu() predicate.SpuImg {
	return predicate.SpuImg(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SpuTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SpuTable, SpuColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSpuWith applies the HasEdge predicate on the "spu" edge with a given conditions (other predicates).
func HasSpuWith(preds ...predicate.Spu) predicate.SpuImg {
	return predicate.SpuImg(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SpuInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SpuTable, SpuColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SpuImg) predicate.SpuImg {
	return predicate.SpuImg(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SpuImg) predicate.SpuImg {
	return predicate.SpuImg(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.SpuImg) predicate.SpuImg {
	return predicate.SpuImg(func(s *sql.Selector) {
		p(s.Not())
	})
}