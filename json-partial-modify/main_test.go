package main

import "testing"

func BenchmarkByStd(b *testing.B) {
	var result []byte
	for i := 0; i < b.N; i++ {
		result, _ = byStd()
	}
	_ = result
}

func BenchmarkByMapStructure(b *testing.B) {
	var result []byte
	for i := 0; i < b.N; i++ {
		result, _ = byMapStructure()
	}
	_ = result
}

func BenchmarkByMapStructurePartial(b *testing.B) {
	var result []byte
	for i := 0; i < b.N; i++ {
		result, _ = byMapStructurePartial()
	}
	_ = result
}

func BenchmarkBySjson(b *testing.B) {
	var result []byte
	for i := 0; i < b.N; i++ {
		result, _ = bySjson()
	}
	_ = result
}
