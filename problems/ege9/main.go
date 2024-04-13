/*
Задание 9
(А. Богданов) В файле электронной таблицы 9-227.xls в каждой строке записаны четыре натуральных числа. Определите количество строк таблицы, для которых выполнены следующие условия:
– только одно число встречается в строке дважды;
– сумма двух самых больших чисел строки более чем в два раза больше суммы двух самых малых;
– максимальное число строки не кратно минимальному.
В ответе запишите только число.
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	buf := bufio.NewScanner(f)
	count := 0
	for buf.Scan() {
		nums, err := parse(buf.Text())
		if err != nil {
			log.Fatal(err)
		}
		if valid(nums) {
			count++
		}
	}
	if buf.Err() != nil {
		log.Fatal(buf.Err())
	}
	fmt.Println(count)
}

func parse(line string) ([4]int, error) {
	var nums [4]int
	numStrings := strings.Split(line, ",")
	if len(numStrings) != 4 {
		return nums, fmt.Errorf("unexpected number count in line (expect 4): %v", numStrings)
	}
	for i, s := range numStrings {
		var err error
		nums[i], err = strconv.Atoi(strings.TrimSpace(s))
		if err != nil {
			return nums, fmt.Errorf("parse numbers in line %v: %w", s, err)
		}
	}
	return nums, nil
}

func valid(nums [4]int) bool {
	slices.Sort(nums[:])
	// только одно число встречается дважды:
	// после удаления всех дублей в слайсе должно остаться ровно три числа
	distinctNums := slices.Compact(slices.Clone(nums[:]))
	if len(distinctNums) != 3 {
		return false
	}

	// сумма двух самых больших чисел строки более чем в два раза больше суммы двух самых малых
	if (nums[0]+nums[1])*2 > nums[2]+nums[3] {
		return false
	}

	// максимальное число строки не кратно минимальному
	if nums[3]%nums[0] == 0 {
		return false
	}

	return true
}
