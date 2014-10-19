package elf

import (
	"testing"
)

func TestELFParse(t *testing.T) {
	e, err := Parse("test.elf")
	if err != nil {
		t.Error(err)
	}

	h := e.ElfHeader

	if h.Class != 1 {
		t.Error("h does not say 32 bit")
	}
	if h.Endianness != 1 {
		t.Error("h does not say LittleEndian")
	}
	if h.Version != 1 {
		t.Error("h does not say Version=1")
	}
	if h.Machine != 83 {
		t.Error("h does not say Machine=EM_AVR")
	}
	if h.Flags != 0xB3 {
		t.Error("h does not say Flags is 0xB3")
	}
	if h.PHOffset != 0x34 {
		t.Error("h does not say PHOffset is 0x34")
	}
	if h.NumPHEntries != 2 {
		t.Error("h does not say PHEntries is 0x2")
	}

	if h.NumSHEntries != 13 {
		t.Errorf("h does not say SHEntries is 13, %d instead", h.NumSHEntries)
	}
	phs := e.ProgramHeaders

	if len(phs) != 2 {
		t.Error("e does not have 2 PHs")
	}

	ph0 := phs[0]
	if ph0.Type != PHLoad {
		t.Error("ph0 does not have Type Load")
	}
	if ph0.Offset != 0x74 {
		t.Error("ph0 does not Offset 0x74")
	}
	if ph0.VAddr != 0x0 {
		t.Error("ph0 does not Vaddr 0x0")
	}
	if ph0.PAddr != 0x0 {
		t.Error("ph0 does not Paddr 0x0")
	}
	if ph0.FileSize != 0x112 {
		t.Error("ph0 does not FileSize 0x112")
	}
	if ph0.MemSize != 0x112 {
		t.Error("ph0 does not MemSize 0x112")
	}

	ph1 := phs[1]
	if ph1.Type != PHLoad {
		t.Error("ph1 does not have Type Load")
	}
	if ph1.Offset != 0x186 {
		t.Error("ph0 does not Offset 0x186")
	}
	if ph1.VAddr != 0x800200 {
		t.Error("ph1 does not Vaddr 0x80200")
	}
	if ph1.PAddr != 0x112 {
		t.Error("ph1 does not Paddr 0x112")
	}
	if ph1.FileSize != 0x0 {
		t.Error("ph1 does not FileSize 0x0")
	}
	if ph1.MemSize != 0x0 {
		t.Error("ph1 does not MemSize 0x0")
	}

	shs := e.SectionHeaders

	if len(shs) != 13 {
		t.Error("SectionHeaders count is not 13")
	}

	if shs[1].Offset != 0x186 {
		t.Errorf("Section1 offset is not 0x186, %#x", shs[1].Offset)
	}
	if shs[1].Addr != 0x800200 {
		t.Errorf("Section1 Addr is not 0x800200, %#x", shs[1].Addr)
	}
	if shs[1].Type != ProgBitsSection {
		t.Error("Section1 Type is not ProgBitsSection")
	}
	if shs[1].Flags != (Write | Alloc) {
		t.Error("Section1 Flags is not Write | Alloc")
	}
}
