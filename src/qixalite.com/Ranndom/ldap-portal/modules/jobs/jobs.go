package jobs

import (
    jobs_lib "github.com/albrow/jobs"
    "qixalite.com/Ranndom/ldap-portal/modules/settings"
)

var (
    UpdateUserJob *jobs_lib.Type
)

func InitJobs() {
    jobs_lib.Config.Db.Address = settings.Jobs.Address
    jobs_lib.Config.Db.Network = settings.Jobs.Network
    jobs_lib.Config.Db.Database = settings.Jobs.Database
    jobs_lib.Config.Db.Password = settings.Jobs.Password


    UpdateUserJob, _ = jobs_lib.RegisterType("updateUser", 3, UpdateUserRun)
}

