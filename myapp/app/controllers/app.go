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
	(*models.DB).AutoMigrate(&models.Table2{})
	return c.Redirect(App.Index)
}

func (c App) InsertSample() revel.Result {

	table1_model := models.Table1{
		Title: "タイトル１",
		Number: 1,
	}

	(*models.DB).Create(&table1_model)

	return c.Redirect(App.Index)
}

func (c App) InsertSample2() revel.Result {

	table2_model1 := models.Table2{
		Title: "タイトル１",
		Number: 1,
		Table1_id: 1,
	}
	table2_model2 := models.Table2{
		Title: "タイトル2",
		Number: 2,
		Table1_id: 1,
	}

	(*models.DB).Create(&table2_model1)
	(*models.DB).Create(&table2_model2)

	return c.Redirect(App.Index)
}

func (c App) JointStruct() revel.Result {

	var join_form []forms.Table1JoinTable2

	(*models.DB).Table("table1").
		Select(`
			table1.title as table1_title,
			table1.number as table1_number,
			table2.title as table2_title,
			table2.number as table2_number`).
		Joins(`JOIN table2 ON table1.ID = table2.table1_id`).
		Scan(&join_form)

	return c.Redirect(App.Index)
}
