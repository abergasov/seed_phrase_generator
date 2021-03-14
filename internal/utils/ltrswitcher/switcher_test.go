package ltrswitcher

import "testing"

type testStr struct {
	expArr []string
	srcArr []string
	exp    string
}

func TestLetterSwitcher_EncodeString(t *testing.T) {
	s := NewSwitcher()
	table := []testStr{
		{
			srcArr: []string{"a", "b", "c"}, // "a b c", YSBiIGM= base 64
			exp:    "YS Bi IG M",
		},
		{
			srcArr: []string{"hello", "world", "from", "golang"}, // "hello world from golang", aGVsbG8gd29ybGQgZnJvbSBnb2xhbmc= base 64
			exp:    "aGVsbG8 gd29ybG QgZnJvb SBnb2xh bmc",        // 31 / 4 = 7 letters per word + rest
		},
		{
			srcArr: []string{"Lorem", "ipsum", "dolor", "sit", "amet"}, // "Lorem ipsum dolor sit amet", TG9yZW0gaXBzdW0gZG9sb3Igc2l0IGFtZXQ= base 64
			exp:    "TG9yZW0 gaXBzdW 0gZG9sb 3Igc2l0 IGFtZXQ",          // 35 / 5 = 7 letters per word, no rest
		},
	}
	for _, v := range table {
		res := s.EncodeString(v.srcArr)
		if v.exp != res {
			t.Fatalf("result mismatch. expect %s, got %s", v.exp, res)
		}
	}
}

func TestLetterSwitcher_RotateWords(t *testing.T) {
	s := NewSwitcher()
	table := []testStr{
		{
			srcArr: []string{"a", "b", "c"}, // "a b c" -> a | b | c -> c b a
			expArr: []string{"c", "b", "a"},
		},
		{
			srcArr: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l"}, // "a,d,g,j | b,e,h,k | c,f,i,l" -> l,i,f,c + k,h,e,b + j,g,d,a
			expArr: []string{"l", "i", "f", "c", "k", "h", "e", "b", "j", "g", "d", "a"},
		},
		{
			srcArr: []string{"hello", "world", "from", "golang"}, // -> hello, golang | world | from -> from world golang hello
			expArr: []string{"from", "world", "golang", "hello"},
		},
	}
	for _, v := range table {
		res := s.RotateWords(v.srcArr)
		if len(v.expArr) != len(res) {
			t.Fatalf("result mismatch. expect %s, got %s", v.exp, res)
		}
		for i := range res {
			if v.expArr[i] != res[i] {
				t.Fatalf("result mismatch. expect %s, got %s", v.exp, res)
			}
		}
	}
}

func TestLetterSwitcher_ReplaceLetters(t *testing.T) {
	s := NewSwitcher()
	table := []testStr{
		{
			srcArr: []string{"a", "Б", "д", "Ы", "s", "-"},
			expArr: []string{"a", "B", "e", "", "s", ""},
		},
	}
	for _, v := range table {
		res := s.ReplaceLetters(v.srcArr)
		if len(v.expArr) != len(res) {
			t.Fatalf("length mismatch. expect %s, got %s", v.exp, res)
		}
		for i := range res {
			if v.expArr[i] != res[i] {
				t.Fatalf("result mismatch. expect %s, got %s", v.exp, res)
			}
		}
	}
}
