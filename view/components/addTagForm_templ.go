// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func AddTagForm() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<form class=\"flex flex-row w-auto items-center justify-center\" hx-put=\"/tag\" hx-target=\"#tags\" hx-swap=\"beforeend\"><input type=\"text\" class=\"w-auto rounded-full text-xl outline-none\n			text-unfocused bg-feature px-3 py-2\" placeholder=\"Add a new tag...\" name=\"tag\"> <button class=\"mx-2\" type=\"submit\" value=\"submit\"><svg class=\"hover:cursor-pointer ease-in-out hover:transition-all hover:bg-feature rounded-full\" width=\"30px\" height=\"30px\" viewBox=\"-2.4 -2.4 28.80 28.80\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\" stroke=\"#a7a7a7\"><g id=\"SVGRepo_bgCarrier\" stroke-width=\"0\"></g> <g id=\"SVGRepo_tracerCarrier\" stroke-linecap=\"round\" stroke-linejoin=\"round\"></g> <g id=\"SVGRepo_iconCarrier\"><path d=\"M4 12H20M12 4V20\" stroke=\"#a7a7a7\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\"></path></g></svg></button></form>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
