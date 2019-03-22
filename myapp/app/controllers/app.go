package controllers

import (
	"encoding/json"
	"myapp/app/forms"
	"fmt"
	"github.com/revel/revel"
	"myapp/app/models"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) BindInput() revel.Result {

	var input_bind_form forms.InputBindFrom
	c.Params.Bind(&input_bind_form, "InputBindFrom")

	return c.Redirect(App.Index)
}

func (c App) JsonUnmarshal() revel.Result {

	json_str := `{"name":"野球", "number": 9}`

	var json_unmarshal_form forms.JsonUnmarshalForm
	json_bytes := ([]byte)(json_str)
	err := json.Unmarshal(json_bytes, &json_unmarshal_form)
	fmt.Println(err)
	fmt.Println(json_unmarshal_form)

	return c.Redirect(App.Index)
}

func (c App) CreateTableSample() revel.Result {
	(*models.DB).AutoMigrate(&models.Table1{})

	return c.Redirect(App.Index)
}
