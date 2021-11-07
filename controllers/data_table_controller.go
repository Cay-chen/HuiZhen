package controllers

type DataTableController struct {
	BaseController
}

func (c *DataTableController) Get() {

	if c.IsLogin {

	} else {
		c.Redirect("/error/405", 302)
	}

}
