package physicalfile

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"

	"github.com/ah-its-andy/goconf"
)

type ResolvePhysicalFile func(io.Reader) (map[string]*goconf.ExtractedValue, error)

var _ goconf.Source = (*PhysicalFileSource)(nil)

type PhysicalFileSource struct {
	fileInfo os.FileInfo

	// 文件路径
	FilePath string
	reader   io.Reader
	fn       ResolvePhysicalFile
}

func PhysicalFile(filePath string, fn ResolvePhysicalFile) *PhysicalFileSource {
	path := filePath
	if !filepath.IsAbs(path) {
		path, _ = filepath.Abs(path)
	}
	return &PhysicalFileSource{
		FilePath: path,
		fn:       fn,
	}
}

func FromReader(reader io.Reader, fn ResolvePhysicalFile) *PhysicalFileSource {
	return &PhysicalFileSource{
		reader: reader,
		fn:     fn,
	}
}

func (source *PhysicalFileSource) BuildProvider(builder goconf.Builder) goconf.Provider {
	return &PhysicalFileProvider{
		ConfigurationProvider: *goconf.NewConfigurationProvider(),
		Source:                source,
		LoadFile:              source.fn,
	}
}

func (source *PhysicalFileSource) ResolveFile() error {
	if source.fileInfo == nil &&
		len(source.FilePath) > 0 {
		fileInfo, err := os.Stat(source.FilePath)
		if err != nil {
			if os.IsNotExist(err) {
				return os.ErrNotExist
			} else {
				return err
			}
		}
		source.fileInfo = fileInfo
	}
	return nil
}

var _ goconf.Provider = (*PhysicalFileProvider)(nil)

type PhysicalFileProvider struct {
	goconf.ConfigurationProvider

	Source *PhysicalFileSource

	LoadFile ResolvePhysicalFile
}

func NewPhysicalFileProvider(source *PhysicalFileSource, loadFile ResolvePhysicalFile) *PhysicalFileProvider {
	return &PhysicalFileProvider{
		ConfigurationProvider: *goconf.NewConfigurationProvider(),
		Source:                source,
		LoadFile:              loadFile,
	}
}

func (provider *PhysicalFileProvider) String() string {
	return fmt.Sprintf("%s for '%s'", reflect.TypeOf(provider).Elem().Name(), provider.Source.FilePath)
}

func (provider *PhysicalFileProvider) Load() error {
	if provider.LoadFile == nil {
		return fmt.Errorf("load file func not set")
	}

	err := provider.Source.ResolveFile()
	if err != nil {
		return err
	}

	if provider.Source == nil ||
		(provider.Source.fileInfo == nil && provider.Source.reader == nil) {
		return fmt.Errorf("source not initialized")
	}
	reader := provider.Source.reader
	if provider.Source.fileInfo != nil {
		file, err := os.OpenFile(provider.Source.FilePath, os.O_RDONLY, 0755)
		if err != nil {
			return err
		}
		defer file.Close()
		reader = file
	}
	data, err := provider.LoadFile(reader)
	if err != nil {
		return err
	}
	provider.ConfigurationProvider.Data = data
	return nil
}
