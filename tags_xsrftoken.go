package pongo2

import (
	"github.com/astaxie/beego"
	p2 "github.com/flosch/pongo2"
)

var xsrfTemplate = p2.Must(p2.FromString(`<input type="hidden" name="_xsrf" value="{{ _xsrf }}">`))

type tagXSRFTokenNode struct{}

func (node *tagXSRFTokenNode) Execute(ctx *p2.ExecutionContext, writer p2.TemplateWriter) *p2.Error {
	if !beego.EnableXSRF {
		return nil
	}

	xsrftoken := ctx.Public["_xsrf"]
	err := xsrfTemplate.ExecuteWriter(p2.Context{"_xsrf": xsrftoken}, writer)
	if err != nil {
		return err.(*p2.Error)
	}
	return nil
}

// tagXSRFParser implements a {% xsrftoken %} tag that inserts <input type="hidden" name="_xsrf" value="{{ _xsrf }}">
// just like Django's {% csrftoken %}. Note that we follow Beego's convention by using "XSRF" and not "CSRF".
func tagXSRFParser(doc *p2.Parser, start *p2.Token, arguments *p2.Parser) (p2.INodeTag, *p2.Error) {
	return &tagXSRFTokenNode{}, nil
}

func init() {
	p2.RegisterTag("xsrftoken", tagXSRFParser)
}
