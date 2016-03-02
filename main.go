package main

import (
    "gopkg.in/macaron.v1"
    "github.com/go-macaron/session"
    "github.com/go-macaron/csrf"
    "github.com/go-macaron/pongo2"
    "qixalite.com/Ranndom/ldap-portal/controllers"
    "qixalite.com/Ranndom/ldap-portal/middleware"
    "qixalite.com/Ranndom/ldap-portal/modules/settings"

    "log"
)

func main() {
    m := CreateWeb()

    RegisterRoutes(m)

    // Debugging settings
    settings.NewContext()
    log.Printf(settings.LDAP.Hostname)

    m.Run()
}

func CreateWeb() *macaron.Macaron {
    m := macaron.New()

    // Enable logger middleware
    m.Use(macaron.Logger())

    // Enable static file serving
    m.Use(macaron.Static(
        "public",
        macaron.StaticOptions{
            SkipLogging: false,
        },
    ))

    // Enable template rendering
    m.Use(pongo2.Pongoer(pongo2.Options{
        Directory: "views",
        Extensions: []string{".tmpl"},
        Charset: "UTF-8",
        IndentJSON: true,
        IndentXML: true,
        HTMLContentType: "text/html",
    }))

    // Enable sessions
    m.Use(session.Sessioner())

    // Enable CSRF protection
    m.Use(csrf.Csrfer(csrf.Options{
        Secret: "hcmeit87fksmcuf9r4jfhg987fhfmvivcicvgshvochg38fgn",
        SetCookie: true,
        Header: "X-CSRF-Token",
    }))

    m.Use(middleware.TemplateSessioner())

    return m
}

func RegisterRoutes(m *macaron.Macaron) {
    m.Get("/", controllers.Home)

    m.Group("/account", func() {
        m.Get("/details", controllers.AccountDetails)
        m.Get("/ssh_keys", controllers.AccountSSHKeys)
        m.Get("/change_password", controllers.AccountChangePassword)

        m.Post("/details", controllers.POSTAccountDetails)
        m.Post("/ssh_keys", controllers.POSTAccountSSHKeys)
        m.Post("/change_password", controllers.POSTAccountChangePassword)
    })

    m.Group("/auth", func() {
        m.Get("/login", controllers.AuthLogin)
        m.Get("/logout", controllers.AuthLogout)

        m.Post("/login", controllers.POSTAuthLogin)
    })
}

