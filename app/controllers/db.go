package controllers

import "database/sql"
import "github.com/revel/revel"
import "github.com/jinzhu/gorm"
import "qixalite.com/Ranndom/ldap-portal/app"

type TransactionalController struct {
    *revel.Controller
    Txn *gorm.DB
}

func (c *TransactionalController) Begin() revel.Result {
    txn := app.DB.Begin()
    if txn.Error != nil {
        panic(txn.Error)
    }

    c.Txn = txn
    return nil
}

func (c *TransactionalController) Commit() revel.Result {
    if c.Txn == nil {
        return nil
    }

    c.Txn.Commit()
    if err := c.Txn.Error; err != nil && err != sql.ErrTxDone {
        panic(err)
    }

    c.Txn = nil
    return nil
}

func (c *TransactionalController) Rollback() revel.Result {
    if c.Txn == nil {
        return nil
    }

    c.Txn.Rollback()
    if err := c.Txn.Error; err != nil && err != sql.ErrTxDone {
        panic(err)
    }

    c.Txn = nil
    return nil
}

