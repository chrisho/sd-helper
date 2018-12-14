package idcard

import (
	"github.com/chrisho/sd-helper/valid"
	"math"
	"strconv"
	"time"
)

func GetAge(idNo string) int {
	if !valid.IsIdcard(idNo) {
		return 0
	}

	idcardDateStr := string([]byte(idNo)[6:14])
	nowTimeDateStr := time.Now().Format("20060102")
	idcardDateInt, _ := strconv.Atoi(idcardDateStr)
	nowTimeDateInt, _ := strconv.Atoi(nowTimeDateStr)

	return int(math.Floor(float64(nowTimeDateInt-idcardDateInt)) / 10000)
}