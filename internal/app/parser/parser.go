package parser

import (
	"io/ioutil"
	"packages/internal/app/models"
	"sync"
)

// Parser ...
type Parser struct {
	filesMutex *sync.Mutex
	Files      map[string][]*models.File
}

// NewParser ...
func NewParser() *Parser {
	return &Parser{
		filesMutex: new(sync.Mutex),
		Files:      make(map[string][]*models.File),
	}
}

// Read ...
func (parser *Parser) Read(directory string, token string) error {
	parser.filesMutex.Lock()
	defer parser.filesMutex.Unlock()
	fileInfo, err := ioutil.ReadDir(directory)
	if err != nil {
		return err
	}
	for _, file := range fileInfo {
		f := &models.File{
			Directory: directory,
			Name:      file.Name(),
			Size:      file.Size(),
		}
		parser.Files[token] = append(parser.Files[token], f)
	}
	return nil
}
