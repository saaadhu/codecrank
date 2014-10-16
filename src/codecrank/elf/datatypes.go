package elf

type Segment struct {
	Name   string
	Offset string
}

type Section struct {
	Name     string
	Offset   int
	Contents []byte
}

type ElfHeader struct {
	Magic              uint32
	Class              uint8
	Endianness         uint8
	Version            uint8
	OsABI              uint8
	ABIVersion         uint8
	_                  uint8
	_                  uint8
	_                  uint8
	_                  uint32
	Type               uint16
	Machine            uint16
	_                  uint32
	EntryPoint         uint32
	PHOffset           uint32
	SHOffset           uint32
	Flags              uint32
	HeaderSize         uint16
	PHEntrySize        uint16
	PHEntries          uint16
	SHEntrySize        uint16
	SHEntry            uint16
	SectionNameSHIndex uint16
}

type Elf struct {
	Filepath  string
	ElfHeader ElfHeader
	//ProgramHeader ProgramHeader
	//SectionHeader SectionHeader
}
