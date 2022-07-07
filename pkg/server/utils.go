package server

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/vmihailenco/msgpack"
	"io"
	"strconv"
	"sync/atomic"
	"time"
)

func GetId() string {
	value := int64(time.Now().Nanosecond()) << 32
	next := atomic.AddInt64(&value, 1)
	return strconv.FormatInt(next, 10)
}

func GetJobId(Handle, Params string) string {
	md5Ctx := md5.New()
	timeStr := strconv.FormatInt(int64(time.Now().Nanosecond()), 10)
	val := Handle + Params + timeStr
	//md5Ctx.Write([]byte(val))

	io.WriteString(md5Ctx, val)
	md5Str := md5Ctx.Sum(nil)

	return hex.EncodeToString(md5Str)
}

func GetBuffer(n int) (buf []byte) {
	buf = make([]byte, n)
	return
}

func IsMulParams(params []byte) bool {
	var decParams []string
	err := msgpack.Unmarshal(params, &decParams)
	if err == nil {
		plen := len(decParams)
		if plen > 1 {
			return true
		} else {
			return false
		}
	}

	return false
}
