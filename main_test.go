package main

import (
	"gorm.io/gorm"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type A struct {
	gorm.Model

	BID uint `gorm:"not null"`
	B   B
}

type B struct {
	gorm.Model

	As []A
}

func TestGORM(t *testing.T) {
	m := DB.Migrator()

	m.DropTable("as")
	m.DropTable("bs")

	m.CreateTable(&A{})
}
