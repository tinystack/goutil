// Copyright 2018 tinystack Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package goutil

import (
    "strings"
)

func JoinStrings(multiString ...string) string {
    return strings.Join(multiString, "")
}

