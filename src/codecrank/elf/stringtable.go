package elf

func (e *Elf) getStringAt(index uint32) (string, error) {
	shIndex := e.ElfHeader.SectionNameSHIndex

	if len(e.Sections) != 0 {
	}

	sh := e.SectionHeaders[shIndex]
	sOffset := sh.Offset + index

	eOffset := sOffset
	for ; e.Contents[eOffset] != 0; eOffset++ {
	}

	return string(e.Contents[sOffset:eOffset]), nil
}
