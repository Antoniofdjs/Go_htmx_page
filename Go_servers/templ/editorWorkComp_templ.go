// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"Go_servers/models"
	"fmt"
)

/* Base layout for the editor view */

func EditorSection(works []models.WorkFrontEnd) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<section class=\"px-0 dark:bg-gray-800 dark:text-gray-900 justify-between w-screen\" id=\"works-section\"><!-- images outer container --><div class=\"mx-auto px-0 flex flex-col items-center justify-between w-screen overflow-hidden\"><div class=\"-m-1 flex flex-wrap md:-m-2 flex-col items-center w-screen md:gap-1 gap-3\" id=\"gallery\"><!-- IMAGES HERE BELOW: -->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, work := range works {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex w-full h-[500px] flex-wrap pb-10 object-center relative\" id=\"work-container-{{ .Position }}\"><img alt=\"gallery\" class=\"block w-full h-full object-cover object-center absolute top-0 left-0\" src=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var2 string
			templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(work.Path)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `editorWorkComp.templ`, Line: 20, Col: 31}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><!-- Image text container --><div class=\"w-full h-full hover:opacity-100 absolute top-0 left-0 flex items-center\"><div class=\"bg-black bg-opacity-50 w-full h-full flex flex-col items-baseline\"><h2 class=\"text-white md:pl-2 w-full text-left mt-2 font-serif text-xl\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(work.Title)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `editorWorkComp.templ`, Line: 25, Col: 100}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h2>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = ButtonsContainer(work).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div></section>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

/* Container for all buttons visible on the editor view.
These include 'Edit Title', 'Change Picture', 'Delete', 'Insert Work Above', 'Insert Work Below'.
To call these buttons use @EditorButton()
}
*/

func ButtonsContainer(work models.WorkFrontEnd) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var4 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var4 == nil {
			templ_7745c5c3_Var4 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"w-1/2 md:w-2/12 flex flex-col gap-y-2 ml-3 bg-gradient-to-r from-slate-400 to-slate-200 rounded-lg overflow-hidden shadow-gray-700 shadow-md\" id=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 string
		templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(templ.JoinStringErrs("buttons-container-" + work.Position))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `editorWorkComp.templ`, Line: 44, Col: 71}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = EditorButton("Edit", work).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = EditorButton("InsertAbove", work).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = EditorButton("InsertBelow", work).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = EditorButton("Delete", work).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

/*
Individual buttons of the container for the editor mode. Pass into it the component name and work data.
Example if you need the delete button: @EditorButton("Delete", work)
The button will have the correct redirects and hx-vals and methods to call in backend.
*/

func EditorButton(component string, work models.WorkFrontEnd) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var6 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var6 == nil {
			templ_7745c5c3_Var6 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<button type=\"button\" class=\"h-14 bg-slate-200 hover:bg-slate-400 text-gray-700 font-bold py-2 px-4\" hx-get=\"/editor/components\" hx-target=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var7 string
		templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(templ.JoinStringErrs("#buttons-container-" + work.Position))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `editorWorkComp.templ`, Line: 64, Col: 79}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-swap=\"outerHTML\" hx-vals=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var8 string
		templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf(`{"Position": "%s", "Component": "%s", "Title": "%s", "Description": "%s"}`, work.Position, component, work.Title, work.Description))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `editorWorkComp.templ`, Line: 66, Col: 162}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var9 string
		templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(TextForButton(component))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `editorWorkComp.templ`, Line: 68, Col: 34}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</button>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func Delete(work models.WorkFrontEnd) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var10 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var10 == nil {
			templ_7745c5c3_Var10 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex flex-col h-1/3 w-11/12 mx-4 md:w-1/2 gap-2 p-6 rounded-lg shadow-md bg-white dark:text-gray-800 md:items-start text-left justify-center\" id=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var11 string
		templ_7745c5c3_Var11, templ_7745c5c3_Err = templ.JoinStringErrs(templ.JoinStringErrs("delete-work-component-" + work.Position))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `editorWorkComp.templ`, Line: 75, Col: 69}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var11))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><h2 class=\"flex items-center gap-2 text-xl font-semibold leading-tight tracking-wide\"><!-- The button is for closing this component and retreiving the previous one --><button type=\"button\" class=\"flex w-6 h-6 p-1 bg-transparent border-none outline-none cursor-pointer hover:bg-gray-200 justify-center items-center text-center\" aria-label=\"Close\" hx-headers=\"{&#34;Content-Type&#34;: &#34;application/json&#34;}\" hx-get=\"/editor/components\" hx-vals=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var12 string
		templ_7745c5c3_Var12, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf(`{"Position": "%s", "Component": "ButtonsEditor", "Title": "%s", "Description": "%s"}`, work.Position, work.Title, work.Description))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `editorWorkComp.templ`, Line: 85, Col: 160}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var12))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-target=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var13 string
		templ_7745c5c3_Var13, templ_7745c5c3_Err = templ.JoinStringErrs(templ.JoinStringErrs("#delete-work-component-" + work.Position))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `editorWorkComp.templ`, Line: 86, Col: 81}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var13))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-swap=\"outerHTML\"><!-- Add button text or icon here if needed --><span class=\"text-red-500 hover:text-red-600\">X</span><!-- You can style this text as needed --></button> DELETE WORK: ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var14 string
		templ_7745c5c3_Var14, templ_7745c5c3_Err = templ.JoinStringErrs(work.Title)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `editorWorkComp.templ`, Line: 93, Col: 29}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var14))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h2><div class=\"flex flex-col justify-between t md:gap-3 mt-6 sm:flex-row items-center w-full text-left\"><p class=\"px-6 py-2 rounded-md md:w-1/2 w-full text-center\">---WARNING---</p><button id=\"delete-button\" type=\"submit\" class=\"py-1 rounded-md md:w-1/6 w-1/4 bg-red-500 text-center hover:bg-red-800 border border-red-900\" hx-post=\"/editor/del\" hx-target=\"#works-section\" hx-swap=\"outerHTML\" hx-vals=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var15 string
		templ_7745c5c3_Var15, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf(`{"Position": "%s"}`, work.Position))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `editorWorkComp.templ`, Line: 104, Col: 64}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var15))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">DELETE</button></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func DeleteView() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var16 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var16 == nil {
			templ_7745c5c3_Var16 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex md:flex-col flex-col justify-start gap-3 md:gap-2 items-center h-full w-full rounded-md md:h-full\"><div class=\"flex flex-col gap-3 md:gap-2 rounded-md md:w-full w-full h-3/4 md:h-full align-text-top items-center md:items-end justify-end text-black flex-1\"><div class=\" flex items-center md:items-end justify-center w-1/2 md:w-1/3 h-1/3 md:h-1/3 \"><button type=\"submit\" class=\"md:font-semibold bg-red-500 hover:bg-red-600 dark:text-gray-100 w-full h-full md:h-3/5  rounded-xl shadow-md shadow-slate-500\">Delete</button></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func FileUploadView(operation string) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var17 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var17 == nil {
			templ_7745c5c3_Var17 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex md:flex-col flex-col justify-start gap-3 md:gap-2 items-center bg-slate-200 h-full w-full rounded-md md:h-full\"><div class=\"flex justify-center items-center w-full h-1/4 md:h-1/5 sm:text-sm md:text-lg rounded-lg  bg-slate-400 shadow-sm shadow-slate-500 border border-slate-500\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if operation == "Edit" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<p class=\"flex justify-center items-center h-full w-full text-center text-black font-medium\">Change cover image</p>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if operation == "InsertAbove" || operation == "InsertBelow" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<p class=\"flex justify-center items-center h-full w-full text-center text-black font-medium\">Choose an image  file below:</p>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div class=\"flex flex-col gap-3 md:gap-2 rounded-md md:w-full w-full h-3/4 md:h-2/3 align-text-top items-end justify-end text-black flex-1\"><input class=\"text-center items-center justify-center w-full h-4/5 md:h-2/3 flex text-black rounded-lg bg-slate-200 pt-2 pl-2 shadow-sm shadow-slate-500 cursor-pointer hover:bg-slate-400 border border-blue-100\" type=\"file\" id=\"picture\" name=\"picture\" accept=\"image/*\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if operation == "Insert" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" required=\"true\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("><div class=\" flex items-end justify-center w-1/2 md:w-1/3 h-1/2 md:h-1/3 \"><button type=\"submit\" class=\"md:font-semibold bg-emerald-400 hover:bg-emerald-500 dark:text-gray-100 w-full h-full md:h-3/4  rounded-xl shadow-md shadow-slate-500\">Submit</button></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

// Decide method to use base on the operation
func GetHxMethod(operation string) string {
	var method string
	if operation == "EditTitle" {
		method = `hx-put="/editor"`
	}
	if operation == "InsertAbove" {
		method = `hx-post="/editor"`
	}
	return method
}

// Generate <input> fields for edit, insert, or a <p> for delete. Used inside the ButtonView() component in the -Title input- area.
func GetInputsOrMessage(operation string, title string) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var18 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var18 == nil {
			templ_7745c5c3_Var18 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		if operation != "Delete" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input class=\"w-full h-full sm:text-sm md:text-lg rounded-lg dark:text-gray-800 bg-slate-200 text-center hover:bg-slate-400 hover:shadow-md\" type=\"text\" name=\"Title\" id=\"title\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if operation == "Edit" {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" value=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var19 string
				templ_7745c5c3_Var19, templ_7745c5c3_Err = templ.JoinStringErrs(title)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `editorWorkComp.templ`, Line: 191, Col: 23}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var19))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			if operation == "InsertAbove" || operation == "InsertBelow" {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" placeholder=\"Title\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" required=\"true\" name=\"Title\"> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if operation == "Delete" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<p class=\" flex justify-center items-center bg-red-200 h-full w-full rounded-md md:h-full text-center border border-red-500\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var20 string
			templ_7745c5c3_Var20, templ_7745c5c3_Err = templ.JoinStringErrs(title)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `editorWorkComp.templ`, Line: 202, Col: 15}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var20))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		return templ_7745c5c3_Err
	})
}

// Generate content for message displayed when in a specific operation: delete,insert,edit. Used inside the ButtonView() in the input title div.
func TextAreaContent(operation string, description string) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var21 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var21 == nil {
			templ_7745c5c3_Var21 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		if operation == "Edit" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<textarea class=\"rounded-md md:w-full w-full h-full align-text-top items-start justify-start bg-slate-200 text-black resize-none pl-1 hover:shadow-md hover:bg-slate-400\" name=\"Description\" required=\"true\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var22 string
			templ_7745c5c3_Var22, templ_7745c5c3_Err = templ.JoinStringErrs(description)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `editorWorkComp.templ`, Line: 214, Col: 21}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var22))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</textarea> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if operation == "InsertAbove" || operation == "InsertBelow" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<textarea class=\"rounded-md md:w-full w-full h-full align-text-top items-start justify-start bg-slate-200 text-black resize-none pl-1 hover:shadow-md hover:bg-slate-400\" name=\"Description\" placeholder=\"Enter work description here...\" required=\"true\"></textarea> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if operation == "Delete" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<p class=\"rounded-md md:w-full w-full h-full align-text-top items-start justify-start bg-red-200 text-black resize-none pl-1 hover:shadow-md border border-red-500\" name=\"Description\">WARNING: This work will be permanently deleted and the gallery associated with it. If you do not wish to delete this work, press the 'X' to cancel and return.</p>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		return templ_7745c5c3_Err
	})
}

// Text for button based on component called
func TextForButton(component string) string {
	ComponentsButtonText := map[string]string{
		"Edit":        "Edit",
		"Delete":      "Delete",
		"InsertAbove": "Insert Work Above",
		"InsertBelow": "Insert Work Below",
	}
	return ComponentsButtonText[component]
}

func ButtonView(operation string, work models.WorkFrontEnd) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var23 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var23 == nil {
			templ_7745c5c3_Var23 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex flex-col h-full md:h-2/3 w-11/12 mx-4 md:w-1/2 gap-2 p-6 rounded-lg shadow-md bg-slate-200 dark:text-gray-800 md:items-start text-left justify-start mb-2\" id=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var24 string
		templ_7745c5c3_Var24, templ_7745c5c3_Err = templ.JoinStringErrs(templ.JoinStringErrs("insert-work-" + work.Position))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `editorWorkComp.templ`, Line: 248, Col: 59}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var24))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><!-- Go back-close --><h2 class=\"flex items-center gap-2 text-xl font-semibold leading-tight tracking-wide\"><button type=\"button\" class=\"flex w-6 h-6 p-1 bg-slate-200 border-none outline-none cursor-pointer hover:bg-slate-400 justify-center items-center text-center\" aria-label=\"Close\" hx-headers=\"{&#34;Content-Type&#34;: &#34;application/json&#34;}\" hx-get=\"/editor/components\" hx-vals=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var25 string
		templ_7745c5c3_Var25, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf(`{"Position": "%s", "Component": "ButtonsEditor", "Title": "%s", "Description": "%s"}`, work.Position, work.Title, work.Description))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `editorWorkComp.templ`, Line: 258, Col: 160}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var25))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-target=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var26 string
		templ_7745c5c3_Var26, templ_7745c5c3_Err = templ.JoinStringErrs(templ.JoinStringErrs("#insert-work-" + work.Position))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `editorWorkComp.templ`, Line: 259, Col: 71}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var26))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-swap=\"outerHTML\"><span class=\"text-red-500 hover:text-red-600\">X</span></button> ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var27 string
		templ_7745c5c3_Var27, templ_7745c5c3_Err = templ.JoinStringErrs(TextForButton(operation))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `editorWorkComp.templ`, Line: 264, Col: 30}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var27))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h2><!-- Form div --><div class=\"flex md:flex-row flex-col justify-center items-center text-left bg-slate-200 h-full w-full rounded-md\"><form id=\"form-container\" class=\"flex md:flex-row flex-col justify-center text-left bg-slate-200 h-full md:h-full w-full gap-4\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if operation == "InsertAbove" || operation == "InsertBelow" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" hx-post=\"/editor\" hx-include=\"[name=&#39;Title&#39;], [name=&#39;Position&#39;], [name=&#39;Description&#39;]\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if operation == "Edit" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" hx-put=\"/editor\" hx-include=\"[name=&#39;Title&#39;], [name=&#39;Position&#39;], [name=&#39;Description&#39;]\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" hx-encoding=\"multipart/form-data\" hx-trigger=\"submit\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if operation == "Delete" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" hx-post=\"/editor/del\" hx-vals=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var28 string
			templ_7745c5c3_Var28, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf(`{"Position": "%s"}`, work.Position))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `editorWorkComp.templ`, Line: 285, Col: 66}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var28))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-confirm=\"Are you sure you want to delete?\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("><!-- Title input --><div class=\"flex md:flex-col flex-col justify-end gap-2 items-center bg-slate-200 h-full w-full rounded-md md:h-full flex\"><div class=\"flex justify-center items-center border border-blue-100 w-full h-1/4 md:h-1/3 md:h-1/5 sm:text-sm md:text-lg rounded-lg dark:text-gray-800 bg-slate-200 text-center shadow-sm shadow-slate-500\"><!-- Generate <input> for 'Edit','Insert' or <p> message for 'Delete'-->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = GetInputsOrMessage(operation, work.Title).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div class=\"rounded-md md:w-full  border border-blue-100 w-full h-5/6 md:h-4/5 align-text-top items-start justify-start bg-slate-200 text-black flex-1 resize-none shadow-sm shadow-slate-500\"><!-- Generate <text area> or <p> with message depending on operation-->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = TextAreaContent(operation, work.Description).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if work.PositionBelow != "" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input type=\"hidden\" name=\"Position\" value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var29 string
			templ_7745c5c3_Var29, templ_7745c5c3_Err = templ.JoinStringErrs(work.PositionBelow)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `editorWorkComp.templ`, Line: 299, Col: 73}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var29))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input type=\"hidden\" name=\"Position\" value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var30 string
			templ_7745c5c3_Var30, templ_7745c5c3_Err = templ.JoinStringErrs(work.Position)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `editorWorkComp.templ`, Line: 300, Col: 74}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var30))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div><!-- /Title Input --><!-- File upload  -->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if operation == "Delete" {
			templ_7745c5c3_Err = DeleteView().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			templ_7745c5c3_Err = FileUploadView(operation).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!-- /File upload --></form></div><!-- /Form div --></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
