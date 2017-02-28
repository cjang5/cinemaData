package cinemaData

import "time"

// movieNode represents a node in our Graph implementation, specifically for a movie
type movieNode struct {
	title       string
	releaseDate time.Time
	boxOffice   int
	cast        map[string]int
}

// newMovie creates an empty movieNode
func newMovie() *movieNode {
	m := new(movieNode)
	m.title = ""
	m.releaseDate = *new(time.Time)
	m.boxOffice = 0
	m.cast = make(map[string]int)

	return m
}

// setTitle sets the movieNode's title to the given string
func (m *movieNode) setTitle(t string) {
	m.title = t
}

// getTitle returns the movieNode's title
func (m *movieNode) getTitle() string {
	return m.title
}

// setReleaseDate sets the movieNode's releaseDate to the given time.Time struct
func (m *movieNode) setReleaseDate(t time.Time) {
	m.releaseDate = t
}

// getReleaseDate returns the movieNode's releaseDate (time.Time)
func (m *movieNode) getReleaseDate() time.Time {
	return m.releaseDate
}

// setBoxOffice sets the movieNode's boxOffice to the given value (int)
func (m *movieNode) setBoxOffice(b int) {
	m.boxOffice = b
}

// getBoxOffice returns the movieNode's boxOffice (int)
func (m *movieNode) getBoxOffice() int {
	return m.boxOffice
}

// addCastMember adds a new cast member -> gross value (string -> int) key-value pair
// to the movieNode's cast map (map[string]int)
func (m *movieNode) addCastMember(s string, i int) {
	m.cast[s] = i
}
