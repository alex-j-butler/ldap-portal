package jobs

import (
    jobs_lib "github.com/albrow/jobs"
    "qixalite.com/Ranndom/ldap-portal/modules/settings"

    "log"
)

var (
    UpdateUserJob *jobs_lib.Type
)

func InitJobs() {
    log.Printf("Memes!")

    jobs_lib.Config.Db.Address = settings.Jobs.Address
    jobs_lib.Config.Db.Network = settings.Jobs.Network
    jobs_lib.Config.Db.Database = settings.Jobs.Database
    jobs_lib.Config.Db.Password = settings.Jobs.Password

    var err error

    pool, err := jobs_lib.NewPool(nil)
    if err != nil {
        log.Printf("%s", err)
    }
    if err := pool.Start(); err != nil {
        log.Printf("%s", err)
    }

    UpdateUserJob, err = jobs_lib.RegisterType("updateUser", 3, UpdateUserRun)
    if err != nil {
        log.Printf("%s", err)
    }
}

