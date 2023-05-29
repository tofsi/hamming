package hamming

import (
	"testing"
)

func TestNew(t *testing.T){
	c := New(3)
	h := [][]bool{{true, true, false, true, true, false, false},
				  {true, false, true, true, false, true, false},
				  {false, true, true, true, false, false, true}}
	g := [][]bool{{true, false, false, false, true, true, false},
				  {false, true, false, false, true, false, true},
				  {false, false, true, false, false, true, true},
				  {false, false, false, true, true, true, true}}

	if c.m != 3{
		t.Errorf("got c.m = %d. want 3 (New(3))", c.m)
	}
	if c.k != 4{
		t.Errorf("got c.k = %d. want 4 (New(3))", c.k)
	}
	if c.n != 7{
		t.Errorf("got c.n = %d. want 7 (New(3))", c.n)
	}
	for i := 0; i < 3; i++{
		for j := 0; j < 7; j++{
			if c.h[i][j] != h[i][j]{
				t.Errorf("got c.h[%d][%d] = %t. want %t (New(3))", i, j, c.h[i][j], h[i][j])
			}
		}
	}
	for i := 0; i < 4; i++{
		for j := 0; j < 7; j++{
			if c.g[i][j] != g[i][j]{
				t.Errorf("got c.g[%d][%d] = %t. want %t (New(3))", i, j, c.g[i][j], g[i][j])
			}
		}
	}
}

func TestEncodeBin(t *testing.T){
	c := New(3)
	w := []bool{true, false, true, true, false, true}
	exp := []bool{true, false, true, true, false, true, false,
				  false, true, false, false, true, false, true}
	res := c.EncodeBin(w)
	for i := 0; i < 14; i++{
		if res[i] != exp[i]{
			t.Errorf("EncodeBin(101101)[%d] = %t, expect %t", i, res[i], exp[i])
		}
	}
}

func TestDecodeBin(t *testing.T) {
	c := New(3)
	cw := []bool{true, false, true, true, false, true, false,
				  false, true, false, false, true, false, true}
	exp := []bool{true, false, true, true, false, true}
	var res struct{
		word []bool
		bitFlips []int
	}
	// Without any bit flips
	res.word, res.bitFlips = c.DecodeBin(cw, 2)
	for i := 0; i < 6; i++{
		if res.word[i] != exp[i]{
			t.Errorf("DecodeBin(10110100100101, 2)[0][%d] = %t, expect %t", i, res.word[i], exp[i])
		}
	}
	if len(res.bitFlips) != 0{
		t.Errorf("DecodeBin(10110100100101, 2)[1] not empty, expect []")
	}
	// With bit flips
	cw[3] = !cw[3]
	res.word, res.bitFlips = c.DecodeBin(cw, 2)
	for i := 0; i < 6; i++{
		if res.word[i] != exp[i]{
			t.Errorf("DecodeBin(10100100100101, 2)[0][%d] = %t, expect %t", i, res.word[i], exp[i])
		}
	}
	if res.bitFlips[0] != 0{
		t.Errorf("DecodeBin(10100100100101, 2)[1][0] = %d, expect 0", res.bitFlips[0])
	}
}

func TestEncodeString(t *testing.T) {
	c := New(4) // Note, m = 4 in this test
	s := "hello"
	binS := []bool{false,true,true,false,true,false,false,false,
				   false,true,true,false,false,true,false,true,
				   false,true,true,false,true,true,false,false,
				   false,true,true,false,true,true,false,false,
				   false,true,true,false,true,true,true,true} // ASCII codes
	exp := c.EncodeBin(binS)
	res := c.EncodeString(s)
	for i := 0; i < 60; i++{
		if exp[i] != res[i]{
			t.Errorf("EncodeString(\"%s\")[%d] = %t, expect %t", "hello", i, res[i], exp[i])
		}
	}
}

func TestDecodeString(t *testing.T) {
	c := New(3)
	s := "hello, world"
	encodedS := c.EncodeString(s)
	var res struct{
		message string
		errorPlacements []int
	}
	res.message, res.errorPlacements = c.DecodeString(encodedS, 0)
	for i, ch := range s{
		if rune(res.message[i]) != ch{
			t.Errorf("DecodeString(EncodeString(\"%s\"))[%d] = %s, expect %s", s, i, string(rune(res.message[i])), string(ch))
		}
	}
}
