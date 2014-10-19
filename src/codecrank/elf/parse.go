package elf

import (
	"encoding/binary"
	"os"
)

func Parse(filePath string) (elf *Elf, err error) {
	r, err := os.Open(filePath)
	if err != nil {
		return
	}

	var eh ElfHeader
	if err = binary.Read(r, binary.LittleEndian, &eh); err != nil {
		return
	}

	r.Seek(int64(eh.PHOffset), 0)

	phs := make([]ProgramHeader, eh.NumPHEntries)
	for i := 0; i < int(eh.NumPHEntries); i++ {
		var ph ProgramHeader
		if err = binary.Read(r, binary.LittleEndian, &ph); err != nil {
			return
		}

		phs[i] = ph
	}

	r.Seek(int64(eh.SHOffset), 0)

	shs := make([]SectionHeader, eh.NumSHEntries)
	for i := 0; i < int(eh.NumSHEntries); i++ {
		var sh SectionHeader
		if err = binary.Read(r, binary.LittleEndian, &sh); err != nil {
			return
		}

		shs[i] = sh
	}

	return &Elf{ElfHeader: eh, ProgramHeaders: phs, SectionHeaders: shs}, nil

}
