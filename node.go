package main

import (
	"errors"
	"strconv"
)

var (
	ErrNotFound  = errors.New("NOT_FOUND")
	ErrWrongPath = errors.New("WRONG_PATH")
)

type NMType byte

const (
	NMDataType NMType = iota
	NMMapType
	NMSliceType
)

type DataNode struct {
	Type  NMType
	Map   map[string]*DataNode
	Slice []*DataNode
	Value interface{}
}

//Field1[*raw][0].Field2[0].Field3[*new]",
// Field1,*raw,0,Field2,0,Field3,*new
func (n *DataNode) Field(path []string) (interface{}, error) {
	switch n.Type {
	case NMDataType:
		if len(path) != 0 {
			return nil, ErrNotFound
		}
		return n.Value, nil
	case NMMapType:
		if len(path) == 0 {
			return n.Map, nil
		}
		node, has := n.Map[path[0]]
		if !has {
			return nil, ErrNotFound
		}
		return node.Field(path[1:])
	case NMSliceType:
		if len(path) == 0 {
			return n.Slice, nil
		}
		idx, err := strconv.Atoi(path[0])
		if err != nil {
			return nil, err
		}
		if idx < 0 {
			idx += len(n.Slice)
		}
		if idx < 0 || idx >= len(n.Slice) {
			return nil, ErrNotFound
		}
		return n.Slice[idx].Field(path[1:])
	}
	return nil, ErrWrongPath
}

func (n *DataNode) Set(val interface{}, path []string) error {
	if len(path) == 0 {
		switch v := val.(type) {
		case map[string]*DataNode:
			n.Type = NMMapType
			n.Map = v
		case []*DataNode:
			n.Type = NMSliceType
			n.Slice = v
		default:
			n.Type = NMDataType
			n.Value = val
		}
		return nil
	}
	switch n.Type {
	case NMDataType:
		return ErrWrongPath
	case NMMapType:
		node, has := n.Map[path[0]]
		if !has {
			node = NewDataNode(path[1:])
			n.Map[path[0]] = node
		}
		return node.Set(val, path[1:])
	case NMSliceType:
		idx, err := strconv.Atoi(path[0])
		if err != nil {
			return err
		}
		if idx == len(n.Slice) {
			node := NewDataNode(path[1:])
			n.Slice = append(n.Slice, node)
			return node.Set(val, path[1:])
		}
		if idx < 0 {
			idx += len(n.Slice)
		}
		if idx > len(n.Slice) {
			return ErrNotFound
		}
		return n.Slice[idx].Set(val, path[1:])
	}
	return ErrWrongPath
}

func (n DataNode) IsEmpty() bool {
	return n.Value == nil ||
		len(n.Map) == 0 ||
		len(n.Slice) == 0
}

func (n *DataNode) Remove(path []string) error {
	if len(path) == 0 {
		n.Map = nil
		n.Slice = nil
		n.Value = nil
		return nil
	}
	switch n.Type {
	case NMDataType:
		return ErrWrongPath
	case NMMapType:
		node, has := n.Map[path[0]]
		if !has {
			return nil
		}
		err := node.Remove(path[1:])
		if node.IsEmpty() {
			delete(n.Map, path[0])
		}
		return err
	case NMSliceType:
		idx, err := strconv.Atoi(path[0])
		if err != nil {
			return err
		}
		if idx < 0 {
			idx += len(n.Slice)
		}
		if idx >= len(n.Slice) {
			return nil
		}
		err = n.Slice[idx].Remove(path[1:])
		if n.Slice[idx].IsEmpty() {
			n.Slice = append(n.Slice[:idx], n.Slice[idx+1:]...)
		}
		return err
	}
	return nil
}

func NewDataNode(path []string) (n *DataNode) {
	n = new(DataNode)
	if len(path) == 0 {
		return
	}
	obj := NewDataNode(path[1:])
	if path[0] == "0" { // only support the 0 index when creating new array
		n.Type = NMSliceType
		n.Slice = []*DataNode{obj}
		return
	}
	n.Type = NMMapType
	n.Map = map[string]*DataNode{path[0]: obj}
	return
}
