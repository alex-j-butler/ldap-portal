package main

import (
    "gopkg.in/macaron.v1"
    "github.com/go-macaron/session"
    "github.com/go-macaron/csrf"
    "github.com/go-macaron/pongo2"
    "qixalite.com/Ranndom/ldap-portal/controllers"
    "qixalite.com/Ranndom/ldap-portal/routes"
    "qixalite.com/Ranndom/ldap-portal/middleware"
    "qixalite.com/Ranndom/ldap-portal/modules/settings"
    "qixalite.com/Ranndom/ldap-portal/modules/database"
    "qixalite.com/Ranndom/ldap-portal/modules/jobs"
)

func main() {
    // Load settings
    settings.NewContext()

    // Load database
    database.InitDatabase()

    // Init jobs
    jobs.InitJobs()

    m := CreateWeb()
    RegisterRoutes(m)
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
    m.Use(session.Sessioner(session.Options{
        Provider: settings.Session.Provider,
        ProviderConfig: settings.Session.ProviderConfig,
        CookieName: settings.Session.CookieName,
        Secure: settings.Session.Secure,
        IDLength: settings.Session.IDLength,
    }))

    // Enable CSRF protection
    m.Use(csrf.Csrfer(csrf.Options{
        Secret: settings.CSRF.Secret,
        SetCookie: true,
        Header: "X-CSRF-Token",
    }))

    // Enable template sessions
    m.Use(middleware.TemplateSessioner())

    return m
}

func RegisterRoutes(m *macaron.Macaron) {
    m.Get("/", controllers.Home)

    routes.Account(m)
    routes.Auth(m)
}

