package forms

type InputBindFrom struct {
  Food string
  Number int
  Bool bool
}

type JsonUnmarshalForm struct {
  Name string `json:"name"`
  Number int `json:"number"`
}

type NestParent struct {
  Sport []Sport `json:"sport" validate:"dive"`
}

type Sport struct {
  Name   string `json:"name" validate:"required"`
  Member int    `json:"member" validate:"required"`
}

type Table1JoinTable2 struct {
  Table1_title string
  Table1_number int
  Table2_title string
  Table2_number int
}
