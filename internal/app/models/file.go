package models

// File ...
type File struct {
	Directory string
	Name      string
	Size      int64
	Data      []byte
	Icon      []byte
}

// MainImages ...
type MainImages struct {
	FileIcon   []byte
	FolderIcon []byte
}

// HelpImages ...
type HelpImages struct {
	BackLinkIcon []byte
	DownloadIcon []byte
}
