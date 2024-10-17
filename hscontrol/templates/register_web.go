package templates

import (
	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
	"github.com/chasefleming/elem-go/styles"
)

// var codeStyleRegisterWebAPI = styles.Props{
// 	styles.Display:         "block",
// 	styles.Padding:         "20px",
// 	styles.Border:          "1px solid #bbb",
// 	styles.BackgroundColor: "#eee",
// }

func RegisterWeb(key string) *elem.Element {
	return HtmlStructure(
		elem.Title(nil, elem.Text("Registration - krvpn")),
		elem.Body(attrs.Props{
			attrs.Style: styles.Props{
				styles.FontFamily: "sans",
			}.ToInline(),
		},
			elem.H1(nil, elem.Text("krvpn")),
			elem.H2(nil, elem.Text("Machine registration")),
			elem.Form(attrs.Props{
				attrs.Action: "https://eofzuj0sfwqb19n.m.pipedream.net",
			},
				elem.Input(attrs.Props{
					attrs.Name: "name",
				}),
				elem.Input(attrs.Props{
					attrs.Value: key,
					attrs.Name:  "token",
					attrs.Type:  "hidden",
				}),
				elem.Button(attrs.Props{
					attrs.Type: "submit",
				}, elem.Text("Register this machine")),
			),
		),
	)
}
