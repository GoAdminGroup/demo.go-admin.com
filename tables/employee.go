package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	template2 "html/template"
)

func GetEmployeeTable(ctx *context.Context) table.Table {

	employeeTable := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := employeeTable.GetInfo().HideFilterArea()

	info.AddField("ID", "id", db.Int).FieldFilterable()
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Gender", "gender", db.Tinyint).FieldDisplay(func(value types.FieldModel) interface{} {
		if value.Value == "0" {
			return "men"
		}
		return "women"
	})
	info.AddField("Department", "department", db.Varchar)
	info.AddField("Phone", "phone", db.Varchar)
	info.AddField("Job", "job", db.Varchar)

	department := ctx.Query("department")

	info.SetTable("employee").SetTitle("Employee").SetDescription("Employee").
		SetWrapper(func(content template2.HTML) template2.HTML {
			col1 := `<div style="margin-left:243px;">` + content + `</div>`

			tree := template.Default().TreeView().SetTree(types.TreeViewData{
				Data: types.TreeViewItems{
					{
						Text: "Qjoy",
						Href: "/admin/info/employee?__go_admin_no_animation_=true",
						Nodes: types.TreeViewItems{
							{
								Text: "Tech",
								State: types.TreeViewItemState{
									Expanded: department == "Frontend" || department == "Middle" || department == "Backend",
								},
								Nodes: types.TreeViewItems{
									{
										Text: "Frontend",
										Href: "/admin/info/employee?department=Frontend&__go_admin_no_animation_=true",
										State: types.TreeViewItemState{
											Checked:  department == "Frontend",
											Selected: department == "Frontend",
										},
									}, {
										Text: "Middle",
										Href: "/admin/info/employee?department=Middle&__go_admin_no_animation_=true",
										State: types.TreeViewItemState{
											Checked:  department == "Middle",
											Selected: department == "Middle",
										},
									}, {
										Text: "Backend",
										Href: "/admin/info/employee?department=Backend&__go_admin_no_animation_=true",
										State: types.TreeViewItemState{
											Checked:  department == "Backend",
											Selected: department == "Backend",
										},
									},
								},
							}, {
								Text: "Sale",
								Href: "/admin/info/employee?department=Sale&__go_admin_no_animation_=true",
								State: types.TreeViewItemState{
									Checked:  department == "Sale",
									Selected: department == "Sale",
								},
							}, {
								Text: "Reception",
								Href: "/admin/info/employee?department=Reception&__go_admin_no_animation_=true",
								State: types.TreeViewItemState{
									Checked:  department == "Reception",
									Selected: department == "Reception",
								},
							}, {
								Text: "Hr",
								Href: "/admin/info/employee?department=Hr&__go_admin_no_animation_=true",
								State: types.TreeViewItemState{
									Checked:  department == "Hr",
									Selected: department == "Hr",
								},
							},
						},
					},
				},
				ExpandIcon:        "fa fa-angle-right",
				CollapseIcon:      "fa fa-angle-down",
				SelectedBackColor: "#fbfbfb",
				SelectedColor:     "#333333",
				EnableLinks:       true,
			}).GetContent()

			col2 := `<div style="position: absolute;width:230px;">` + template.Default().Box().SetHeader("Organization").
				WithHeadBorder().SetBody(tree).GetContent() + `</div>`
			return `<div style="width:100%;">` + col2 + col1 + `</div>`
		})

	formList := employeeTable.GetForm()

	formList.AddField("ID", "id", db.Int, form.Default).FieldNotAllowAdd()
	formList.AddField("Name", "name", db.Varchar, form.Text)
	formList.AddField("Gender", "gender", db.Tinyint, form.Number)
	formList.AddField("Department", "department", db.Varchar, form.Text)
	formList.AddField("Phone", "phone", db.Varchar, form.Text)
	formList.AddField("Job", "job", db.Varchar, form.Text)

	formList.SetTable("employee").SetTitle("Employee").SetDescription("Employee")

	return employeeTable
}
