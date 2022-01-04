// Code generated by entc, DO NOT EDIT.

package model

import (
	"fmt"
	"mall-go/app/app/service/internal/data/model/charge"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Charge is the model entity for the Charge schema.
type Charge struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// DeleteTime holds the value of the "delete_time" field.
	DeleteTime time.Time `json:"delete_time,omitempty"`
	// UserID holds the value of the "user_id" field.
	// user表外键
	UserID int64 `json:"user_id,omitempty"`
	// Amount holds the value of the "amount" field.
	// 充值金额
	Amount string `json:"amount,omitempty"`
	// ChargeNo holds the value of the "charge_no" field.
	// 充值单号
	ChargeNo string `json:"charge_no,omitempty"`
	// TransactionID holds the value of the "transaction_id" field.
	// 支付平台流水号
	TransactionID string `json:"transaction_id,omitempty"`
	// PayWay holds the value of the "pay_way" field.
	// 支付方式：1微信支付，2支付宝支付
	PayWay int `json:"pay_way,omitempty"`
	// ClientType holds the value of the "client_type" field.
	// 客户端类型：1安卓，2IOS，3PC
	ClientType int `json:"client_type,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Charge) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case charge.FieldID, charge.FieldUserID, charge.FieldPayWay, charge.FieldClientType:
			values[i] = new(sql.NullInt64)
		case charge.FieldAmount, charge.FieldChargeNo, charge.FieldTransactionID:
			values[i] = new(sql.NullString)
		case charge.FieldCreateTime, charge.FieldUpdateTime, charge.FieldDeleteTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Charge", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Charge fields.
func (c *Charge) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case charge.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int64(value.Int64)
		case charge.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				c.CreateTime = value.Time
			}
		case charge.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				c.UpdateTime = value.Time
			}
		case charge.FieldDeleteTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field delete_time", values[i])
			} else if value.Valid {
				c.DeleteTime = value.Time
			}
		case charge.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				c.UserID = value.Int64
			}
		case charge.FieldAmount:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value.Valid {
				c.Amount = value.String
			}
		case charge.FieldChargeNo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field charge_no", values[i])
			} else if value.Valid {
				c.ChargeNo = value.String
			}
		case charge.FieldTransactionID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field transaction_id", values[i])
			} else if value.Valid {
				c.TransactionID = value.String
			}
		case charge.FieldPayWay:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field pay_way", values[i])
			} else if value.Valid {
				c.PayWay = int(value.Int64)
			}
		case charge.FieldClientType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field client_type", values[i])
			} else if value.Valid {
				c.ClientType = int(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Charge.
// Note that you need to call Charge.Unwrap() before calling this method if this Charge
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Charge) Update() *ChargeUpdateOne {
	return (&ChargeClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Charge entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Charge) Unwrap() *Charge {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("model: Charge is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Charge) String() string {
	var builder strings.Builder
	builder.WriteString("Charge(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(c.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(c.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", delete_time=")
	builder.WriteString(c.DeleteTime.Format(time.ANSIC))
	builder.WriteString(", user_id=")
	builder.WriteString(fmt.Sprintf("%v", c.UserID))
	builder.WriteString(", amount=")
	builder.WriteString(c.Amount)
	builder.WriteString(", charge_no=")
	builder.WriteString(c.ChargeNo)
	builder.WriteString(", transaction_id=")
	builder.WriteString(c.TransactionID)
	builder.WriteString(", pay_way=")
	builder.WriteString(fmt.Sprintf("%v", c.PayWay))
	builder.WriteString(", client_type=")
	builder.WriteString(fmt.Sprintf("%v", c.ClientType))
	builder.WriteByte(')')
	return builder.String()
}

// Charges is a parsable slice of Charge.
type Charges []*Charge

func (c Charges) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}