package main

import (
	"fmt"
	"reflect"
)

var otherThing float32 = -32.34

type Thing interface {
	AMethod(o string) string
	BMethod(i int) int
}

type AThing struct {
	AnotherVal int
}

type BThing struct {
	Val int
}

func (a AThing) AMethod(val string) string {
	return fmt.Sprintf("%s AThing", val)
}

func (b BThing) AMethod(val string) string {
	return fmt.Sprintf("%s Bthing", val)
}

func (a AThing) BMethod(i int) int {
	return 1 * i
}

func (b BThing) BMethod(i int) int {
	return 2 * i
}

func init() {
}

func main() {
	a := AThing{}
	b := BThing{}
	MethodThatTakesInterface(a)
	MethodThatTakesInterface(b)
}

func doAnotherThing(float float32) string {
	return fmt.Sprintf("%f - is the %x", float, reflect.TypeOf(float))
}

func MethodThatTakesInterface(thing Thing) {
	fmt.Println(fmt.Sprintf("%s %+v", thing.AMethod("value"), thing))
}

func UseMultiplyAndReturn() int {
	result := multiplyAndReturn(2)

	return result * 4
}

func multiplyAndReturn(number int) int {
	var multiplier int
	multiplier = 2
	var newValue int
	newValue = number * multiplier
	return newValue
}
