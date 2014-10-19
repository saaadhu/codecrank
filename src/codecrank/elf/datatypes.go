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
	NumPHEntries       uint16
	SHEntrySize        uint16
	SHEntry            uint16
	SectionNameSHIndex uint16
}

const (
	Null    = iota
	Load    = iota
	Dynamic = iota
	Interp  = iota
	Note    = iota
	ShLib   = iota
	PHeader = iota
	LOPROC  = 0x70000000
	HIPROC  = 0x7ffffff
)

type ProgramHeader struct {
	Type     uint32
	Offset   uint32
	VAddr    uint32
	PAddr    uint32
	FileSize uint32
	MemSize  uint32
	Flags    uint32
	Align    uint32
}

type Elf struct {
	Filepath       string
	ElfHeader      ElfHeader
	ProgramHeaders []ProgramHeader
	//SectionHeader SectionHeader
}
