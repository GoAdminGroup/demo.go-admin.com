package pages

import (
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db"
	template2 "github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetForm1Content() (types.Panel, error) {

	components := template2.Get(config.Get().Theme)

	aform := components.Form().
		SetTabHeaders([]string{"input", "select"}).
		SetTabContents([][]types.FormField{
			{
				{
					Field:    "name",
					TypeName: db.Varchar,
					Head:     "Name",
					Default:  "jane",
					Editable: true,
					FormType: form.Text,
					Value:    "jane",
					Options:  types.FieldOptions{},
				},
				{
					Field:    "age",
					TypeName: "int",
					Head:     "Age",
					Default:  "11",
					Value:    "11",
					Editable: true,
					FormType: form.Number,
					Options:  types.FieldOptions{},
				},
				{
					Field:    "homepage",
					TypeName: db.Varchar,
					Head:     "HomePage",
					Default:  "http://google.com",
					Editable: true,
					FormType: form.Url,
					Value:    "",
					Options:  types.FieldOptions{},
				},
				{
					Field:    "email",
					TypeName: db.Varchar,
					Head:     "Email",
					Default:  "xxxx@xxx.com",
					Editable: true,
					FormType: form.Email,
					Value:    "",
					Options:  types.FieldOptions{},
				},
				{
					Field:    "birthday",
					TypeName: db.Varchar,
					Head:     "Birthday",
					Default:  "2010-09-05",
					Editable: true,
					FormType: form.Datetime,
					Value:    "",
					Options:  types.FieldOptions{},
				},
				{
					Field:    "password",
					TypeName: db.Varchar,
					Head:     "Password",
					Default:  "",
					Editable: true,
					FormType: form.Password,
					Value:    "",
					Options:  types.FieldOptions{},
				},
				{
					Field:    "ip",
					TypeName: db.Varchar,
					Head:     "Ip",
					Default:  "",
					Editable: true,
					FormType: form.Ip,
					Value:    "",
					Options:  types.FieldOptions{},
				},
				{
					Field:    "currency",
					TypeName: db.Int,
					Head:     "Currency",
					Default:  "",
					Editable: true,
					FormType: form.Currency,
					Value:    "",
					Options:  types.FieldOptions{},
				},
				{
					Field:    "content",
					TypeName: db.Text,
					Head:     "Content",
					Default:  "",
					Editable: true,
					FormType: form.RichText,
					Value:    "",
					Options:  types.FieldOptions{},
				},
			},
			{
				{
					Field:    "website",
					TypeName: db.Tinyint,
					Head:     "Switch",
					Default:  "",
					Editable: true,
					FormType: form.Switch,
					Value:    "",
					Options: types.FieldOptions{
						{Text: "website", Value: "0"},
						{Text: "website", Value: "1"},
					},
				},
				{
					Field:    "fruit",
					TypeName: db.Varchar,
					Head:     "Fruit",
					Default:  "",
					Editable: true,
					FormType: form.SelectBox,
					Value:    "",
					Options: types.FieldOptions{
						{Text: "apple", Value: "apple"},
						{Text: "banana", Value: "banana"},
						{Text: "watermelon", Value: "watermelon"},
						{Text: "pear", Value: "pear"},
					},
					FieldDisplay: types.FieldDisplay{
						Display: func(value types.FieldModel) interface{} {
							return []string{"梨"}
						},
					},
				},
				{
					Field:    "gender",
					TypeName: db.Tinyint,
					Head:     "Gender",
					Default:  "0",
					Editable: true,
					FormType: form.Radio,
					Value:    "",
					Options: types.FieldOptions{
						{Text: "male", Value: "0"},
						{Text: "female", Value: "1"},
					},
				},
				{
					Field:    "drink",
					TypeName: db.Varchar,
					Head:     "Drink",
					Default:  "beer",
					Editable: true,
					FormType: form.Select,
					Value:    "",
					Options: types.FieldOptions{
						{Text: "beer", Value: "beer"},
						{Text: "juice", Value: "juice"},
						{Text: "water", Value: "water"},
						{Text: "red bull", Value: "red bull"},
					},
				},
				{
					Field:    "experience",
					TypeName: db.Tinyint,
					Head:     "Work experience",
					Default:  "",
					Editable: true,
					FormType: form.SelectSingle,
					Value:    "",
					Options: types.FieldOptions{
						{Text: "two years", Value: "0"},
						{Text: "three years", Value: "1"},
						{Text: "four years", Value: "2"},
						{Text: "five years", Value: "3"},
					},
				},
			},
		}).
		SetPrefix(config.Get().PrefixFixSlash()).
		SetUrl("/").
		SetTitle("Form").
		SetInfoUrl("/admin").
		GetContent()

	return types.Panel{
		Content:     aform,
		Title:       "Form",
		Description: "form example",
	}, nil
}
