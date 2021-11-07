package controllers

type DepMangeController struct {
	BaseController
}

func (c *DepMangeController) Get() {
	if c.IsLogin {
		if c.IsAdmin || c.IsYwb {
			c.TplName = "iframe_dev_manage.html"
		}
	} else {
		c.Redirect("/error/1245", 302)
	}

}
