package sonic

import "testing"

type omitDefaultRecord struct {
	Count   int32
	Name    string
	Tags    []string
	Flag    bool
	Nested  struct{ A int }
	PtrNil  *int32
	PtrZero *int32
	Kept    int32 `json:"kept,omitempty"`
}

func TestOmitEmptyNonPointer(t *testing.T) {
	zero := int32(0)
	record := omitDefaultRecord{PtrZero: &zero}

	SetOmitEmptyNonPointer(true)
	defer SetOmitEmptyNonPointer(false)
	got, err := Marshal(record)
	if err != nil {
		t.Fatal(err)
	}
	// Non-pointer zero values vanish; nil pointer stays null; pointer-to-zero stays 0; the struct itself is never omitted (its own zero fields are).
	want := `{"Nested":{},"PtrNil":null,"PtrZero":0}`
	if string(got) != want {
		t.Fatalf("got %s want %s", got, want)
	}
}
