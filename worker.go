package main

import (
    jobs_lib "github.com/albrow/jobs"
    "qixalite.com/Ranndom/ldap-portal/modules/jobs"
    "qixalite.com/Ranndom/ldap-portal/modules/settings"

    "log"
)

func main() {
    // Load settings
    settings.NewContext()

    // Init jobs
    jobs.InitJobs()
    StartPool()
}

func StartPool() {
    pool, err := jobs_lib.NewPool(&jobs_lib.PoolConfig{
        NumWorkers: settings.JobsWorker.MaxWorkers,
        BatchSize: settings.JobsWorker.BatchSize,
    })

    if err != nil {
        log.Fatal("Error creating pool: ", err)
    }

    if err := pool.Start(); err != nil {
        log.Fatal("Error starting pool: ", err)
    }

    defer func() {
        pool.Close()
        if err := pool.Wait(); err != nil {
            log.Fatal("Error running pool: ", err)
        }
    }()

    MainLoop(pool)
}

func MainLoop(pool *jobs_lib.Pool) {
    for true {
        pool.Wait()
    }
}

