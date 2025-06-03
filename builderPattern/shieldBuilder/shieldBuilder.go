package shieldbuilder

import "strings"

type shield struct {
	front bool
	back  bool
	left  bool
	right bool
}

type shieldBuilder struct {
	code string
}

func NewShieldBuilder() *shieldBuilder {
	return new(shieldBuilder)
}

func (sb *shieldBuilder) RaiseFront() *shieldBuilder {
	sb.code += "F"
	return sb
}

func (sb *shieldBuilder) RaiseBack() *shieldBuilder {
	sb.code += "B"
	return sb
}

func (sb *shieldBuilder) RaiseRight() *shieldBuilder {
	sb.code += "R"
	return sb
}

func (sb *shieldBuilder) RaiseLeft() *shieldBuilder {
	sb.code += "L"
	return sb
}

func (sb *shieldBuilder) Build() *shield {
	code := sb.code
	return &shield{
		front: strings.Contains(code, "F"),
		back:  strings.Contains(code, "B"),
		left:  strings.Contains(code, "L"),
		right: strings.Contains(code, "R"),
	}
}
