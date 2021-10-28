package controllers

type ExitControoler struct {
	BaseController
}

func (c *ExitControoler) Get() {
	err := c.DelSession("loginUser")
	if err != nil {
		return
	}
	c.Redirect("/", 302)

}
