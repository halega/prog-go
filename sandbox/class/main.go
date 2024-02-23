package main

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
	"unicode/utf8"

	"github.com/halega/prog-go/sandbox/class/student"
)

type formattedGrades student.Grades

func (g formattedGrades) String(subjects []string) string {
	var line string
	for _, subj := range subjects {
		if gr, ok := g[subj]; ok {
			line += fmt.Sprintf("%*d ", len(subj), gr)
		} else {
			line += fmt.Sprintf("%*s ", len(subj), "")
		}
	}
	return line
}

type formattedClass student.Class

func (c formattedClass) String() string {
	line := c.Header() + "\n"
	w := c.width()
	subjects := c.Subjects()
	for _, s := range c {
		line += fmt.Sprintf("%-*s %v\n", w, s.Name(), formattedGrades(s.Grades).String(subjects))
	}
	return line
}

func (c formattedClass) Header() string {
	w := c.width()
	header := fmt.Sprintf("%-*s ", w, "Name")
	header += strings.Join(c.Subjects(), " ")
	return header
}

func (c formattedClass) Subjects() []string {
	subjects := make(map[string]bool)
	for _, st := range c {
		for subj := range st.Grades {
			subjects[subj] = true
		}
	}

	res := make([]string, 0, len(subjects))
	for s := range subjects {
		res = append(res, s)
	}
	slices.Sort(res)
	return res
}

// width returns width of the Name column.
func (c formattedClass) width() int {
	if len(c) == 0 {
		return len("Name")
	}

	longestName := slices.MaxFunc(c, func(a, b student.Student) int {
		return cmp.Compare(utf8.RuneCountInString(a.Name()), utf8.RuneCountInString(b.Name()))
	}).Name()

	return utf8.RuneCountInString(longestName)
}

func sample() student.Class {
	return student.Class{
		{FirstName: "Veronika", LastName: "Dobrodeeva",
			Grades: student.Grades{"math": 5, "phys": 4, "chem": 3, "rus": 5, "eng": 4}},
		{FirstName: "Angelina", LastName: "Strikachova",
			Grades: student.Grades{"math": 5, "phys": 5, "eng": 5, "geo": 5, "bio": 5}},
		{FirstName: "Ivan", LastName: "Antonov",
			Grades: student.Grades{"math": 4, "phys": 4, "rus": 4, "eng": 4, "geo": 3}},
	}
}

func main() {
	c := sample()
	fmt.Println(formattedClass(c))
	fmt.Println()
	fmt.Println("Best student(s): ", c.Best())
	fmt.Println("Worst student(s): ", c.Worst())
}
