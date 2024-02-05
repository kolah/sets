package sets_test

import (
	"testing"

	"github.com/kolah/sets"
	"github.com/stretchr/testify/require"
)

func TestTrackableSet_Add(t *testing.T) {
	t.Parallel()

	ts := sets.NewTrackable(1, 2, 3)
	ts.Add(4)

	require.True(t, ts.Contains(4), "Add method did not add the element to the trackable set")
	require.True(t, ts.ToBeAdded().Contains(4), "Element was not marked as to be added")
}

func TestTrackableSet_Remove(t *testing.T) {
	t.Parallel()

	ts := sets.NewTrackable(1, 2, 3)
	ts.Remove(2)

	require.False(t, ts.Contains(2), "Remove method did not remove the element from the trackable set")
	require.True(t, ts.ToBeRemoved().Contains(2), "Element was not marked as to be removed")
}

func TestTrackableSet_RemoveNonExisting(t *testing.T) {
	t.Parallel()

	ts := sets.NewTrackable(1, 2, 3)
	ts.Remove(4)

	require.False(t, ts.ToBeRemoved().Contains(4), "Remove method marked a non-existing element as to be removed")
}

func TestTrackableSet_Intersection(t *testing.T) {
	t.Parallel()

	ts1 := sets.NewTrackable(1, 2, 3, 4)
	ts2 := sets.NewTrackable(3, 4, 5)
	intersection := ts1.Intersection(ts2)

	require.True(t, intersection.Contains(3), "Intersection method did not return the expected intersection set")
	require.True(t, intersection.Contains(4), "Intersection method did not return the expected intersection set")
	require.False(t, intersection.Contains(1), "Intersection method returned an unexpected element")
	require.False(t, intersection.Contains(5), "Intersection method returned an unexpected element")
}

func TestTrackableSet_Intersects(t *testing.T) {
	t.Parallel()

	ts1 := sets.NewTrackable(1, 2, 3)
	ts2 := sets.NewTrackable(3, 4, 5)

	require.True(t, ts1.Intersects(ts2), "Intersects method did not detect intersection between trackable sets")
	require.False(
		t,
		ts1.Intersects(sets.NewTrackable(6, 7, 8)),
		"Intersects method incorrectly detected intersection between trackable sets with no common elements",
	)
}

func TestTrackableSet_Contains(t *testing.T) {
	t.Parallel()

	ts := sets.NewTrackable(1, 2, 3)

	require.True(t, ts.Contains(1), "Contains method did not find an existing element in the trackable set")
	require.False(t, ts.Contains(4), "Contains method incorrectly found a non-existing element in the trackable set")
}

func TestTrackableSet_Items(t *testing.T) {
	t.Parallel()

	ts := sets.NewTrackable(1, 2, 3, 4)
	expectedItems := []int{1, 2, 3, 4}

	require.ElementsMatch(t, expectedItems, ts.Items(), "Items method did not return the expected items")
}

func TestTrackableSet_Clone(t *testing.T) {
	t.Parallel()

	ts := sets.NewTrackable(1, 2, 3)
	clonedTS := ts.Clone()

	require.Equal(t, ts.ToBeAdded(), clonedTS.ToBeAdded(), "Cloned trackable set's to-be-added values are not equal")
	require.Equal(t, ts.ToBeRemoved(), clonedTS.ToBeRemoved(), "Cloned trackable set's to-be-removed values are not equal")
	require.ElementsMatch(t, ts.Items(), clonedTS.Items(), "Cloned trackable set's elements are not equal")

	clonedTS.Add(4)
	require.False(t, ts.Contains(4), "Modifying cloned trackable set should not affect the original trackable set")
}

func TestTrackableSet_HasChanges(t *testing.T) {
	t.Parallel()

	ts := sets.NewTrackable(1, 2, 3)
	ts.Add(4)

	require.True(t, ts.HasChanges(), "HasChanges method did not detect changes in the trackable set")
}

func TestTrackableSet_Flush(t *testing.T) {
	t.Parallel()

	ts := sets.NewTrackable(1, 2, 3)
	ts.Add(4)
	ts.Remove(2)
	ts.Flush()

	require.Empty(t, ts.ToBeAdded(), "Flush method did not clear to-be-added set")
	require.Empty(t, ts.ToBeRemoved(), "Flush method did not clear to-be-removed set")
	require.False(t, ts.HasChanges(), "Flush method did not clear changes in the trackable set")
}

func TestTrackableSet_Count(t *testing.T) {
	t.Parallel()

	// Create a trackable set with initial values
	ts := sets.NewTrackable(1, 2, 3, 4)

	// Check the count of elements in the trackable set
	require.Equal(t, 4, ts.Count(), "Count method on trackable set did not return the correct count")

	// Add an element to the trackable set and check the count again
	ts.Add(5)
	require.Equal(t, 5, ts.Count(), "Count method on trackable set did not update correctly after adding an element")

	// Remove an element from the trackable set and check the count again
	ts.Remove(2)
	require.Equal(t, 4, ts.Count(), "Count method on trackable set did not update correctly after removing an element")
}

func TestTrackableSet_Diff(t *testing.T) {
	t.Parallel()

	ts1 := sets.NewTrackable(1, 2, 3, 4)
	ts2 := sets.NewTrackable(3, 4, 5)

	diff := ts1.Diff(ts2)

	expectedDiff := sets.NewTrackable(1, 2)

	require.Equal(t, expectedDiff, diff, "Diff method did not return the expected diff set")
}
