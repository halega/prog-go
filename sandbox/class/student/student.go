package student

import (
	"cmp"
	"slices"
)

type Grades map[string]int

func (g Grades) Score() int {
	var sum int
	for _, v := range g {
		sum += v
	}
	return sum
}

// Subjects returns sorted list of subjects.
func (g Grades) Subjects() []string {
	subjects := make([]string, 0, len(g))
	for subj := range g {
		subjects = append(subjects, subj)
	}
	slices.Sort(subjects)
	return subjects
}

type Student struct {
	FirstName string
	LastName  string
	Grades    Grades
}

func (s Student) Name() string {
	return s.LastName + " " + s.FirstName
}

type Class []Student

func (c Class) MaxScore() int {
	if len(c) == 0 {
		return 0
	}
	return slices.MaxFunc(c, func(a, b Student) int {
		return cmp.Compare(a.Grades.Score(), b.Grades.Score())
	}).Grades.Score()
}

// Best returns the best students in the class.
func (c Class) Best() []Student {
	maxScore := c.MaxScore()
	// there're no students with grades
	if maxScore == 0 {
		return nil
	}

	var best []Student
	for _, s := range c {
		if s.Grades.Score() == maxScore {
			best = append(best, s)
		}
	}
	return best
}

func (c Class) MinScore() int {
	if len(c) == 0 {
		return 0
	}
	return slices.MinFunc(c, func(a, b Student) int {
		return cmp.Compare(a.Grades.Score(), b.Grades.Score())
	}).Grades.Score()
}

// Worst returns the worst students in the class.
func (c Class) Worst() []Student {
	minScore := c.MinScore()
	// there're no students with grades
	if minScore == 0 {
		return nil
	}

	var worst []Student
	for _, s := range c {
		if s.Grades.Score() == minScore {
			worst = append(worst, s)
		}
	}
	return worst
}
