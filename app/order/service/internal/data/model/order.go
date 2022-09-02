// Code generated by entc, DO NOT EDIT.

package model

import (
	"fmt"
	"mall-go/app/order/service/internal/data/model/order"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Order is the model entity for the Order schema.
type Order struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// DeleteTime holds the value of the "delete_time" field.
	DeleteTime time.Time `json:"delete_time,omitempty"`
	// OrderNo holds the value of the "order_no" field.
	OrderNo string `json:"order_no,omitempty"`
	// TransactionID holds the value of the "transaction_id" field.
	// 支付平台流水号
	TransactionID string `json:"transaction_id,omitempty"`
	// UserID holds the value of the "user_id" field.
	// user表外键
	UserID int64 `json:"user_id,omitempty"`
	// TotalPrice holds the value of the "total_price" field.
	TotalPrice float64 `json:"total_price,omitempty"`
	// TotalCount holds the value of the "total_count" field.
	TotalCount int `json:"total_count,omitempty"`
	// FinalTotalPrice holds the value of the "final_total_price" field.
	FinalTotalPrice float64 `json:"final_total_price,omitempty"`
	// Status holds the value of the "status" field.
	Status int `json:"status,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrderQuery when eager-loading is set.
	Edges OrderEdges `json:"edges"`
}

// OrderEdges holds the relations/edges for other nodes in the graph.
type OrderEdges struct {
	// OrderSnap holds the value of the order_snap edge.
	OrderSnap []*OrderSnap `json:"order_snap,omitempty"`
	// OrderSub holds the value of the order_sub edge.
	OrderSub []*OrderSub `json:"order_sub,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// OrderSnapOrErr returns the OrderSnap value or an error if the edge
// was not loaded in eager-loading.
func (e OrderEdges) OrderSnapOrErr() ([]*OrderSnap, error) {
	if e.loadedTypes[0] {
		return e.OrderSnap, nil
	}
	return nil, &NotLoadedError{edge: "order_snap"}
}

// OrderSubOrErr returns the OrderSub value or an error if the edge
// was not loaded in eager-loading.
func (e OrderEdges) OrderSubOrErr() ([]*OrderSub, error) {
	if e.loadedTypes[1] {
		return e.OrderSub, nil
	}
	return nil, &NotLoadedError{edge: "order_sub"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Order) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case order.FieldTotalPrice, order.FieldFinalTotalPrice:
			values[i] = new(sql.NullFloat64)
		case order.FieldID, order.FieldUserID, order.FieldTotalCount, order.FieldStatus:
			values[i] = new(sql.NullInt64)
		case order.FieldOrderNo, order.FieldTransactionID:
			values[i] = new(sql.NullString)
		case order.FieldCreateTime, order.FieldUpdateTime, order.FieldDeleteTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Order", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Order fields.
func (o *Order) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case order.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			o.ID = int64(value.Int64)
		case order.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				o.CreateTime = value.Time
			}
		case order.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				o.UpdateTime = value.Time
			}
		case order.FieldDeleteTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field delete_time", values[i])
			} else if value.Valid {
				o.DeleteTime = value.Time
			}
		case order.FieldOrderNo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field order_no", values[i])
			} else if value.Valid {
				o.OrderNo = value.String
			}
		case order.FieldTransactionID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field transaction_id", values[i])
			} else if value.Valid {
				o.TransactionID = value.String
			}
		case order.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				o.UserID = value.Int64
			}
		case order.FieldTotalPrice:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field total_price", values[i])
			} else if value.Valid {
				o.TotalPrice = value.Float64
			}
		case order.FieldTotalCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field total_count", values[i])
			} else if value.Valid {
				o.TotalCount = int(value.Int64)
			}
		case order.FieldFinalTotalPrice:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field final_total_price", values[i])
			} else if value.Valid {
				o.FinalTotalPrice = value.Float64
			}
		case order.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				o.Status = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryOrderSnap queries the "order_snap" edge of the Order entity.
func (o *Order) QueryOrderSnap() *OrderSnapQuery {
	return (&OrderClient{config: o.config}).QueryOrderSnap(o)
}

// QueryOrderSub queries the "order_sub" edge of the Order entity.
func (o *Order) QueryOrderSub() *OrderSubQuery {
	return (&OrderClient{config: o.config}).QueryOrderSub(o)
}

// Update returns a builder for updating this Order.
// Note that you need to call Order.Unwrap() before calling this method if this Order
// was returned from a transaction, and the transaction was committed or rolled back.
func (o *Order) Update() *OrderUpdateOne {
	return (&OrderClient{config: o.config}).UpdateOne(o)
}

// Unwrap unwraps the Order entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (o *Order) Unwrap() *Order {
	tx, ok := o.config.driver.(*txDriver)
	if !ok {
		panic("model: Order is not a transactional entity")
	}
	o.config.driver = tx.drv
	return o
}

// String implements the fmt.Stringer.
func (o *Order) String() string {
	var builder strings.Builder
	builder.WriteString("Order(")
	builder.WriteString(fmt.Sprintf("id=%v", o.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(o.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(o.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", delete_time=")
	builder.WriteString(o.DeleteTime.Format(time.ANSIC))
	builder.WriteString(", order_no=")
	builder.WriteString(o.OrderNo)
	builder.WriteString(", transaction_id=")
	builder.WriteString(o.TransactionID)
	builder.WriteString(", user_id=")
	builder.WriteString(fmt.Sprintf("%v", o.UserID))
	builder.WriteString(", total_price=")
	builder.WriteString(fmt.Sprintf("%v", o.TotalPrice))
	builder.WriteString(", total_count=")
	builder.WriteString(fmt.Sprintf("%v", o.TotalCount))
	builder.WriteString(", final_total_price=")
	builder.WriteString(fmt.Sprintf("%v", o.FinalTotalPrice))
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", o.Status))
	builder.WriteByte(')')
	return builder.String()
}

// Orders is a parsable slice of Order.
type Orders []*Order

func (o Orders) config(cfg config) {
	for _i := range o {
		o[_i].config = cfg
	}
}
