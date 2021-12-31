// Code generated by entc, DO NOT EDIT.

package model

import (
	"fmt"
	"mall-go/internal/data/model/userpoint"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// UserPoint is the model entity for the UserPoint schema.
type UserPoint struct {
	config `json:"-"`
	// ID of the ent.
	// user_id
	ID int64 `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// DeleteTime holds the value of the "delete_time" field.
	DeleteTime time.Time `json:"delete_time,omitempty"`
	// Value holds the value of the "value" field.
	Value int `json:"value,omitempty"`
	// Status holds the value of the "status" field.
	Status int8 `json:"status,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserPoint) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case userpoint.FieldID, userpoint.FieldValue, userpoint.FieldStatus:
			values[i] = new(sql.NullInt64)
		case userpoint.FieldCreateTime, userpoint.FieldUpdateTime, userpoint.FieldDeleteTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type UserPoint", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserPoint fields.
func (up *UserPoint) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case userpoint.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			up.ID = int64(value.Int64)
		case userpoint.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				up.CreateTime = value.Time
			}
		case userpoint.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				up.UpdateTime = value.Time
			}
		case userpoint.FieldDeleteTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field delete_time", values[i])
			} else if value.Valid {
				up.DeleteTime = value.Time
			}
		case userpoint.FieldValue:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value.Valid {
				up.Value = int(value.Int64)
			}
		case userpoint.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				up.Status = int8(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this UserPoint.
// Note that you need to call UserPoint.Unwrap() before calling this method if this UserPoint
// was returned from a transaction, and the transaction was committed or rolled back.
func (up *UserPoint) Update() *UserPointUpdateOne {
	return (&UserPointClient{config: up.config}).UpdateOne(up)
}

// Unwrap unwraps the UserPoint entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (up *UserPoint) Unwrap() *UserPoint {
	tx, ok := up.config.driver.(*txDriver)
	if !ok {
		panic("model: UserPoint is not a transactional entity")
	}
	up.config.driver = tx.drv
	return up
}

// String implements the fmt.Stringer.
func (up *UserPoint) String() string {
	var builder strings.Builder
	builder.WriteString("UserPoint(")
	builder.WriteString(fmt.Sprintf("id=%v", up.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(up.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(up.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", delete_time=")
	builder.WriteString(up.DeleteTime.Format(time.ANSIC))
	builder.WriteString(", value=")
	builder.WriteString(fmt.Sprintf("%v", up.Value))
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", up.Status))
	builder.WriteByte(')')
	return builder.String()
}

// UserPoints is a parsable slice of UserPoint.
type UserPoints []*UserPoint

func (up UserPoints) config(cfg config) {
	for _i := range up {
		up[_i].config = cfg
	}
}
