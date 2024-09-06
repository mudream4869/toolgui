package tcinput

import (
	"github.com/mudream4869/toolgui/toolgui/tgcomp/tcutil"
	"github.com/mudream4869/toolgui/toolgui/tgframe"
)

var _ tgframe.Component = &fileuploadComponent{}
var fileuploadComponentName = "fileupload_component"

type fileuploadComponent struct {
	*tgframe.BaseComponent
	Label  string `json:"label"`
	Accept string `json:"accept"`
}

func newFileuploadComponent(label, accept string) *fileuploadComponent {
	return &fileuploadComponent{
		BaseComponent: &tgframe.BaseComponent{
			Name: fileuploadComponentName,
			ID:   tcutil.NormalID(fileuploadComponentName, label),
		},
		Label:  label,
		Accept: accept,
	}
}

type FileObject struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Size int    `json:"size"`

	Bytes []byte `json:"_"`
}

// Fileupload create a fileupload and return its selected file.
func Fileupload(s *tgframe.State, c *tgframe.Container, label, accept string) *FileObject {
	comp := newFileuploadComponent(label, accept)
	c.AddComponent(comp)

	var fileObj *FileObject
	err := s.GetObject(comp.ID, &fileObj)
	if err != nil {
		panic(err)
	}

	if fileObj == nil {
		return nil
	}

	fileObj.Bytes = s.GetFile(fileObj.Name)

	return fileObj
}