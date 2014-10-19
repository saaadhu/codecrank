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

	return &Elf{ElfHeader: eh, ProgramHeaders: phs}, nil

}
