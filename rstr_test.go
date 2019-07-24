/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : str.go
#   Created       : 2019-04-30 16:14
#   Last Modified : 2019-04-30 16:14
#   Describe      :
#
# ====================================================*/
package gotask

import (
	"testing"
)

const n = 20

func BenchmarkRandStringBytesMaskImprSrc(b *testing.B) {
	var temp = make(map[string]bool, b.N)
	for i := 0; i < b.N; i++ {
		temp[RandStringBytesMaskImprSrc(n)] = true
	}
	if len(temp) != b.N {
		b.Error("error len:", len(temp))
	}
}

func TestRandStringBytesMaskImprSrc(t *testing.T) {
	const length = 2000
	var temp = make(map[string]bool, length)
	for i := 0; i < length; i++ {
		temp[RandStringBytesMaskImprSrc(n)] = true
	}
	if len(temp) != length {
		t.Error("error len:", len(temp))
	}
}
