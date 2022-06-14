package formatter

import (
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
)

var cppHeaderBlockFmt = `
/*%s*/
/********************************************************
*
********************************************************/

/*%s*/
`

var cppFuncBlockFmt = `
/* %s                                                  */
/********************************************************
*description：函数详细描述
*arg        ：%s
*return     ：type %s; %s
********************************************************/
`

func CppHeaderBlockFmtFormatter(blockName string) string {
	log.Debug("[var]blockName:", blockName)
	splitStr := strings.Split(blockName, "\r\n")[0]
	log.Debug("[var]splitStr:", splitStr)
	blockNameLen := len(splitStr)
	log.Debug("[var]blockNameLen:", blockNameLen)

	excluedLen := 54 - blockNameLen
	blockNameStr := ""
	endExcluedLen := 54 - (blockNameLen + 6) // plus _End
	endBlockStr := " " + splitStr + " End "
	if (blockNameLen % 2) == 0 {
		needLen := (excluedLen / 2) - 1
		nameStr01 := "_"
		nameStr02 := ""
		for i := 0; i < needLen; i++ {
			nameStr01 += "_"
			nameStr02 += "_"
		}
		blockNameStr = nameStr01 + splitStr + nameStr02

		nameStr01 = "-"
		nameStr02 = ""
		for i := 0; i < (endExcluedLen/2)-1; i++ {
			nameStr01 += "-"
			nameStr02 += "-"
		}
		endBlockStr = nameStr01 + endBlockStr + nameStr02
	} else {
		needLen := (excluedLen / 2) - 1
		nameStr01 := ""
		nameStr02 := ""
		for i := 0; i < needLen; i++ {
			nameStr01 += "_"
			nameStr02 += "_"
		}
		blockNameStr = nameStr01 + " " + splitStr + " " + nameStr02

		nameStr01 = ""
		nameStr02 = ""
		for i := 0; i < (endExcluedLen / 2); i++ {
			nameStr01 += "-"
			nameStr02 += "-"
		}
		endBlockStr = nameStr01 + endBlockStr + nameStr02
	}
	log.Debug("[Var]blockNameStr:", blockNameStr)
	log.Debug("[Var]endBlockStr:", endBlockStr)

	lastStr := fmt.Sprintf(cppHeaderBlockFmt, blockNameStr, endBlockStr)
	return lastStr
}

func CppFuncBlockFmtFormatter() string {
	return cppFuncBlockFmt
}
