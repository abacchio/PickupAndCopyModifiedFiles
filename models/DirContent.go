package models

// DirContent is information of Current Directory
type DirContent struct {
	Root     string     `json:"Root"`
	LogDate  string     `json:"LogDate"`
	Contents []FileInfo `json:"Contents"`
}

// FileInfo is information of file
type FileInfo struct {
	Name    string `json:"Name"`
	ModTime string `json:"ModifiedTime"`
}
