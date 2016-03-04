package routes

import (
    "gopkg.in/macaron.v1"
    "github.com/go-macaron/binding"
    "qixalite.com/Ranndom/ldap-portal/controllers"
    "qixalite.com/Ranndom/ldap-portal/models"
)

func Account(m *macaron.Macaron) {
    m.Group("/account", func() {
	m.Get("/details", controllers.AccountDetails)
	m.Get("/ssh_keys", controllers.AccountSSHKeys)
	m.Get("/change_password", controllers.AccountChangePassword)

	m.Post("/details",
            binding.Bind(models.AccountDetailsForm{}),
            controllers.POSTAccountDetails,
    	)
        m.Post("/ssh_keys", controllers.POSTAccountSSHKeys)
        m.Post("/change_password",
            binding.Bind(models.AccountChangePasswordForm{}),
            controllers.POSTAccountChangePassword,
        )
    })
}

