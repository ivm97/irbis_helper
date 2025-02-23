package files

type Files struct {
	//serverPath    string
	clientPath    string
	clientVersion string
}

func NewFiles(clientPath string, clientVersion string) *Files {
	return &Files{clientPath: clientPath, clientVersion: clientVersion}
}

func (f *Files) ClientPath() string {
	return f.clientPath
}
func (f *Files) ClientVersion() string {
	return f.clientVersion
}
