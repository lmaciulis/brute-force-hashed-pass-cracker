package char

import "testing"

func TestNewHolder(t *testing.T) {
	c := []rune("f")[0]
	l := 5

	h := NewHolder(l, c)

	if h.charLen != 5 {
		t.Errorf("holder length incorrect, got: %d, want: %d.", h.charLen, 5)
	}

	for _, v := range h.charList {
		if v != c {
			t.Errorf("holder char incorrect, got: %d, want: %d.", v, c)
		}
	}
}

func TestCloneHolder(t *testing.T) {
	rC := []rune("c")[0]
	rD := []rune("d")[0]
	idxD := 2
	l := 3

	h := NewHolder(l, rC)
	h.Set(idxD, rD)

	hCloned := CloneHolder(h)

	if hCloned.charLen != l {
		t.Errorf("holder length incorrect, got: %d, want: %d.", hCloned.charLen, l)
	}
	if hCloned.charList[0] != rC {
		t.Errorf("holder char incorrect at index %d, got: %d, want: %d.", 0, hCloned.charList[0], rC)
	}
	if hCloned.charList[1] != rC {
		t.Errorf("holder char incorrect at index %d, got: %d, want: %d.", 1, hCloned.charList[1], rC)
	}
	if hCloned.charList[2] != rD {
		t.Errorf("holder char incorrect at index %d, got: %d, want: %d.", 2, hCloned.charList[2], rD)
	}
}

func TestCloneHolderWithPrefix(t *testing.T) {
	rC := []rune("c")[0]
	prefix := []rune("ačiū")
	l := 3

	h := NewHolder(l, rC)
	hCloned := CloneHolderWithPrefix(h, prefix)

	if hCloned.charLen != l+len(prefix) {
		t.Errorf("holder length incorrect, got: %d, want: %d.", hCloned.charLen, l+len(prefix))
	}
	if hCloned.ToString() != "ačiūccc" {
		t.Errorf("holder value incorrect, got: %s, want: %s.", hCloned.ToString(), "ačiūccc")
	}
}

func TestCloneHolderWithSuffix(t *testing.T) {
	rC := []rune("c")[0]
	suffix := []rune("ačiū")
	l := 3

	h := NewHolder(l, rC)
	hCloned := CloneHolderWithSuffix(h, suffix)

	if hCloned.charLen != l+len(suffix) {
		t.Errorf("holder length incorrect, got: %d, want: %d.", hCloned.charLen, l+len(suffix))
	}
	if hCloned.ToString() != "cccačiū" {
		t.Errorf("holder value incorrect, got: %s, want: %s.", hCloned.ToString(), "cccačiū")
	}
}

func TestNewHolderFromCharList(t *testing.T) {
	chars := []rune("aBcD")
	h := newHolderFromCharList(chars)

	for i, v := range h.charList {
		if v != chars[i] {
			t.Errorf("holder char incorrect at index %d, got: %d, want: %d.", i, v, chars[i])
		}
	}
}

func TestHolder_Set(t *testing.T) {
	h := newHolderFromCharList([]rune("abc"))
	h.Set(2, []rune("f")[0])

	if h.ToString() != "abf" {
		t.Errorf("holder value incorrect, got: %s, want: %s.", h.ToString(), "abf")
	}
}

func TestHolder_ToString(t *testing.T) {
	h := newHolderFromCharList([]rune("abc"))

	if h.ToString() != "abc" {
		t.Errorf("holder value incorrect, got: %s, want: %s.", h.ToString(), "abc")
	}
}

func TestHolder_ToBytes(t *testing.T) {
	h := newHolderFromCharList([]rune("01č"))
	exp := []byte{48, 49, 196, 141}
	res := h.ToBytes()

	for i, v := range res {
		if v != exp[i] {
			t.Errorf("holder byte incorrect at index %d, got: %d, want: %d.", i, exp[i], v)
		}
	}
}
