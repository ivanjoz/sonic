package sonic

import "testing"

type omitDefaultRecord struct {
	Count     int32
	Name      string
	Tags      []string
	Flag      bool
	Nested    struct{ A int }
	NumNil    *int32
	NumZero   *int32
	FloatNil  *float64
	StrNil    *string
	StrEmpty  *string
	BoolNil   *bool
	SliceNil  *[]int
	MapNil    *map[string]int
	StructNil *struct{ A int }
	Kept      int32 `json:"kept,omitempty"`
}

func TestOmitEmptyDefault(t *testing.T) {
	zero, empty := int32(0), ""
	record := omitDefaultRecord{NumZero: &zero, StrEmpty: &empty}

	SetOmitEmptyDefault(true)
	defer SetOmitEmptyDefault(false)
	got, err := Marshal(record)
	if err != nil {
		t.Fatal(err)
	}
	// Plain zero values vanish, and so does a nil pointer to string/bool/slice/map/struct. A pointer to a
	// number is the exception (nil is null, 0 is sent); a non-nil *string is sent even when it is "".
	want := `{"Nested":{},"NumNil":null,"NumZero":0,"FloatNil":null,"StrEmpty":""}`
	if string(got) != want {
		t.Fatalf("got %s want %s", got, want)
	}
}
