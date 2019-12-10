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

	objects := make(map[string]*object)

	fmt.Println("Part one = ", partOne(input, objects))
	fmt.Println("Part two = ", partTwo(objects))
}

func partOne(input []string, objects map[string]*object) int {
	var res int

	for _, v := range input {
		buildConnection(objects, v)
	}

	for _, v := range objects {
		res += numberOfOrbits(objects, v)
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

func numberOfOrbits(objects map[string]*object, obj *object) int {
	if obj.name == "COM" {
		return 0
	}

	if obj.parent == nil {
		panic("does not have parent " + obj.name)
	}

	return 1 + numberOfOrbits(objects, objects[obj.parent.name])
}

func partTwo(objects map[string]*object) int {
	destinationChain := make(map[string]string)
	buildDestinationChain(objects, objects["SAN"], destinationChain)

	originTransferCount, originTransferOjectName :=
		transfersFromOrigin(objects, objects["YOU"], destinationChain)

	destinationTransferCount, destinationTransferObjectName :=
		transfersToDestination(destinationChain[originTransferOjectName], destinationChain)

	objects[destinationTransferObjectName].child = append(objects[destinationTransferObjectName].child, objects["YOU"])
	objects["YOU"].parent = objects[destinationTransferObjectName]

	return originTransferCount + destinationTransferCount
}

func buildDestinationChain(objects map[string]*object, obj *object, chain map[string]string) {
	if obj.name == "COM" {
		return
	}

	if obj.parent == nil {
		panic("does not have parent " + obj.name)
	}

	chain[obj.parent.name] = obj.name
	buildDestinationChain(objects, objects[obj.parent.name], chain)
}

func transfersFromOrigin(objects map[string]*object, obj *object, san map[string]string) (int, string) {
	if _, ok := san[obj.name]; ok {
		return 0, obj.name
	}

	if obj.parent == nil {
		panic("does not have parent " + obj.name)
	}

	i, objName := transfersFromOrigin(objects, objects[obj.parent.name], san)

	return i + 1, objName
}

func transfersToDestination(objName string, m map[string]string) (int, string) {
	v, ok := m[objName]
	if v == "SAN" {
		return 0, objName
	}

	if !ok {
		panic("does not have child " + objName)
	}

	i, objName := transfersToDestination(v, m)
	return 1 + i, objName
}
