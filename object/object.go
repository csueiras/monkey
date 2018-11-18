package object

import (
	"bytes"
	"fmt"
	"github.com/csueiras/monkey/ast"
	"hash/fnv"
	"strings"
)

type ObjectType string

const (
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
	FUNCTION_OBJ     = "FUNCTION"
	STRING_OBJ       = "STRING"
	BUILTIN_OBJ      = "BUILTIN"
	ARRAY_OBJ        = "ARRAY"
	HASH_OBJ         = "HASH"
	SET_OBJ          = "SET"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int64
	hash  *HashKey
}

func (i *Integer) Type() ObjectType { return INTEGER_OBJ }
func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }

type Boolean struct {
	Value bool
	hash  *HashKey
}

func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }

type Null struct{}

func (n *Null) Type() ObjectType { return NULL_OBJ }
func (n *Null) Inspect() string  { return "null" }

type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }
func (rv *ReturnValue) Inspect() string  { return rv.Value.Inspect() }

type Error struct {
	Message string
}

func (e *Error) Type() ObjectType { return ERROR_OBJ }
func (e *Error) Inspect() string  { return "ERROR: " + e.Message }

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() ObjectType { return FUNCTION_OBJ }
func (f *Function) Inspect() string {
	var out bytes.Buffer

	var params []string
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ","))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}

type String struct {
	Value string
	hash  *HashKey
}

func (s *String) Type() ObjectType { return STRING_OBJ }
func (s *String) Inspect() string  { return s.Value }

type BuiltinFunction func(args ...Object) Object

type Builtin struct {
	Fn BuiltinFunction
}

func (b *Builtin) Type() ObjectType { return BUILTIN_OBJ }
func (b *Builtin) Inspect() string  { return "builtin function" }

type Array struct {
	Elements []Object
}

func (ao *Array) Type() ObjectType { return ARRAY_OBJ }
func (ao *Array) Inspect() string {
	var out bytes.Buffer

	elements := []string{}
	for _, e := range ao.Elements {
		elements = append(elements, e.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ","))
	out.WriteString("]")

	return out.String()
}

type HashKey struct {
	Type  ObjectType
	Value uint64
}

func (b *Boolean) HashKey() HashKey {
	if b.hash == nil {
		var value uint64

		if b.Value {
			value = 1
		} else {
			value = 0
		}

		b.hash = &HashKey{Type: b.Type(), Value: value}
	}
	return *b.hash
}

func (i *Integer) HashKey() HashKey {
	if i.hash == nil {
		i.hash = &HashKey{Type: i.Type(), Value: uint64(i.Value)}
	}
	return *i.hash
}

func (s *String) HashKey() HashKey {
	if s.hash == nil {
		h := fnv.New64a()
		h.Write([]byte(s.Value))
		s.hash = &HashKey{Type: s.Type(), Value: h.Sum64()}
	}

	return *s.hash
}

type HashPair struct {
	Key   Object
	Value Object
}

type Hash struct {
	Pairs map[HashKey]HashPair
}

func (h *Hash) Type() ObjectType { return HASH_OBJ }
func (h *Hash) Inspect() string {
	var out bytes.Buffer

	var pairs []string
	for _, pair := range h.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s", pair.Key.Inspect(), pair.Value.Inspect()))
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}

type Hashable interface {
	HashKey() HashKey
}

type Set struct {
	Data map[HashKey]HashPair
}

func NewSet(elements []Object) *Set {
	pairs := make(map[HashKey]HashPair)
	for _, element := range elements {
		hashable, ok := element.(Hashable)
		if !ok {
			return nil
		}

		hashKey := hashable.HashKey()
		pairs[hashKey] = HashPair{
			Key:   element,
			Value: nil,
		}
	}
	return &Set{
		Data: pairs,
	}
}

func (s *Set) Type() ObjectType { return SET_OBJ }
func (s *Set) Inspect() string {
	var out bytes.Buffer

	var elements []string
	for _, e := range s.Elements() {
		elements = append(elements, e.Inspect())
	}

	out.WriteString("set")
	out.WriteString("(")
	out.WriteString(strings.Join(elements, ","))
	out.WriteString(")")

	return out.String()
}

func (s *Set) Elements() []Object {
	if s.Data == nil || len(s.Data) == 0 {
		return []Object{}
	}

	var elements []Object

	for _, v := range s.Data {
		elements = append(elements, v.Key)
	}
	return elements
}
