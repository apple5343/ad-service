package converter

import "database/sql"

func ToNullString(str *string) sql.NullString {
	if str != nil {
		return sql.NullString{String: *str, Valid: true}
	}
	return sql.NullString{}
}

func ToNullInt64(num *int) sql.NullInt64 {
	if num != nil {
		return sql.NullInt64{Int64: int64(*num), Valid: true}
	}
	return sql.NullInt64{}
}

func ToString(str sql.NullString) *string {
	if str.Valid {
		return &str.String
	}
	return nil
}

func ToInt(num sql.NullInt64) *int {
	if num.Valid {
		i := int(num.Int64)
		return &i
	}
	return nil
}

func FromStringPtr(str *string) string {
	if str != nil {
		return *str
	}
	return ""
}

func FromIntPtr(num *int) int {
	if num != nil {
		return *num
	}
	return 0
}

func FromFloat64Ptr(num *float64) float64 {
	if num != nil {
		return *num
	}
	return 0
}
