// Code generated by ent, DO NOT EDIT.

package refund

import (
	"mall-go/app/app/service/internal/data/model/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// DeleteTime applies equality check predicate on the "delete_time" field. It's identical to DeleteTimeEQ.
func DeleteTime(v time.Time) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeleteTime), v))
	})
}

// RefundNo applies equality check predicate on the "refund_no" field. It's identical to RefundNoEQ.
func RefundNo(v string) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRefundNo), v))
	})
}

// TransactionID applies equality check predicate on the "transaction_id" field. It's identical to TransactionIDEQ.
func TransactionID(v string) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTransactionID), v))
	})
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int64) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// Reason applies equality check predicate on the "reason" field. It's identical to ReasonEQ.
func Reason(v string) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReason), v))
	})
}

// OrderID applies equality check predicate on the "order_id" field. It's identical to OrderIDEQ.
func OrderID(v int64) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderID), v))
	})
}

// OrderSubID applies equality check predicate on the "order_sub_id" field. It's identical to OrderSubIDEQ.
func OrderSubID(v int64) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderSubID), v))
	})
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v int) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Refund {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Refund(func(s *sql.Selector) {
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
func CreateTimeNotIn(vs ...time.Time) predicate.Refund {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Refund(func(s *sql.Selector) {
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
func CreateTimeGT(v time.Time) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	})
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Refund {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Refund(func(s *sql.Selector) {
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
func UpdateTimeNotIn(vs ...time.Time) predicate.Refund {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Refund(func(s *sql.Selector) {
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
func UpdateTimeGT(v time.Time) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateTime), v))
	})
}

// DeleteTimeEQ applies the EQ predicate on the "delete_time" field.
func DeleteTimeEQ(v time.Time) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeleteTime), v))
	})
}

// DeleteTimeNEQ applies the NEQ predicate on the "delete_time" field.
func DeleteTimeNEQ(v time.Time) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeleteTime), v))
	})
}

// DeleteTimeIn applies the In predicate on the "delete_time" field.
func DeleteTimeIn(vs ...time.Time) predicate.Refund {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Refund(func(s *sql.Selector) {
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
func DeleteTimeNotIn(vs ...time.Time) predicate.Refund {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Refund(func(s *sql.Selector) {
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
func DeleteTimeGT(v time.Time) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeleteTime), v))
	})
}

// DeleteTimeGTE applies the GTE predicate on the "delete_time" field.
func DeleteTimeGTE(v time.Time) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeleteTime), v))
	})
}

// DeleteTimeLT applies the LT predicate on the "delete_time" field.
func DeleteTimeLT(v time.Time) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeleteTime), v))
	})
}

// DeleteTimeLTE applies the LTE predicate on the "delete_time" field.
func DeleteTimeLTE(v time.Time) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeleteTime), v))
	})
}

// DeleteTimeIsNil applies the IsNil predicate on the "delete_time" field.
func DeleteTimeIsNil() predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeleteTime)))
	})
}

// DeleteTimeNotNil applies the NotNil predicate on the "delete_time" field.
func DeleteTimeNotNil() predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeleteTime)))
	})
}

// RefundNoEQ applies the EQ predicate on the "refund_no" field.
func RefundNoEQ(v string) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRefundNo), v))
	})
}

// RefundNoNEQ applies the NEQ predicate on the "refund_no" field.
func RefundNoNEQ(v string) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRefundNo), v))
	})
}

// RefundNoIn applies the In predicate on the "refund_no" field.
func RefundNoIn(vs ...string) predicate.Refund {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Refund(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRefundNo), v...))
	})
}

// RefundNoNotIn applies the NotIn predicate on the "refund_no" field.
func RefundNoNotIn(vs ...string) predicate.Refund {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Refund(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRefundNo), v...))
	})
}

// RefundNoGT applies the GT predicate on the "refund_no" field.
func RefundNoGT(v string) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRefundNo), v))
	})
}

// RefundNoGTE applies the GTE predicate on the "refund_no" field.
func RefundNoGTE(v string) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRefundNo), v))
	})
}

// RefundNoLT applies the LT predicate on the "refund_no" field.
func RefundNoLT(v string) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRefundNo), v))
	})
}

// RefundNoLTE applies the LTE predicate on the "refund_no" field.
func RefundNoLTE(v string) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRefundNo), v))
	})
}

// RefundNoContains applies the Contains predicate on the "refund_no" field.
func RefundNoContains(v string) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRefundNo), v))
	})
}

// RefundNoHasPrefix applies the HasPrefix predicate on the "refund_no" field.
func RefundNoHasPrefix(v string) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRefundNo), v))
	})
}

// RefundNoHasSuffix applies the HasSuffix predicate on the "refund_no" field.
func RefundNoHasSuffix(v string) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRefundNo), v))
	})
}

// RefundNoEqualFold applies the EqualFold predicate on the "refund_no" field.
func RefundNoEqualFold(v string) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRefundNo), v))
	})
}

// RefundNoContainsFold applies the ContainsFold predicate on the "refund_no" field.
func RefundNoContainsFold(v string) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRefundNo), v))
	})
}

// TransactionIDEQ applies the EQ predicate on the "transaction_id" field.
func TransactionIDEQ(v string) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTransactionID), v))
	})
}

// TransactionIDNEQ applies the NEQ predicate on the "transaction_id" field.
func TransactionIDNEQ(v string) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTransactionID), v))
	})
}

// TransactionIDIn applies the In predicate on the "transaction_id" field.
func TransactionIDIn(vs ...string) predicate.Refund {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Refund(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTransactionID), v...))
	})
}

// TransactionIDNotIn applies the NotIn predicate on the "transaction_id" field.
func TransactionIDNotIn(vs ...string) predicate.Refund {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Refund(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTransactionID), v...))
	})
}

// TransactionIDGT applies the GT predicate on the "transaction_id" field.
func TransactionIDGT(v string) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTransactionID), v))
	})
}

// TransactionIDGTE applies the GTE predicate on the "transaction_id" field.
func TransactionIDGTE(v string) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTransactionID), v))
	})
}

// TransactionIDLT applies the LT predicate on the "transaction_id" field.
func TransactionIDLT(v string) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTransactionID), v))
	})
}

// TransactionIDLTE applies the LTE predicate on the "transaction_id" field.
func TransactionIDLTE(v string) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTransactionID), v))
	})
}

// TransactionIDContains applies the Contains predicate on the "transaction_id" field.
func TransactionIDContains(v string) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTransactionID), v))
	})
}

// TransactionIDHasPrefix applies the HasPrefix predicate on the "transaction_id" field.
func TransactionIDHasPrefix(v string) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTransactionID), v))
	})
}

// TransactionIDHasSuffix applies the HasSuffix predicate on the "transaction_id" field.
func TransactionIDHasSuffix(v string) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTransactionID), v))
	})
}

// TransactionIDEqualFold applies the EqualFold predicate on the "transaction_id" field.
func TransactionIDEqualFold(v string) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTransactionID), v))
	})
}

// TransactionIDContainsFold applies the ContainsFold predicate on the "transaction_id" field.
func TransactionIDContainsFold(v string) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTransactionID), v))
	})
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int64) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int64) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	})
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int64) predicate.Refund {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Refund(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUserID), v...))
	})
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int64) predicate.Refund {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Refund(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	})
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v int64) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserID), v))
	})
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v int64) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserID), v))
	})
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v int64) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserID), v))
	})
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v int64) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserID), v))
	})
}

// UserIDIsNil applies the IsNil predicate on the "user_id" field.
func UserIDIsNil() predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUserID)))
	})
}

// UserIDNotNil applies the NotNil predicate on the "user_id" field.
func UserIDNotNil() predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUserID)))
	})
}

// ReasonEQ applies the EQ predicate on the "reason" field.
func ReasonEQ(v string) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReason), v))
	})
}

// ReasonNEQ applies the NEQ predicate on the "reason" field.
func ReasonNEQ(v string) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldReason), v))
	})
}

// ReasonIn applies the In predicate on the "reason" field.
func ReasonIn(vs ...string) predicate.Refund {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Refund(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldReason), v...))
	})
}

// ReasonNotIn applies the NotIn predicate on the "reason" field.
func ReasonNotIn(vs ...string) predicate.Refund {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Refund(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldReason), v...))
	})
}

// ReasonGT applies the GT predicate on the "reason" field.
func ReasonGT(v string) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldReason), v))
	})
}

// ReasonGTE applies the GTE predicate on the "reason" field.
func ReasonGTE(v string) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldReason), v))
	})
}

// ReasonLT applies the LT predicate on the "reason" field.
func ReasonLT(v string) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldReason), v))
	})
}

// ReasonLTE applies the LTE predicate on the "reason" field.
func ReasonLTE(v string) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldReason), v))
	})
}

// ReasonContains applies the Contains predicate on the "reason" field.
func ReasonContains(v string) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldReason), v))
	})
}

// ReasonHasPrefix applies the HasPrefix predicate on the "reason" field.
func ReasonHasPrefix(v string) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldReason), v))
	})
}

// ReasonHasSuffix applies the HasSuffix predicate on the "reason" field.
func ReasonHasSuffix(v string) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldReason), v))
	})
}

// ReasonEqualFold applies the EqualFold predicate on the "reason" field.
func ReasonEqualFold(v string) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldReason), v))
	})
}

// ReasonContainsFold applies the ContainsFold predicate on the "reason" field.
func ReasonContainsFold(v string) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldReason), v))
	})
}

// OrderIDEQ applies the EQ predicate on the "order_id" field.
func OrderIDEQ(v int64) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderID), v))
	})
}

// OrderIDNEQ applies the NEQ predicate on the "order_id" field.
func OrderIDNEQ(v int64) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOrderID), v))
	})
}

// OrderIDIn applies the In predicate on the "order_id" field.
func OrderIDIn(vs ...int64) predicate.Refund {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Refund(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOrderID), v...))
	})
}

// OrderIDNotIn applies the NotIn predicate on the "order_id" field.
func OrderIDNotIn(vs ...int64) predicate.Refund {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Refund(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOrderID), v...))
	})
}

// OrderIDGT applies the GT predicate on the "order_id" field.
func OrderIDGT(v int64) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOrderID), v))
	})
}

// OrderIDGTE applies the GTE predicate on the "order_id" field.
func OrderIDGTE(v int64) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOrderID), v))
	})
}

// OrderIDLT applies the LT predicate on the "order_id" field.
func OrderIDLT(v int64) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOrderID), v))
	})
}

// OrderIDLTE applies the LTE predicate on the "order_id" field.
func OrderIDLTE(v int64) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOrderID), v))
	})
}

// OrderIDIsNil applies the IsNil predicate on the "order_id" field.
func OrderIDIsNil() predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldOrderID)))
	})
}

// OrderIDNotNil applies the NotNil predicate on the "order_id" field.
func OrderIDNotNil() predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldOrderID)))
	})
}

// OrderSubIDEQ applies the EQ predicate on the "order_sub_id" field.
func OrderSubIDEQ(v int64) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderSubID), v))
	})
}

// OrderSubIDNEQ applies the NEQ predicate on the "order_sub_id" field.
func OrderSubIDNEQ(v int64) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOrderSubID), v))
	})
}

// OrderSubIDIn applies the In predicate on the "order_sub_id" field.
func OrderSubIDIn(vs ...int64) predicate.Refund {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Refund(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOrderSubID), v...))
	})
}

// OrderSubIDNotIn applies the NotIn predicate on the "order_sub_id" field.
func OrderSubIDNotIn(vs ...int64) predicate.Refund {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Refund(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOrderSubID), v...))
	})
}

// OrderSubIDGT applies the GT predicate on the "order_sub_id" field.
func OrderSubIDGT(v int64) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOrderSubID), v))
	})
}

// OrderSubIDGTE applies the GTE predicate on the "order_sub_id" field.
func OrderSubIDGTE(v int64) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOrderSubID), v))
	})
}

// OrderSubIDLT applies the LT predicate on the "order_sub_id" field.
func OrderSubIDLT(v int64) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOrderSubID), v))
	})
}

// OrderSubIDLTE applies the LTE predicate on the "order_sub_id" field.
func OrderSubIDLTE(v int64) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOrderSubID), v))
	})
}

// OrderSubIDIsNil applies the IsNil predicate on the "order_sub_id" field.
func OrderSubIDIsNil() predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldOrderSubID)))
	})
}

// OrderSubIDNotNil applies the NotNil predicate on the "order_sub_id" field.
func OrderSubIDNotNil() predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldOrderSubID)))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v int) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v int) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...int) predicate.Refund {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Refund(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStatus), v...))
	})
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...int) predicate.Refund {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Refund(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v int) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStatus), v))
	})
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v int) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStatus), v))
	})
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v int) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStatus), v))
	})
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v int) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStatus), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Refund) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Refund) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
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
func Not(p predicate.Refund) predicate.Refund {
	return predicate.Refund(func(s *sql.Selector) {
		p(s.Not())
	})
}
