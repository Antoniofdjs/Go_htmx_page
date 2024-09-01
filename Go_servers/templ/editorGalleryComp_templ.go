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

func GalleryEditor(work models.WorkFrontEnd, galleryItems []models.GalleryItemFrontEnd) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<section class=\"px-0 dark:bg-gray-800 dark:text-gray-900 justify-center w-screen h-auto items-center flex flex-col\" id=\"gallery-section\"><!-- Work Container --><div class=\"flex flex-col dark:bg-gray-50 dark:text-gray-800 justify-center items-center w-full h-600 text-left\"><!-- Cover Picture Container  --><div class=\"h-full w-full flex\" id=\"cover-image-container\"><img src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(work.Path)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `editorGalleryComp.templ`, Line: 15, Col: 27}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" alt=\"\" class=\"object-cover h-full w-full rounded-sm dark:bg-gray-500 aspect-video\" id=\"cover-image\" loading=\"lazy\"></div></div><!-- / Work Container --><!-- Work Info Container --><div id=\"work-info-container\" class=\"w-full h-1/4 md:justify-center justify-center flex flex-col text-left md:text-center md:items-center py-12\"><h2 id=\"work-title\" class=\"text-xl md:text-3xl font-semibold text-center\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(work.Title)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `editorGalleryComp.templ`, Line: 27, Col: 98}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h2><span class=\"block md:pb-1 md:pt-2 pb-4 text-sm dark:text-gray-600 font-thin text-center\">07/26/2024</span><div class=\"md:w-3/5 w-full font-serif text-lg h-auto flex justify-center items-center\"><p id=\"work-description\" class=\"md:w-full w-3/4 font-serif text-lg h-auto text-center\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(work.Description)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `editorGalleryComp.templ`, Line: 31, Col: 34}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p></div></div><!-- / Work Info Container --><!-- Gallery Container --><div class=\"px-0 dark:bg-gray-800 dark:text-gray-900 justify-center items-start flex pb-4 w-full h-auto\" id=\"gallery-container\"><!-- Gallery Grid --><div class=\" flex h-auto grid md:grid-cols-3 auto-cols-max grid-cols-2 w-11/12 md:gap-3 gap-1\" id=\"gallery-grid\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, item := range galleryItems {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"w-full md:h-96 h-80 flex relative\" id=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var5 string
			templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(templ.JoinStringErrs("image-container-" + item.Position))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `editorGalleryComp.templ`, Line: 43, Col: 122}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><img class=\"w-full h-full object-cover\" src=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var6 string
			templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(item.Path)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `editorGalleryComp.templ`, Line: 46, Col: 39}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" alt=\"\" loading=\"lazy\" id=\"my-pic\"> <button type=\"button\" id=\"delete-pic\" class=\"absolute bottom-2 right-2 h-12 w-12 bg-red-500 text-black px-1 py-1 justify-center items-center flex rounded-md shadow  shadow-md shadow-slate-700 hover:bg-red-400 hover:shadow-none transition duration-150 ease-in-out\" hx-get=\"/editor/update\" hx-target=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var7 string
			templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(templ.JoinStringErrs("#image-container-" + item.Position))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `editorGalleryComp.templ`, Line: 56, Col: 89}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-swap=\"innerHTML transition:true\" hx-vals=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var8 string
			templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf(`{"Opacity": "true", "PicUrl": "%s", "Position": "%s"}`, item.Path, item.Position))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `editorGalleryComp.templ`, Line: 58, Col: 124}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><img src=\"../static/images/trash-icon.png\" alt=\"Delete\" class=\"h-full w-full object-contain\"></button></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"w-full md:h-96 h-80 bg-slate-200  rounded-lg shadow-md shadow-slate-600 flex justify-center items-center flex-col overflow-hidden\" id=\"file-upload-container\"><form hx-post=\"/editor/gallery\" hx-trigger=\"change from:#upload-pics\" hx-target=\"#files-list\" hx-swap=\"innerHTML\" hx-include=\"#upload-pics\" enctype=\"multipart/form-data\" class=\"h-1/3 md:h-1/2 w-full flex justify-center items-center\"><label for=\"upload-pics\" class=\"flex hover:opacity-70 hover:cursor-pointer\"><svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 32 31\" class=\" w-20 w-20 md:h-28 md:w-28 fill-current text-gray-500 rounded-lg shadow shadow-lg shadow-slate-700 hover:text-emerald-500 hover:shadow-slate-400 transition duration-150 ease-in-out\" style=\"fill-rule: evenodd;\"><path d=\"m30 19.59-4.29-4.29a1 1 0 0 0-1.41 0L19 20.59l-6.29-6.29a1 1 0 0 0-1.41 0L2 23.59V3a1 1 0 0 1 1-1h18V0H3a3 3 0 0 0-3 3v26a3 3 0 0 0 3 3h26a3 3 0 0 0 3-3V12h-2z\"></path> <path d=\"M10 8a4 4 0 1 0 4-4 4 4 0 0 0-4 4zm6 0a2 2 0 1 1-2-2 2 2 0 0 1 2 2zM28 4V0h-2v4h-4v2h4v4h2V6h4V4h-4z\"></path></svg> <input type=\"file\" name=\"pictures\" id=\"upload-pics\" multiple class=\"hidden\" accept=\"image/*\"></label></form><div class=\"w-full h-2/3 md:h-1/2 overflow-y-auto  justify-around flex flex-col gap-2 items-center\" id=\"files-list\"><div class=\"w-10/12 h-3/4 justify-center items-center flex flex-col mt-2\"><p class=\"text-slate-500\">No files selected</p></div></div></div></div><!-- /Gallery Grid --></div><!-- /Gallery Container --></section>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func UpdatePicStatus(opacity string, picUrl string, position string) templ.Component {
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
		templ_7745c5c3_Var9 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var9 == nil {
			templ_7745c5c3_Var9 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<img class=\"w-full h-full object-cover\" src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var10 string
		templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(picUrl)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `editorGalleryComp.templ`, Line: 100, Col: 16}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" alt=\"\" loading=\"lazy\" id=\"my-pic\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if opacity == "true" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" style=\"opacity: 0.50; filter: blur(2px)\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("> <button type=\"button\" id=\"delete-pic\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if opacity == "false" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" class=\"absolute bottom-2 right-2 h-12 w-12 bg-red-500 text-black px-1 py-1 justify-center items-center flex rounded-md shadow  shadow-md shadow-slate-700 hover:bg-red-400 hover:shadow-none transition duration-150 ease-in-out\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if opacity == "true" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" class=\"absolute bottom-2 right-2 h-12 w-12 bg-emerald-200 border border-emerald-700 text-black px-1 py-1 justify-center items-center flex rounded-md shadow-md shadow-slate-700 hover:bg-emerald-400 hover:shadow-none transition duration-150 ease-in-out\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" hx-get=\"/editor/update\" hx-target=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var11 string
		templ_7745c5c3_Var11, templ_7745c5c3_Err = templ.JoinStringErrs(templ.JoinStringErrs("#image-container-" + position))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `editorGalleryComp.templ`, Line: 118, Col: 67}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var11))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-swap=\"innerHTML transition:true\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if opacity == "true" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" hx-vals=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var12 string
			templ_7745c5c3_Var12, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf(`{"Opacity": "false", "PicUrl": "%s", "Position": "%s"}`, picUrl, position))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `editorGalleryComp.templ`, Line: 121, Col: 100}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var12))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if opacity == "false" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" hx-vals=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var13 string
			templ_7745c5c3_Var13, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf(`{"Opacity": "true", "PicUrl": "%s", "Position": "%s"}`, picUrl, position))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `editorGalleryComp.templ`, Line: 124, Col: 99}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var13))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if opacity == "true" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<img src=\"../static/images/undo-symboll.png\" alt=\"Delete\" class=\"h-full w-full object-contain\"> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if opacity == "false" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<img src=\"../static/images/trash-icon.png\" alt=\"Delete\" class=\"h-full w-full object-contain\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</button>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

/* List the files seclecte to be uploaded */

func FilesSelectedContainer() templ.Component {
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
		templ_7745c5c3_Var14 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var14 == nil {
			templ_7745c5c3_Var14 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<ul class=\"w-10/12 max-h-3/4 h-auto overflow-y-auto justify-around items-center flex flex-col gap-1 border border border-slate-400 rounded-lg mt-2\"><li class=\"w-10/12 px-2 py-2 text-gray-600\">File 1 Pic Name Here</li><li class=\"w-10/12 px-2 py-2 text-gray-600\">File 1 Pic Name Here</li><li class=\"w-10/12 px-2 py-2 text-gray-600\">File 1 Pic Name Here</li><li class=\"w-10/12 px-2 py-2 text-gray-600\">File 1 Pic Name Here</li><li class=\"w-10/12 px-2 py-2 text-gray-600\">File 1 Pic Name Here</li><li class=\"w-10/12 px-2 py-2 text-gray-600\">File 1 Pic Name Here</li><li class=\"w-10/12 px-2 py-2 text-gray-600\">File 1 Pic Name Here</li><li class=\"w-10/12 px-2 py-2 text-gray-600\">File 1 Pic Name Here</li><li class=\"w-10/12 px-2 py-2 text-gray-600\">File 1 Pic Name Here</li><li class=\"w-10/12 px-2 py-2 text-gray-600\">File 1 Pic Name Here</li></ul><div class=\"w-full h-1/4 flex justify-end items-center pb-2\"><button class=\"w-auto h-auto bg-emerald-500 px-1 py-1 mr-2 rounded-lg\">Upload Pictures </button></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
