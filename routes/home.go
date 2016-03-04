package routes

import (
    "gopkg.in/macaron.v1"
    "qixalite.com/Ranndom/ldap-portal/controllers"
)

func Home(m *macaron.Macaron) {
    m.Get("/", controllers.Home)
}

