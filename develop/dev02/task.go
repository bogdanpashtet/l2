package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
=== Задача на распаковку ===
Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)
В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.
Функция должна проходить все тесты. Код должен проходить проверки go vet и go lint.
*/

func WriteValuesIfBufferNotEmpty(countStr, resultStr *strings.Builder, buffRune *rune) {
	// write "buffRune" "countStr" times in "resultStr" if "buffRune" isn't empty
	if *buffRune != ' ' {
		count, _ := strconv.Atoi(countStr.String())

		if count == 0 {
			count = 1
		}

		for l := 0; l < count; l++ {
			resultStr.WriteRune(*buffRune)
		}
		*buffRune = ' '
		countStr.Reset()
	}
}

func UnpackString(str string) string {

	if str == "" {
		return ""
	}

	runeStr := []rune(str)

	var strBuild, countBuff strings.Builder
	buffRune := ' '

	for i := 0; i < len(runeStr); i++ {
		elemStr := string(runeStr[i])
		switch {
		case (elemStr < "0" || elemStr > "9") && elemStr != `\`:
			WriteValuesIfBufferNotEmpty(&countBuff, &strBuild, &buffRune)
			buffRune = runeStr[i]
		case "0" < elemStr && elemStr < "9" && buffRune != ' ':
			countBuff.WriteRune(runeStr[i])
		case elemStr == `\` && i+1 < len(runeStr):
			WriteValuesIfBufferNotEmpty(&countBuff, &strBuild, &buffRune)
			i++
			buffRune = runeStr[i]
		}
	}

	WriteValuesIfBufferNotEmpty(&countBuff, &strBuild, &buffRune)
	if strBuild.Len() == 0 {
		return "(некорректная строка)"
	}

	return strBuild.String()
}

func main() {
	str := `45`
	fmt.Println("Unpacked string: ", UnpackString(str))
}
