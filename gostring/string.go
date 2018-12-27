// Copyright 2018 tinystack Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gostring

import (
    "strings"
    "strconv"
    "math/rand"
    "crypto/md5"
    "encoding/hex"
    "encoding/json"
    "time"
    "sort"
)

func JoinStrings(multiString ...string) string {
    return strings.Join(multiString, "")
}

func JoinSepStrings(sep string, multiString ...string) string {
    return strings.Join(multiString, sep)
}

func JoinIntSlice2String(intSlice []int, sep string) string {
    return strings.Join(IntSlice2StrSlice(intSlice), sep)
}

func StrSlice2IntSlice(strSlice []string) []int {
    var intSlice []int
    for _, s := range strSlice {
        i, _ := strconv.Atoi(s)
        intSlice = append(intSlice, i)
    }
    return intSlice
}

func StrSplit2IntSlice(str, sep string) []int {
    return StrSlice2IntSlice(StrFilterSliceEmpty(strings.Split(str, sep)))
}

func IntSlice2StrSlice(intSlice []int) []string {
    var strSlice []string
    for _, i := range intSlice {
        s := strconv.Itoa(i)
        strSlice = append(strSlice, s)
    }
    return strSlice
}

func StrFilterSliceEmpty(strSlice []string) []string {
    var filterSlice []string
    for _, s := range strSlice {
        if s != "" {
            filterSlice = append(filterSlice, s)
        }
    }
    return filterSlice
}

func Str2Int(s string) int {
    i, _ := strconv.Atoi(s)
    return i
}

func Int2Str(i int) string {
    return strconv.Itoa(i)
}

func StrRandom(l int) string {
    str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    bytes := []byte(str)
    result := []byte{}
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    for i := 0; i < l; i++ {
        result = append(result, bytes[r.Intn(len(bytes))])
    }
    return string(result)
}

func StrMd5(s string) string {
    md5Ctx := md5.New()
    md5Ctx.Write([]byte(s))
    return hex.EncodeToString(md5Ctx.Sum(nil))
}

func StringSliceRsort(s []string) []string {
    sort.Strings(s)
    for from, to := 0, len(s)-1; from < to; from, to = from + 1, to - 1 {
        s[from], s[to] = s[to], s[from]
    }
    return s
}

func JsonEncode(obj interface{}) (string, error) {
    jsonBytes, err := json.Marshal(obj)
    if err != nil {
        return "", err
    }
    return string(jsonBytes), nil
}
