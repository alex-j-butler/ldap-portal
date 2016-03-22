package main

import (
    "fmt"
    "strings"
    "net/http"

    "gopkg.in/macaron.v1"
    "github.com/go-macaron/session"
    "github.com/go-macaron/csrf"
    "github.com/go-macaron/pongo2"
    "github.com/go-macaron/binding"
    "github.com/go-macaron/bindata"
    "qixalite.com/Ranndom/ldap-portal/public"
    "qixalite.com/Ranndom/ldap-portal/controllers"
    "qixalite.com/Ranndom/ldap-portal/models"
    "qixalite.com/Ranndom/ldap-portal/middleware"
    "qixalite.com/Ranndom/ldap-portal/modules/settings"
    "qixalite.com/Ranndom/ldap-portal/modules/database"
    "qixalite.com/Ranndom/ldap-portal/modules/jobs"
    "qixalite.com/Ranndom/ldap-portal/modules/helpers"
    "qixalite.com/Ranndom/ldap-portal/modules/logging"

    _ "github.com/go-macaron/session/redis"

    pongo "gopkg.in/flosch/pongo2.v3"
)

func main() {
    // Load settings
    settings.NewContext()

    // Load loggers
    logging.NewContext()

    // Load database
    database.InitDatabase()

    // Init jobs
    jobs.InitJobs()

    m := CreateWeb()
    RegisterRoutes(m)

    logging.AppLogger.Info("Listening on %s:%d", settings.Web.Address, settings.Web.Port)
    http.ListenAndServe(fmt.Sprintf("%s:%d", settings.Web.Address, settings.Web.Port), m)
//    m.Run()
}

func CreateWeb() *macaron.Macaron {
    m := macaron.New()

    // Enable logger middleware
    // m.Use(macaron.Logger())
    m.Use(middleware.Logger())

    // Enable static file serving
    m.Use(macaron.Static(
        "public",
        macaron.StaticOptions{
            SkipLogging: false,
            FileSystem: bindata.Static(bindata.Options{
                Asset: public.Asset,
                AssetDir: public.AssetDir,
                AssetNames: public.AssetNames,
                AssetInfo: public.AssetInfo,
                Prefix: "",
            }),
        },
    ))

    m.Use(macaron.Recovery())

    // Enable template rendering
    AddPongoFilters()
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

    m.Group("/profile", func() {
        m.Get("/:name", controllers.UserProfile)
    })

    m.Group("/account", func() {
        m.Get("/details", controllers.AccountDetails)
        m.Get("/change_password", controllers.AccountChangePassword)

        m.Post("/details",
            binding.BindIgnErr(models.AccountDetailsForm{}),
            controllers.POSTAccountDetails,
        )
        m.Post("/change_password",
            binding.BindIgnErr(models.AccountChangePasswordForm{}),
            controllers.POSTAccountChangePassword,
        )
    }, helpers.IsLoggedIn)

    m.Group("/auth", func() {
        m.Get("/login", controllers.AuthLogin)
        m.Get("/logout", controllers.AuthLogout)

        m.Post("/login",
            binding.BindIgnErr(models.LoginForm{}),
            controllers.POSTAuthLogin,
        )
    })
}

func RegisterModels() {
    database.RegisterModel(&models.User{})
}

func AddPongoFilters() {
    pongo.RegisterFilter("split", FilterSplit)
}

func FilterSplit(in *pongo.Value, param *pongo.Value) (*pongo.Value, *pongo.Error) {
    return pongo.AsValue(strings.Split(in.String(), param.String())), nil
}
