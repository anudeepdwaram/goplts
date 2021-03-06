package popcount

import "testing"

const input29 = 275032803564053945

func TestPopCount1(t *testing.T) {
	expected := 29
	res := popCount1(input29)
	if res != expected {
		t.Errorf("expected %d got %d", expected, res)
	}
}

func TestPopCount2(t *testing.T) {
	expected := 29
	res := popCount2(input29)
	if res != expected {
		t.Errorf("expected %d got %d", expected, res)
	}
}

func TestPopCount3(t *testing.T) {
	expected := 29
	res := popCount3(input29)
	if res != expected {
		t.Errorf("expected %d got %d", expected, res)
	}
}

func TestPopCount4(t *testing.T) {
	expected := 29
	res := popCount4(input29)
	if res != expected {
		t.Errorf("expected %d got %d", expected, res)
	}
}

func BenchmarkPopCount1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popCount1(input29)
	}
}

func BenchmarkPopCount2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popCount2(input29)
	}
}

func BenchmarkPopCount3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popCount3(input29)
	}
}

func BenchmarkPopCount4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popCount4(input29)
	}
}
