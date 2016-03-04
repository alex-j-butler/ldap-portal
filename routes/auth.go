package routes

import (
    "gopkg.in/macaron.v1"
    "github.com/go-macaron/binding"
    "qixalite.com/Ranndom/ldap-portal/controllers"
    "qixalite.com/Ranndom/ldap-portal/models"
)

func Auth(m *macaron.Macaron) {
    m.Group("/auth", func() {
        m.Get("/login", controllers.AuthLogin)
        m.Get("/logout", controllers.AuthLogout)

        m.Post("/login",
            binding.Bind(models.LoginForm{}),
            controllers.POSTAuthLogin,
        )   
    })
}

