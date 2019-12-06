package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type object struct {
	name   string
	parent *object
	child  []*object
}

func main() {
	input := make([]string, 0)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		input = append(input, text)
	}

	fmt.Println("Part one = ", partOne(input))
}

func partOne(input []string) int {
	var res int

	objects := make(map[string]*object)
	for _, v := range input {
		buildConnection(objects, v)
	}

	for _, v := range objects {
		res += goToCOM(objects, v)
	}

	return res
}

func buildConnection(objects map[string]*object, or string) {
	data := strings.Split(or, ")")

	around, ok := objects[data[0]]
	if !ok {
		around = &object{}
		around.name = data[0]
	}

	inOrbit, ok := objects[data[1]]
	if !ok {
		inOrbit = &object{
			name: data[1],
		}
	}

	inOrbit.parent = around

	around.child = append(around.child, inOrbit)

	objects[around.name] = around
	objects[inOrbit.name] = inOrbit

	for _, v := range objects {
		for _, c := range v.child {
			if c.name == around.name {
				around.parent = c.parent
			}
		}
	}
}

func goToCOM(objects map[string]*object, obj *object) int {
	if obj.name == "COM" {
		return 0
	}

	if obj.parent == nil {
		panic("does not have parent " + obj.name)
	}

	return 1 + goToCOM(objects, objects[obj.parent.name])
}
