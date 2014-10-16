package elf

import (
	"testing"
)

func TestELFParse(t *testing.T) {
	e, err := Parse("test.elf")
	if err != nil {
		t.Error(err)
	}

	header := e.ElfHeader

	if header.Class != 1 {
		t.Error("Header does not say 32 bit")
	}
	if header.Endianness != 1 {
		t.Error("Header does not say LittleEndian")
	}
	if header.Version != 1 {
		t.Error("Header does not say Version=1")
	}
	if header.Machine != 83 {
		t.Error("Header does not say Machine=EM_AVR")
	}

	if header.Flags != 0xB3 {
		t.Error("Header does not say Flags is 0xB3")
	}
}
