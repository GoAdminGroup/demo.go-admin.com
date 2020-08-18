package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	editType "github.com/GoAdminGroup/go-admin/template/types/table"
	template2 "html/template"
)

// GetPostsTable return the model of table posts.
func GetPostsTable(ctx *context.Context) (postsTable table.Table) {

	postsTable = table.NewDefaultTable(table.DefaultConfig().SetExportable(true))

	info := postsTable.GetInfo()
	info.AddField("ID", "id", db.Int).FieldSortable()
	info.AddField("Title", "title", db.Varchar)
	info.AddField("AuthorID", "author_id", db.Int).FieldDisplay(func(value types.FieldModel) interface{} {
		return template.Default().
			Link().
			SetURL("/admin/info/authors/detail?__goadmin_detail_pk=" + value.Value).
			SetContent(template2.HTML(value.Value)).
			OpenInNewTab().
			SetTabTitle(template.HTML("Author Detail(" + value.Value + ")")).
			GetContent()
	})
	info.AddField("AuthorName", "name", db.Varchar).FieldDisplay(func(value types.FieldModel) interface{} {
		first, _ := value.Row["authors_goadmin_join_first_name"].(string)
		last, _ := value.Row["authors_goadmin_join_last_name"].(string)
		return first + " " + last
	})
	info.AddField("AuthorFirstName", "first_name", db.Varchar).FieldJoin(types.Join{
		Field:     "author_id",
		JoinField: "id",
		Table:     "authors",
	}).FieldHide()
	info.AddField("AuthorLastName", "last_name", db.Varchar).FieldJoin(types.Join{
		Field:     "author_id",
		JoinField: "id",
		Table:     "authors",
	}).FieldHide()
	info.AddField("Description", "description", db.Varchar)
	info.AddField("Content", "content", db.Varchar).FieldEditAble(editType.Textarea)
	info.AddField("Date", "date", db.Varchar)

	info.SetTable("posts").SetTitle("Posts").SetDescription("Posts")

	formList := postsTable.GetForm()
	formList.AddField("ID", "id", db.Int, form.Default).FieldDisplayButCanNotEditWhenUpdate().FieldNotAllowAdd()
	formList.AddField("Title", "title", db.Varchar, form.Text)
	formList.AddField("Description", "description", db.Varchar, form.Text)
	formList.AddField("Content", "content", db.Varchar, form.RichText).FieldEnableFileUpload()
	formList.AddField("Date", "date", db.Varchar, form.Date)

	formList.SetWrapper(func(content template2.HTML) template2.HTML {
		tableContent := template.Default().Table().SetThead(types.Thead{
			{Head: "Total amount of reading"},
			{Head: "Total subscription amount"},
			{Head: "Viewer Today"},
			{Head: "New users"},
			{Head: "Preserve rate"},
		}).SetInfoList([]map[string]types.InfoItem{
			{
				"Total amount of reading": {Content: "1223"},
				"amount":                  {Content: "1433"},
				"Viewer Today":            {Content: "230"},
				"New users":               {Content: "20"},
				"Preserve rate":           {Content: "50%"},
			},
		}).GetContent()
		return template.Default().Box().
			SetBody(tableContent).
			SetNoPadding().
			WithHeadBorder().
			GetContent() + content
	})
	formList.EnableAjax("Success", "Fail")

	formList.SetTable("posts").SetTitle("Posts").SetDescription("Posts")

	return
}
