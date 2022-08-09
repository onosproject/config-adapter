// SPDX-FileCopyrightText: 2022-present Intel Corporation
// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package synchronizer

import (
	"fmt"
)

// BoolToUint32 convert a boolean to an unsigned integer
func BoolToUint32(b bool) uint32 {
	if b {
		return 1
	}
	return 0
}

// DerefStrPtr dereference a string pointer, returning a default if it is nil
func DerefStrPtr(s *string, def string) string {
	if s == nil {
		return def
	}
	return *s
}

// DerefUint32Ptr dereference a uint32 pointer, returning default if it is nil
func DerefUint32Ptr(u *uint32, def uint32) uint32 {
	if u == nil {
		return def
	}
	return *u
}

// DerefUint16Ptr dereference a uint16 pointer, returning default if it is nil
func DerefUint16Ptr(u *uint16, def uint16) uint16 {
	if u == nil {
		return def
	}
	return *u
}

// DerefUint8Ptr dereference a uint8 pointer, returning default if it is nil
func DerefUint8Ptr(u *uint8, def uint8) uint8 {
	if u == nil {
		return def
	}
	return *u
}

// DerefInt8Ptr dereference an int8 pointer, returning default if it is nil
func DerefInt8Ptr(u *int8, def int8) int8 {
	if u == nil {
		return def
	}
	return *u
}

// ProtoStringToProtoNumber converts a protocol name to a number
func ProtoStringToProtoNumber(s string) (uint8, error) {
	n, okay := map[string]uint8{"TCP": 6, "UDP": 17}[s]
	if !okay {
		return 0, fmt.Errorf("Unknown protocol %s", s)
	}
	return n, nil
}

// AStr facilitates easy declaring of pointers to strings
func AStr(s string) *string {
	return &s
}

// ABool facilitates easy declaring of pointers to bools
func ABool(b bool) *bool {
	return &b
}

// AInt8 facilitates easy declaring of pointers to int8
func AInt8(u int8) *int8 {
	return &u
}

// AUint8 facilitates easy declaring of pointers to uint8
func AUint8(u uint8) *uint8 {
	return &u
}

// AUint16 facilitates easy declaring of pointers to uint16
func AUint16(u uint16) *uint16 {
	return &u
}

// AUint32 facilitates easy declaring of pointers to uint32
func AUint32(u uint32) *uint32 {
	return &u
}

// AUint64 facilitates easy declaring of pointers to uint64
func AUint64(u uint64) *uint64 {
	return &u
}
