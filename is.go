// Copyright 2018 tinystack Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package goutil

import (
    "strings"
    "net"
)

func IsInteger(val interface{}) bool {
    switch val.(type) {
    case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
    case string:
        str := val.(string)
        if str == "" {
            return false
        }
        str = strings.TrimSpace(str)
		if str[0] == '-' || str[0] == '+' {
			if len(str) == 1 {
				return false
			}
			str = str[1:]
		}
        for _, v := range str {
            if v < '0' || v > '9' {
				return false
			}
        }
    }
    return true
}

func IsIp(s string) bool {
    ip := net.ParseIP(s)
    if ip == nil {
        return false
    }
    return true
}
