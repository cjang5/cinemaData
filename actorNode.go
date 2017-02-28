package cinemaData

import (
	"time"
)

// actorNode represents a graph node holding info about an actor
type actorNode struct {
	name      string
	birthdate time.Time
	movies    []string
}

// NewActor creates a new actorNode with zero-val in its attributes
func newActor() *actorNode {
	a := new(actorNode)
	a.name = ""
	a.birthdate = *new(time.Time)
	a.movies = make([]string, 0)

	return a
}

// SetName sets the actorNode's name to the given string
func (a *actorNode) setName(n string) {
	a.name = n
}

// SetBirthdate sets the actorNode's birthdate to the given time.Time struct
func (a *actorNode) setBirthdate(t time.Time) {
	a.birthdate = t
}

// AddMovie appends the given movie's Wikipedia URL to the actor's movie slice
func (a *actorNode) addMovie(u ...string) {
	a.movies = append(a.movies, u...)
}
