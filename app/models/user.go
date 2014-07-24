package models

type User struct {
  Name          string `sql:"size:255"`
  Email         string `sql:"size:255"`
  Description   string `sql:"size:255"`
}
