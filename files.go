package main

type FileKind int

func (f FileKind) String() string {
	switch f {
	case Directory:
		return "/"
	}
	return ""
}

const (
	Directory FileKind = iota
	File
)

type FileType struct {
	Name string
	Path string
	Kind FileKind
}

func NewFileType(name string, path string, isDir bool) FileType {
	if isDir {
		return FileType{
			Name: name,
			Path: path,
			Kind: Directory,
		}
	}

	return FileType{
		Name: name,
		Path: path,
		Kind: File,
	}
}

// to implement bubble tea list.Item
func (f FileType) FilterValue() string {
	return f.Name
}

func (f *FileType) String() string {
	return f.Name + f.Kind.String()
}
