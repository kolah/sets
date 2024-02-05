package sets_test

import (
	"testing"

	"github.com/kolah/sets"
	"github.com/stretchr/testify/require"
)

func TestSet_Add(t *testing.T) {
	t.Parallel()

	set := sets.New(1, 2, 3)
	set.Add(4)
	require.True(t, set.Contains(4), "Add method did not add the element to the set")
}

func TestSet_Contains(t *testing.T) {
	t.Parallel()

	set := sets.New(1, 2, 3)
	require.True(t, set.Contains(3), "Contains method did not find the existing element in the set")
	require.False(t, set.Contains(4), "Contains method incorrectly found a non-existing element in the set")
}

func TestSet_Items(t *testing.T) {
	t.Parallel()

	set := sets.New(1, 2, 3, 4)
	expectedItems := []int{1, 2, 3, 4}
	require.ElementsMatch(t, expectedItems, set.Items(), "Items method did not return the expected items")
}

func TestSet_Intersection(t *testing.T) {
	t.Parallel()

	set1 := sets.New(1, 2, 3, 4)
	set2 := sets.New(3, 4, 5)
	intersection := set1.Intersection(set2)
	expectedIntersection := sets.New(3, 4)
	require.Equal(
		t,
		expectedIntersection,
		intersection,
		"Intersection method did not return the expected intersection set",
	)
}

func TestSet_Intersects(t *testing.T) {
	t.Parallel()

	set1 := sets.New(1, 2, 3)
	set2 := sets.New(3, 4, 5)
	require.True(t, set1.Intersects(set2), "Intersects method did not detect intersection between sets")
	require.False(
		t,
		set1.Intersects(sets.New(6, 7, 8)),
		"Intersects method incorrectly detected intersection between sets with no common elements",
	)
}

func TestSet_Remove(t *testing.T) {
	t.Parallel()

	set := sets.New(1, 2, 3, 4)
	require.True(t, set.Contains(2), "Initial set should contain element 2")

	set.Remove(2)
	require.False(t, set.Contains(2), "Remove method did not remove the element from the set")

	set.Remove(5)
	require.False(t, set.Contains(5), "Remove method incorrectly removed a non-existing element")

	expectedItems := []int{1, 3, 4}
	require.ElementsMatch(t, expectedItems, set.Items(), "Remove method altered the set incorrectly")
}

func TestSet_Clone(t *testing.T) {
	t.Parallel()

	set := sets.New(1, 2, 3, 4)
	clonedSet := set.Clone()

	require.Equal(t, set, clonedSet, "Clone method did not produce an equal cloned set")

	clonedSet.Add(5)
	require.False(t, set.Contains(5), "Modifying cloned set should not affect the original set")
}

func TestSet_Count(t *testing.T) {
	t.Parallel()

	emptySet := sets.New[int]()
	setWithValues := sets.New(1, 2, 3, 4)

	require.Equal(t, 0, emptySet.Count(), "Count method on empty set did not return 0")
	require.Equal(t, 4, setWithValues.Count(), "Count method on non-empty set did not return the correct count")
}

func TestSet_Diff(t *testing.T) {
	t.Parallel()

	set1 := sets.New(1, 2, 3, 4)
	set2 := sets.New(3, 4, 5)

	diff := set1.Diff(set2)

	expectedDiff := sets.New(1, 2)

	require.Equal(t, expectedDiff, diff, "Diff method did not return the expected diff set")
}

func TestSet_Flush(t *testing.T) {
	t.Parallel()

	set := sets.New(1, 2, 3, 4)
	set.Add(5)
	set.Remove(2)
	set.Flush()

	expectedItems := make([]int, 0)
	require.ElementsMatch(t, expectedItems, set.Items(), "Flush method did not flush the set")
}
