package controllers

import "github.com/revel/revel"

type App struct {
    *revel.Controller
}

func (c App) Index() revel.Result {
    return c.Render()
}

func (c App) Account() revel.Result {
    return c.Render()
}

func (c App) POST_Account(givenName string, surname string) revel.Result {
    c.Validation.Required(givenName).Message("Please enter a first name")
    c.Validation.Required(surname).Message("Please enter a surname")

    if c.Validation.HasErrors() {
        c.Validation.Keep()
        c.FlashParams()
        return c.Redirect(App.Account)
    }

    c.Flash.Success("Updated account!")

    return c.Redirect(App.Account)
}

