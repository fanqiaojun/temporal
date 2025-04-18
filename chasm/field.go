// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package chasm

import (
	"reflect"
	"strings"

	"google.golang.org/protobuf/proto"
)

const (
	// Used by reflection.
	chasmFieldTypePrefix      = "chasm.Field["
	chasmCollectionTypePrefix = "chasm.Collection["
	internalFieldName         = "Internal"

	fieldNameTag = "name"
)

// This struct needs to be create via reflection
// but reflection can't set prviate fields...
type Field[T any] struct {
	Internal fieldInternal
}

type fieldInternal struct {
	fieldType fieldType
	value     any // Component | Data | Pointer

	// Pointer to the corresponding tree node. Can be nil for the just assigned fields.
	node *Node
}

func (fi fieldInternal) IsEmpty() bool {
	return fi.value == nil
}

func (f Field[T]) Get(Context) (T, error) {
	// remember to handle d == nil case

	panic("not implemented")
}

func NewEmptyField[T any]() Field[T] {
	return Field[T]{}
}

// re. Data v.s. Component.
// Components have behavior and has a lifecycle.
// while Data doesn't and must be attached to a component.
//
// You can define a component just for storing the data,
// that may contain other information like ref count etc.
// most importantly, the framework needs to know when it's safe to delete the data.
// i.e. the lifecycle of that data component reaches completed.
func NewDataField[D proto.Message](
	ctx MutableContext,
	d D,
) Field[D] {
	return Field[D]{
		Internal: fieldInternal{
			fieldType: fieldTypeData,
			value:     d,
		},
	}
}

type componentFieldOptions struct {
	detached bool
}

type ComponentFieldOption func(*componentFieldOptions)

func ComponentFieldDetached() ComponentFieldOption {
	return func(o *componentFieldOptions) {
		o.detached = true
	}
}

func NewComponentField[C Component](
	ctx MutableContext,
	c C,
	options ...ComponentFieldOption,
) Field[C] {
	return Field[C]{
		Internal: fieldInternal{
			fieldType: fieldTypeComponent,
			value:     c,
		},
	}
}

func NewComponentPointerField[C Component](
	ctx MutableContext,
	c C,
) Field[C] {
	panic("not implemented")
}

func NewDataPointerField[D proto.Message](
	ctx MutableContext,
	d D,
) Field[D] {
	panic("not implemented")
}

type Collection[T any] map[string]Field[T]

func genericTypePrefix(t reflect.Type) string {
	tn := t.String()
	bracketPos := strings.Index(tn, "[")
	if bracketPos == -1 {
		return ""
	}
	return tn[:bracketPos+1]
}
