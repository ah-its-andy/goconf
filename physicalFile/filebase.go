package physicalfile

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"

	"github.com/ah-its-andy/goconf"
)

type ResolvePhysicalFile func(*os.File) (map[string]*goconf.ExtractedValue, error)

var _ goconf.Source = (*PhysicalFileSource)(nil)

type PhysicalFileSource struct {
	fileInfo os.FileInfo

	// 文件路径
	FilePath string
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

func (source *PhysicalFileSource) BuildProvider(builder goconf.Builder) goconf.Provider {
	return &PhysicalFileProvider{
		ConfigurationProvider: *goconf.NewConfigurationProvider(),
		Source:                source,
		LoadFile:              loadYamlFile,
	}
}

func (source *PhysicalFileSource) ResolveFile() error {
	if source.fileInfo == nil &&
		len(source.FilePath) > 0 {
		fileInfo, err := os.Stat(source.FilePath)
		if err != nil {
			if err == os.ErrNotExist {
				return fmt.Errorf("file not exists: %s", source.FilePath)
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

	LoadFile func(*os.File) (map[string]*goconf.ExtractedValue, error)
}

func NewPhysicalFileProvider(source *PhysicalFileSource, loadFile func(*os.File) (map[string]*goconf.ExtractedValue, error)) *PhysicalFileProvider {
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
		provider.Source.fileInfo == nil {
		return fmt.Errorf("source not initialized")
	}
	file, err := os.OpenFile(provider.Source.FilePath, os.O_RDONLY, 0755)
	if err != nil {
		return err
	}
	defer file.Close()
	data, err := provider.LoadFile(file)
	if err != nil {
		return err
	}
	provider.ConfigurationProvider.Data = data
	return nil
}
