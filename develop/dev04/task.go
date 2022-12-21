package main

import (
	"bytes"
	"fmt"
	"sort"
)

/*
=== Поиск анаграмм по словарю ===
Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.
Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.
Программа должна проходить все тесты. Код должен проходить проверки go vet и go lint.
*/

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func FindAnogram(arr *[]string) *map[string]*[]string {
	var set = make(map[string]*[]string)
	var buffSet = make(map[string][]string)

	for _, val := range *arr {
		val = string(bytes.ToLower([]byte(val)))
		ss := SortString(val)
		buffSet[ss] = append(buffSet[ss], val)
	}

	for k, v := range buffSet {
		sort.Strings(buffSet[k])
		a := buffSet[k]
		set[v[0]] = &a
	}

	return &set
}

func main() {
	var strArr = []string{"тяпка", "пятак", "лиСток", "пЯтка", "слиТок", "столик"}
	set := FindAnogram(&strArr)
	for k, v := range *set {
		fmt.Printf("%s: %v", k, v)
	}
}
