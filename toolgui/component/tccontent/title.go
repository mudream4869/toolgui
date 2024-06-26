package tccontent

import (
	"github.com/mudream4869/toolgui/toolgui/component/tcutil"
	"github.com/mudream4869/toolgui/toolgui/framework"
)

var _ framework.Component = &titleComponent{}
var titleComponentName = "title_component"

type titleComponent struct {
	*framework.BaseComponent
	Text string `json:"text"`
}

func newTitleComponent(text string) *titleComponent {
	return &titleComponent{
		BaseComponent: &framework.BaseComponent{
			Name: titleComponentName,
			ID:   tcutil.NormalID(titleComponentName, text),
		},
		Text: text,
	}
}

func Title(c *framework.Container, text string) {
	c.AddComponent(newTitleComponent(text))
}
