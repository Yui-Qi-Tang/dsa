// Package disjoints is a practice with the introduction to algorithm section 21.1 Disjoint set operations
//
//	this is a simple implemenation wit 'string slice' which means the element is the byte
package disjointsets

func New() *DisjointSets {
	return &DisjointSets{
		sets: make([]string, 0),
	}
}

// DisjointSets is an object for disjoint sets
type DisjointSets struct {
	// sets is the representative using string type
	// the collection is string which implies the elements if byte
	sets []string
}

// MakeSets makes the collections of disjoint sets
func (d *DisjointSets) MakeSets(bs ...byte) {
	for i := range bs {
		if d.FindSet(bs[i]) == -1 {
			d.sets = append(d.sets, string(bs[i]))
		}
	}
}

// Union makes union of u and v if u and v stands in differet pointer
func (d *DisjointSets) Union(u, v byte) {
	i, j := d.FindSet(u), d.FindSet(v)
	// return if u and v in the same group
	if i == j {
		return
	}

	// this is a tech. problems if the u or v doesn't exist in the sets
	if i == -1 || j == -1 {
		return
	}

	// keep the element with small index that is behind of the bigger one
	if i > j {
		i, j = j, i
	}
	d.sets[i] += string(d.sets[j])
	d.sets = append(d.sets[:j], d.sets[j+1:]...)
}

// FindSet finds the v in the sets, -1 if v does not exist in the sets
func (d *DisjointSets) FindSet(v byte) int {
	for i := range d.sets {
		for j := range d.sets[i] {
			if v == d.sets[i][j] {
				return i
			}
		}
	}
	return -1
}

// Collections returns the collections
func (d DisjointSets) Collections() []string {
	return d.sets
}
