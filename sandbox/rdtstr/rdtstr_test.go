/*
 * Проверка кода из вопроса с Reddit: https://www.reddit.com/r/golang/comments/k8wtk4/strings_question/
 */
package rdtstr_test

import "testing"

type p struct {
	s string
}

func TestStringCopy(t *testing.T) {
	x := p{s: "init"}
	s := x.s
	x.s = "new"
	if s != "init" {
		t.Errorf("want init, got %s", s)
	}
	t.Logf("s = %s, x.s = %s", s, x.s)
}
