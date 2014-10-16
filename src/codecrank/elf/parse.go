package elf

import (
	"encoding/binary"
	"os"
)

func Parse(filePath string) (elf *Elf, err error) {
	reader, err := os.Open(filePath)
	if err != nil {
		return
	}

	var header ElfHeader
	if err = binary.Read(reader, binary.LittleEndian, &header); err != nil {
		return
	}

	return &Elf{ElfHeader: header}, nil

}
