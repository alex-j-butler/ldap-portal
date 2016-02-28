package tests

import "github.com/revel/revel/testing"

type AppTest struct {
    testing.TestSuite
}

func (t *AppTest) Before() {
    println("Set up")
}

func (t *AppTest) TestIndexPageLoads() {
    t.Get("/")
    t.AssertOk()
    t.AssertContentType("text/html; charset=utf-8")
}

func (t *AppTest) TestLoginPageLoads() {
    t.Get("/login")
    t.AssertOk()
    t.AssertContentType("text/html; charset=utf-8")
}

func (t *AppTest) After() {
    println("Tear down")
}
