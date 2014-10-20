package elf

import (
	"bytes"
	"encoding/binary"
	"io"
	"io/ioutil"
	"os"
)

func (e *Elf) readHeaders() (*Elf, error) {
	var eh ElfHeader
	var r io.ReadSeeker = bytes.NewReader(e.Contents)

	if err := binary.Read(r, binary.LittleEndian, &eh); err != nil {
		return nil, err
	}

	r.Seek(int64(eh.PHOffset), 0)

	phs := make([]ProgramHeader, eh.NumPHEntries)
	for i := 0; i < int(eh.NumPHEntries); i++ {
		var ph ProgramHeader
		if err := binary.Read(r, binary.LittleEndian, &ph); err != nil {
			return nil, err
		}

		phs[i] = ph
	}

	r.Seek(int64(eh.SHOffset), 0)

	shs := make([]SectionHeader, eh.NumSHEntries)
	for i := 0; i < int(eh.NumSHEntries); i++ {
		var sh SectionHeader
		if err := binary.Read(r, binary.LittleEndian, &sh); err != nil {
			return nil, err
		}

		shs[i] = sh
	}

	e.ElfHeader = eh
	e.ProgramHeaders = phs
	e.SectionHeaders = shs

	return e, nil
}

func (e *Elf) createSegmentsAndSections() (*Elf, error) {
	sections := make([]Section, len(e.SectionHeaders))
	for i, sh := range e.SectionHeaders {
		name, err := e.getStringAt(sh.NameIndex)
		if err != nil {
			return nil, err
		}
		sections[i] = Section{Name: name}
	}

	e.Sections = sections

	return e, nil
}

func Parse(filePath string) (elf *Elf, err error) {
	r, err := os.Open(filePath)
	defer r.Close()

	if err != nil {
		return nil, err
	}

	e := new(Elf)
	contents, err := ioutil.ReadAll(r)

	if err != nil {
		return nil, err
	}

	e.Contents = contents
	_, err = e.readHeaders()

	if err != nil {
		return nil, err
	}

	e.createSegmentsAndSections()

	if err != nil {
		return nil, err
	}

	return e, nil

}
