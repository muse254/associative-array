package main

import (
	"errors"
	"fmt"
	"reflect"
)

type AssociativeArray struct {
	keys   []interface{}
	values []interface{}

	// keyType allows for type check to make sure all keys are of the same type
	keyType reflect.Type
	// valueType allows for type checks to tmake sure that all values are of the same type
	valueType reflect.Type
}

var (
	ErrKeyNotFound = errors.New("ERROR: the key has not been found in the AssociativeArray index")

	ErrKeyAlreadyExists = func(key interface{}) error {
		return errors.New(fmt.Sprintf("ERROR: the key %v already exists", reflect.ValueOf(key)))
	}

	ErrorWrongType = func(wrongType, correctType reflect.Type, place string) error {
		return errors.New(fmt.Sprintf("ERROR: the %s needs to be of type %s but provided %s", place, correctType, wrongType))
	}
)

func newAssociativeArray() *AssociativeArray {
	return &AssociativeArray{
		keys:   []interface{}{},
		values: []interface{}{},
	}
}

func (a *AssociativeArray) typeCheckPass(newKey, newValue interface{}) error {
	if reflect.TypeOf(newKey) != a.keyType {
		return ErrorWrongType(reflect.TypeOf(newKey), a.keyType, "type")
	} else if reflect.TypeOf(newValue) != a.valueType {
		return ErrorWrongType(reflect.TypeOf(newValue), a.valueType, "value")
	}
	return nil
}

func (a *AssociativeArray) addition(key, value interface{}) {
	// if the associative array is empty
	if len(a.keys) == 0 && len(a.values) == 0 {
		a.keys = append(a.keys, key)
		a.values = append(a.values, value)

		a.keyType = reflect.TypeOf(key)
		a.valueType = reflect.TypeOf(value)
		return
	}

	// type check the pair
	if err := a.typeCheckPass(key, value); err != nil {
		fmt.Println(err)
		return
	}

	// if the key already exists
	for _, otherKey := range a.keys {
		if otherKey == key {
			fmt.Println(ErrKeyAlreadyExists(key))
		}
	}

	// add the pair to the associative array
	a.keys = append(a.keys, key)
	a.values = append(a.values, value)
}

func (a *AssociativeArray) lookup(key string) (retKey, value interface{}, err error) {
	for i, currentKey := range a.keys {
		if currentKey == key {
			return key, a.values[i], nil
		}
	}
	return nil, nil, ErrKeyNotFound
}

func (a *AssociativeArray) remove(keyToRemove string) {
	for i, key := range a.keys {
		if keyToRemove == key {
			// remove operation
			if len(a.keys) > 1 && len(a.values) > 1 {
				a.keys = append(a.keys[:i], a.keys[i+1:]...)
				a.values = append(a.values[:i], a.values[i+1:]...)
			} else if len(a.keys) == 1 {
				a.keys = []interface{}{}
				a.values = []interface{}{}
			}
			return
		}
	}
	fmt.Println(ErrKeyNotFound)
}

func (a *AssociativeArray) printer() {
	var print string
	for i, key := range a.keys {
		print += fmt.Sprintf("[%v:%v] ", key, a.values[i])
	}
	fmt.Println(print)
}

func (a *AssociativeArray) modify(key, newValue interface{}) error {
	for i, currentKey := range a.keys {
		if currentKey == key {
			// modify operation
			a.values[i] = newValue
			return nil
		}
	}
	return ErrKeyNotFound
}

func main() {
	assArray := newAssociativeArray()

	// add 2 key value pairs
	// Shopping list
	assArray.addition("bread", 50)
	assArray.addition("soap", 20)

	assArray.printer()

	// lookup previously added pairs
	key, value, err := assArray.lookup("bread")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("value is: %v, for key: %v\n", value, key)

	// add 1 pair of wrong type, should flag up an error
	assArray.addition(37373.3737, "someRandomNum")

	// modify previously successfully added pair
	assArray.modify("bread", 100)

	assArray.printer()

	// delete a pair
	assArray.remove("bread")

	assArray.printer()
}
