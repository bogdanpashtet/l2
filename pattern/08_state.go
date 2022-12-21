package pattern

import "fmt"

type Document struct {
	state State
}

func (p *Document) setState(st State) {
	p.state = st
}

type State interface {
	publish()
}

type Draft struct {
	doc *Document
}

func (d *Draft) publish() {
	fmt.Printf("Отправлено на модерацию.\n")
	d.doc.setState(&Moderation{d.doc})
}

type Moderation struct {
	doc *Document
}

func (m *Moderation) publish() {
	fmt.Printf("Произошла модерация.\n")
	m.doc.setState(&Published{m.doc})
}

type Published struct {
	doc *Document
}

func (p *Published) publish() {
	fmt.Printf("Опубликовано.\n")
}

func StateFunc() {
	document := Document{}
	document.setState(&Draft{&document})

	document.state.publish()
	document.state.publish()
	document.state.publish()
}
