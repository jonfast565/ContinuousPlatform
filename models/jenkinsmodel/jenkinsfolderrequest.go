package jenkinsmodel

type JenkinsFolderRequest struct {
	Name   string `json:"name"`
	Mode   string `json:"mode"`
	From   string `json:"from"`
	Submit string `json:"Submit"`
}

func NewFolderRequest(folderName string) JenkinsFolderRequest {
	return JenkinsFolderRequest{
		Name:   folderName,
		Mode:   "com.cloudbees.hudson.plugins.folder.Folder",
		From:   "",
		Submit: "OK",
	}
}
