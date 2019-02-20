package shellsort

import "testing"

func TestShellSort4(t *testing.T) {
	values := []int{5, 4, 3, 2, 1, 43, 65, 54, 67, 7, 4, 45, 5, 67, 34, 42, 5, 6, 7, 87}
	ShellSort(values)
	t.Log(values)
}
func TestShellSort1(t *testing.T) {
	values := []int{5, 4, 3, 2, 1}
	ShellSort(values)
	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 4 || values[4] != 5 {
		t.Error("ShellSort() failed, Got", values, "Expected 1 2 3 4 5")
	}
}
func TestShellSort2(t *testing.T) {
	values := []int{5, 5, 5, 2, 1}
	ShellSort(values)
	if values[0] != 1 || values[1] != 2 || values[2] != 5 || values[3] != 5 || values[4] != 5 {
		t.Error("ShellSort() failed, Got", values, "Expected 1 2 5 5 5")
	}
}
func TestShellSort3(t *testing.T) {
	values := []int{5}
	ShellSort(values)
	if values[0] != 5 {
		t.Error("ShellSort() failed, Got", values, "Expected 5")
	}
}
