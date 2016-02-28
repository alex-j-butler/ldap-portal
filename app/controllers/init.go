package controllers

import (
    "github.com/revel/revel"
)

func init() {
    revel.InterceptMethod((*TransactionalController).Begin, revel.BEFORE)
    revel.InterceptMethod((*TransactionalController).Commit, revel.AFTER)
    revel.InterceptMethod((*TransactionalController).Rollback, revel.FINALLY)
}
