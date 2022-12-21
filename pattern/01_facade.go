package pattern

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// asmFile - .asm файл
type asmFile struct {
	name string
	code string
}

// TextEditor - подсистема: текстовый редактор
type TextEditor struct {
}

func (te *TextEditor) WriteCode(name, code string) {
	fmt.Printf("Пишем код: %s\n", code)
	te.SaveCode(name)
}

func (te *TextEditor) SaveCode(name string) {
	fmt.Printf("Сохраняем код в файле: %s\n", name)
}

// Compiler - подсистема: Компилятор
type Compiler struct {
}

func (c *Compiler) CompileCode(name string) string {
	var res strings.Builder
	fmt.Printf("tasm %s\n", name)
	res.WriteString(name[:utf8.RuneCountInString(name)-4])
	res.WriteString(".obj")
	fmt.Printf("[System] Создан файл %s\n", res.String())
	return res.String()
}

// Linker - подсистема: Компоновщик
type Linker struct {
}

func (l *Linker) LinkObjFile(objFileName string) string {
	var res strings.Builder
	fmt.Printf("tlink %s\n", objFileName)
	res.WriteString(objFileName[:utf8.RuneCountInString(objFileName)-4])
	res.WriteString(".exe")
	fmt.Printf("[System] Создан файл %s\n", res.String())
	return res.String()
}

// IDE - фасад, в котором собираются все наши подсистемы воедино
type IDE struct {
	TextEditor
	Compiler
	Linker
}

// BuildProgram - соединение всей бизнес-логики в одной единственной функции фасада IDE
func (ide *IDE) BuildProgram(asm asmFile) {
	ide.WriteCode(asm.name, asm.code)
	objFile := ide.CompileCode(asm.name)
	exeFile := ide.LinkObjFile(objFile)
	fmt.Printf("%s\n[System] Программа запущена...", exeFile)
}

// Facade - test code
func Facade() {
	var ide IDE
	asm := asmFile{
		name: "prog1.asm",
		code: "\n;...\nmov a, 4\nadd 5\n;...",
	}

	ide.BuildProgram(asm)
}
