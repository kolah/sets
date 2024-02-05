package sets

// Set is a set of elements.
type Set[E comparable] map[E]struct{}

// New creates a new set with the given values.
func New[E comparable](values ...E) Set[E] {
	s := Set[E]{}
	for _, v := range values {
		s[v] = struct{}{}
	}
	return s
}

// Add adds an element to the set.
func (s Set[E]) Add(v E) Set[E] {
	s[v] = struct{}{}

	return s
}

// Remove removes an element from the set.
func (s Set[E]) Remove(v E) Set[E] {
	delete(s, v)

	return s
}

// Contains checks if the set contains the given element.
func (s Set[E]) Contains(v E) bool {
	_, ok := s[v]
	return ok
}

// Items returns a slice of the items in the set.
func (s Set[E]) Items() []E {
	result := make([]E, 0, len(s))
	for v := range s {
		result = append(result, v)
	}
	return result
}

// Intersection returns a new set with the intersection of the set and the given set.
func (s Set[E]) Intersection(s2 Set[E]) Set[E] {
	result := New[E]()
	for _, v := range s.Items() {
		if s2.Contains(v) {
			result.Add(v)
		}
	}

	return result
}

// Diff returns a new set with the difference of the set and the given set.
func (s Set[E]) Diff(s2 Set[E]) Set[E] {
	result := New[E]()
	for _, v := range s.Items() {
		if !s2.Contains(v) {
			result.Add(v)
		}
	}

	return result
}

// Intersects checks if the set intersects with the given set.
func (s Set[E]) Intersects(s2 Set[E]) bool {
	for _, v := range s.Items() {
		if s2.Contains(v) {
			return true
		}
	}

	return false
}

// Clone returns a new set with the same elements as the set.
func (s Set[E]) Clone() Set[E] {
	result := New[E]()
	for _, v := range s.Items() {
		result.Add(v)
	}

	return result
}

// Count returns the number of elements in the set.
func (s Set[E]) Count() int {
	return len(s)
}

// Flush removes all elements from the set.
func (s Set[E]) Flush() {
	clear(s)
}
