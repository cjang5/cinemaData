package cinemaData

import "time"

// Graph is a graph representation of Actors associated with their movies
// and vice versa
type Graph struct {
	actors map[string]*actorNode
	movies map[string]*movieNode
}

// New creates and returns a new Graph
// the actors map and movies map will be empty upon creation
func New() *Graph {
	g := new(Graph)
	g.actors = make(map[string]*actorNode)
	g.movies = make(map[string]*movieNode)

	return g
}

// AddActor adds a new actorNode to our Graph
func (g *Graph) AddActor(url string, name string, birthdate time.Time, movies ...string) {
	a := newActor()
	a.setName(name)
	a.setBirthdate(birthdate)
	a.addMovie(movies...)

	g.actors[url] = a
}

// AddMovie adds a new movieNode to our Graph
func (g *Graph) AddMovie(url string, title string, release time.Time, boxOffice int, cast map[string]int) {
	m := newMovie()
	m.setTitle(title)
	m.setReleaseDate(release)
	m.setBoxOffice(boxOffice)

	for k, v := range cast {
		m.addCastMember(k, v)
	}

	g.movies[url] = m
}
