package utils

import "time"

func PtrToInt(v *int) int {
	if v == nil {
		return 0
	}
	return *v
}

func PtrToInt32(v *int32) int32 {
	if v == nil {
		return 0
	}
	return *v
}

func PtrToInt64(v *int64) int64 {
	if v == nil {
		return 0
	}
	return *v
}

func PtrToString(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}

func PtrToBool(v *bool) bool {
	if v == nil {
		return false
	}
	return *v
}

func PtrToFloat32(v *float32) float32 {
	if v == nil {
		return 0
	}
	return *v
}

func PtrToFloat64(v *float64) float64 {
	if v == nil {
		return 0
	}
	return *v
}

func IntToPtr(v int) *int {
	return &v
}

func Int32ToPtr(v int32) *int32 {
	return &v
}

func Int64ToPtr(v int64) *int64 {
	return &v
}

func StringToPtr(v string) *string {
	return &v
}

func BoolToPtr(v bool) *bool {
	return &v
}

func Float32ToPtr(v float32) *float32 {
	return &v
}

func Float64ToPtr(v float64) *float64 {
	return &v
}

func TimeToPtr(v time.Time) *time.Time {
	return &v
}

func PtrToTime(v *time.Time) time.Time {
	if v == nil {
		return time.Time{}
	}
	return *v
}
