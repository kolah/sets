package sets

// Trackable is a set that tracks changes. It keeps track of the elements that have been added and removed
// during its lifetime. It might be useful when interacting with a database, where you want to keep track of
// the changes that have been made to a set of elements and persist them accordingly.
type Trackable[E comparable] struct {
	set Set[E]

	toBeAdded   Set[E]
	toBeRemoved Set[E]
}

// NewTrackable creates a new trackable set with the given values.
func NewTrackable[E comparable](values ...E) Trackable[E] {
	return Trackable[E]{
		set:         New[E](values...),
		toBeAdded:   New[E](),
		toBeRemoved: New[E](),
	}
}

// Add adds an element to the set and marks it as to-be-added.
func (s Trackable[E]) Add(v E) Trackable[E] {
	s.set.Add(v)
	s.toBeAdded.Add(v)

	return s
}

// Remove removes an element from the set and marks it as to-be-removed.
func (s Trackable[E]) Remove(v E) Trackable[E] {
	if !s.set.Contains(v) {
		return s
	}

	s.set.Remove(v)
	s.toBeRemoved.Add(v)

	return s
}

// Contains checks if the set contains the given element.
func (s Trackable[E]) Contains(v E) bool {
	return s.set.Contains(v)
}

// Items returns a slice of the items in the set.
func (s Trackable[E]) Items() []E {
	return s.set.Items()
}

// Intersection returns a new set with the intersection of the set and the given set.
func (s Trackable[E]) Intersection(s2 Trackable[E]) Trackable[E] {
	items := s.set.Intersection(s2.set)

	return NewTrackable[E](items.Items()...)
}

// Intersects checks if the set intersects with the given set.
func (s Trackable[E]) Intersects(s2 Trackable[E]) bool {
	for _, v := range s.Items() {
		if s2.Contains(v) {
			return true
		}
	}

	return false
}

// Diff returns a new set with the difference of the set and the given set.
func (s Trackable[E]) Diff(s2 Trackable[E]) Trackable[E] {
	diff := s.set.Diff(s2.set)

	return NewTrackable[E](diff.Items()...)
}

// ToBeAdded returns a set of elements that have been added to the set.
func (s Trackable[E]) ToBeAdded() Set[E] {
	return s.toBeAdded
}

// ToBeRemoved returns a set of elements that have been removed from the set.
func (s Trackable[E]) ToBeRemoved() Set[E] {
	return s.toBeRemoved
}

// Clone returns a new trackable set with the same elements as the set.
func (s Trackable[E]) Clone() Trackable[E] {
	return Trackable[E]{
		set:         s.set.Clone(),
		toBeAdded:   s.toBeAdded.Clone(),
		toBeRemoved: s.toBeRemoved.Clone(),
	}
}

// Count returns the number of elements in the set.
func (s Trackable[E]) Count() int {
	return s.set.Count()
}

// HasChanges checks if the set has any changes.
func (s Trackable[E]) HasChanges() bool {
	return s.toBeAdded.Count() > 0 || s.toBeRemoved.Count() > 0
}

// Flush clears the to-be-added and to-be-removed sets.
func (s Trackable[E]) Flush() {
	s.toBeAdded.Flush()
	s.toBeRemoved.Flush()
}
