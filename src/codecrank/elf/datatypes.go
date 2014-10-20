package elf

type Segment struct {
	Offset   string
	Sections []Section
	Contents []byte
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
	NumSHEntries       uint16
	SectionNameSHIndex uint16
}

const (
	PHNull = iota
	PHLoad
	PHDynamic
	PHInterp
	PHNote
	PHShLib
	PHPHeader
	PHLOPROC = 0x70000000
	PHHIPROC = 0x7ffffff
)

const (
	NullSection = iota
	ProgBitsSection
	SymTabSection
	StrTabSection
	RelaSection
	HashSection
	DynamicSection
	NoteSection
	NoBitsSection
	RELSection
	SHLibSection
	DynSymSection
	LOPROCSection = 0x70000000
	HIPROCSection = 0x7fffffff
	LOUSERSection = 0x80000000
	HIUSERSection = 0xffffffff
)

const (
	Write = 0x1 << iota
	Alloc
	ExecInstr
	MaskProc = 0xf0000000
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

type SectionHeader struct {
	NameIndex uint32
	Type      uint32
	Flags     uint32
	Addr      uint32
	Offset    uint32
	Size      uint32
	Link      uint32
	Info      uint32
	AddrAlign uint32
	EntSize   uint32
}

type Elf struct {
	Filepath       string
	ElfHeader      ElfHeader
	ProgramHeaders []ProgramHeader
	SectionHeaders []SectionHeader
	Segments       []Segment
	Sections       []Section
	Contents       []byte
}
