package jobs

import (
    "github.com/albrow/jobs"
    "qixalite.com/Ranndom/ldap-portal/modules/settings"
)

var (
    UpdateUserJob *jobs.Type
)

func InitJobs() {
    jobs.Config.Db.Address = settings.Jobs.Address
    jobs.Config.Db.Network = settings.Jobs.Network
    jobs.Config.Db.Database = settings.Jobs.Database
    jobs.Config.Db.Password = settings.Jobs.Password

    UpdateUserJob, _ = jobs.RegisterType("updateUser", 3, UpdateUserRun)
}

