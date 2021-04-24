package pornhubEvaler

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	emptyDataReplacer   = strings.NewReplacer(" ", "", "	", "", "\r\n", "")
	removeCommentsRegex = regexp.MustCompile(`\/\*[\s\S]*?\*\/|\/\/.*`)
	pRegex              = regexp.MustCompile(`varp=(\d+);`)
	sRegex              = regexp.MustCompile(`vars=(\d+);`)
	shiftRegex          = regexp.MustCompile(`if\(\(s>>(\d+)\)`)
	pPlusRegex          = regexp.MustCompile(`p\+=(\d+)\*(\d+);`)
	pMinusRegex         = regexp.MustCompile(`p-=(\d+)\*(\d+);`)
	pLastRegex          = regexp.MustCompile(`p[-|\+]=(\d+);`)
	cookieNumberRegex   = regexp.MustCompile(`":"\+s\+"([\d\s:;]+)"`)
)

// Evalute given JS data(not in base64)
// Returns RNKEY cookie
func Eval(data string) string {
	data = emptyDataReplacer.Replace(removeCommentsRegex.ReplaceAllString(data, ""))
	pValue, _ := strconv.Atoi(pRegex.FindStringSubmatch(data)[1])
	sValue, _ := strconv.Atoi(sRegex.FindStringSubmatch(data)[1])
	shiftValues := sliceStringToUint8(shiftRegex.FindAllStringSubmatch(data, -1))
	pPlusValues := sliceStringToInt(pPlusRegex.FindAllStringSubmatch(data, -1))
	pMinusValues := sliceStringToInt(pMinusRegex.FindAllStringSubmatch(data, -1))
	cookieNumber := cookieNumberRegex.FindStringSubmatch(data)[1]
	pLastValueStr := pLastRegex.FindStringSubmatch(data)
	pLastValue, _ := strconv.Atoi(pLastValueStr[0])
	if ((sValue >> shiftValues[0]) & 1) == 1 {
		pValue += pPlusValues[0][0] * pPlusValues[0][1]
	} else {
		pValue -= pMinusValues[0][0] * pMinusValues[0][1]
	}

	if ((sValue >> shiftValues[1]) & 1) == 1 {
		pValue += pPlusValues[1][0] * pPlusValues[1][1]
	} else {
		pValue -= pMinusValues[1][0] * pMinusValues[1][1]
	}

	if ((sValue >> shiftValues[2]) & 1) == 1 {
		pValue += pPlusValues[2][0] * pPlusValues[2][1]
	} else {
		pValue -= pMinusValues[2][0] * pMinusValues[2][1]
	}

	if ((sValue >> shiftValues[3]) & 1) == 1 {
		pValue += pPlusValues[3][0] * pPlusValues[3][1]
	} else {
		pValue -= pMinusValues[3][0] * pMinusValues[3][1]
	}

	if ((sValue >> shiftValues[4]) & 1) == 1 {
		pValue += pPlusValues[4][0] * pPlusValues[4][1]
	} else {
		pValue -= pMinusValues[4][0] * pMinusValues[4][1]
	}

	if strings.Contains(pLastValueStr[0], "-=") {
		pValue -= pLastValue
	} else {
		pValue += pLastValue
	}
	n := leastFactor(pValue)
	return fmt.Sprintf("RNKEY=%d*%d:%d%s", n, pValue/n, sValue, cookieNumber)
}
