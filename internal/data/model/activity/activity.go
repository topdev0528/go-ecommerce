// Code generated by entc, DO NOT EDIT.

package activity

import (
	"time"
)

const (
	// Label holds the string label denoting the activity type in the database.
	Label = "activity"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldDeleteTime holds the string denoting the delete_time field in the database.
	FieldDeleteTime = "delete_time"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldStartTime holds the string denoting the start_time field in the database.
	FieldStartTime = "start_time"
	// FieldEndTime holds the string denoting the end_time field in the database.
	FieldEndTime = "end_time"
	// FieldRemark holds the string denoting the remark field in the database.
	FieldRemark = "remark"
	// FieldOnline holds the string denoting the online field in the database.
	FieldOnline = "online"
	// FieldEntranceImg holds the string denoting the entrance_img field in the database.
	FieldEntranceImg = "entrance_img"
	// FieldInternalTopImg holds the string denoting the internal_top_img field in the database.
	FieldInternalTopImg = "internal_top_img"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeSpu holds the string denoting the spu edge name in mutations.
	EdgeSpu = "spu"
	// EdgeCoupon holds the string denoting the coupon edge name in mutations.
	EdgeCoupon = "coupon"
	// Table holds the table name of the activity in the database.
	Table = "activity"
	// SpuTable is the table that holds the spu relation/edge. The primary key declared below.
	SpuTable = "activity_spu"
	// SpuInverseTable is the table name for the Spu entity.
	// It exists in this package in order to avoid circular dependency with the "spu" package.
	SpuInverseTable = "spu"
	// CouponTable is the table that holds the coupon relation/edge. The primary key declared below.
	CouponTable = "activity_coupon"
	// CouponInverseTable is the table name for the Coupon entity.
	// It exists in this package in order to avoid circular dependency with the "coupon" package.
	CouponInverseTable = "coupon"
)

// Columns holds all SQL columns for activity fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldDeleteTime,
	FieldTitle,
	FieldDescription,
	FieldStartTime,
	FieldEndTime,
	FieldRemark,
	FieldOnline,
	FieldEntranceImg,
	FieldInternalTopImg,
	FieldName,
}

var (
	// SpuPrimaryKey and SpuColumn2 are the table columns denoting the
	// primary key for the spu relation (M2M).
	SpuPrimaryKey = []string{"activity_id", "spu_id"}
	// CouponPrimaryKey and CouponColumn2 are the table columns denoting the
	// primary key for the coupon relation (M2M).
	CouponPrimaryKey = []string{"activity_id", "coupon_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
)
