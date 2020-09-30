package parser

import (
	"fmt"
	"io/ioutil"
	"packages/internal/app/models"
	"sync"
)

// Parser ...
type Parser struct {
	filesMutex *sync.Mutex
	Files      map[string][]*models.File
	Icons      *models.HelpImages
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
	filesInfo, err := ioutil.ReadDir(directory)
	if err != nil {
		return err
	}
	icons, err := parser.ReadMainIcon()
	if err != nil {
		return err
	}
	for _, file := range filesInfo {
		var fileReader []byte
		var icon []byte
		if !file.IsDir() {
			fileReader, err = ioutil.ReadFile(fmt.Sprintf("%s/%s", directory, file.Name()))
			if err != nil {
				return err
			}
			icon = icons.FileIcon
		} else {
			icon = icons.FolderIcon
		}
		f := &models.File{
			Directory: directory,
			Name:      file.Name(),
			Size:      file.Size(),
			Data:      fileReader,
			Icon:      icon,
		}
		parser.Files[token] = append(parser.Files[token], f)
	}
	return nil
}

// ReadMainIcon ...
func (parser *Parser) ReadMainIcon() (*models.MainImages, error) {
	folder, err := ioutil.ReadFile("../../resources/static/folder_icon.txt")
	if err != nil {
		return nil, err
	}
	file, err := ioutil.ReadFile("../../resources/static/file_icon.txt")
	if err != nil {
		return nil, err
	}
	return &models.MainImages{
		FolderIcon: folder,
		FileIcon:   file,
	}, nil
}
