package filesystem

type FileSystem interface {
	Exists(path string) (bool, error)
	GetContent(path string) ([]byte, error)
	Save(path string, content []byte) error
}
