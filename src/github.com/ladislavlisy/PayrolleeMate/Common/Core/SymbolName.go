package Core

type SymbolName struct {
	code uint32
	name string
}

func NewSymbolName(code uint32, name string) *SymbolName {
	return &SymbolName{code, name}
}

func (s SymbolName) Equals(other SymbolName) bool {
	return (s.code == other.code)
}

func (s SymbolName) CompareTo(other SymbolName) int {
	return (int(s.code) - int(other.code))
}

func (s SymbolName) Gt(other SymbolName) bool {
	return (s.CompareTo(other) > 0)
}

func (s SymbolName) Lt(other SymbolName) bool {
	return (s.CompareTo(other) < 0)
}

