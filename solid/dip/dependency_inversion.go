package main

import (
	"fmt"
)

// Dependency Inversion Principle.
// High Level Module (HLM) not depend on Low Level Module (LLM).
// Both should depend on abstractions.

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	name string
}

type Info struct {
	from         *Person
	relationship Relationship
	to           *Person
}

type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}

type RelationshipInfo struct {
	relations []Info
}

func (r *RelationshipInfo) FindAllChildrenOf(name string) []*Person {
	var results []*Person

	for _, v := range r.relations {
		if v.relationship == Parent && v.from.name == name {
			results = append(results, v.to)
		}
	}

	return results
}

func (r *RelationshipInfo) AddParentAndChild(parent, child *Person) {
	parentInfo := Info{
		relationship: Parent,
		from:         parent,
		to:           child,
	}
	childInfo := Info{
		relationship: Child,
		from:         child,
		to:           parent,
	}
	r.relations = append(r.relations, parentInfo, childInfo)
}

type Research struct {
	browser RelationshipBrowser
}

func (r *Research) Investigate() {
	for _, p := range r.browser.FindAllChildrenOf("John") {
		fmt.Println("John has a child called: ", p.name)
	}
}

func main() {
	parent1 := Person{"John"}
	child1 := Person{"Chris"}
	child2 := Person{"Matt"}

	// low-level module.
	relationship := RelationshipInfo{}
	relationship.AddParentAndChild(&parent1, &child1)
	relationship.AddParentAndChild(&parent1, &child2)

	research := Research{&relationship}
	research.Investigate()
}
