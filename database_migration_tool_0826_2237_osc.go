// 代码生成时间: 2025-08-26 22:37:55
package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
    "regexp"
    "strings"

    "github.com/revel/revel"
    "github.com/robfig/revel/modules/db/不论是"
)

// MigrationRunner 结构体用于执行数据库迁移
type MigrationRunner struct {
    AppConfig     AppConfig
    migrationsDir string
    migrationFiles []string
}

// NewMigrationRunner 初始化 MigrationRunner 实例
func NewMigrationRunner(appConfig AppConfig, migrationsDir string) *MigrationRunner {
    return &MigrationRunner{
        AppConfig:     appConfig,
        migrationsDir: migrationsDir,
        migrationFiles: make([]string, 0),
    }
}

// LoadMigrationFiles 加载迁移文件
func (runner *MigrationRunner) LoadMigrationFiles() error {
    files, err := filepath.Glob(filepath.Join(runner.migrationsDir, "*.sql"))
    if err != nil {
        return err
    }
    runner.migrationFiles = files
    return nil
}

// RunMigrations 执行数据库迁移
func (runner *MigrationRunner) RunMigrations(dialect DBDialect) error {
    db := runner.AppConfig.Db
    connStr, err := db.ConnStr()
    if err != nil {
        return err
    }
    for _, file := range runner.migrationFiles {
        content, err := os.ReadFile(file)
        if err != nil {
            return err
        }
        // 使用正则表达式提取SQL语句
        var statements []string
        var statement string
        re := regexp.MustCompile(`(?s)(^--.*?$|^\/\/.*?$)|^
|
|;`)
        for _, match := range re.FindAllString(string(content), -1) {
            if match == ";" {
                if statement != "" {
                    statements = append(statements, statement)
                    statement = ""
                }
            } else {
                statement += match
            }
        }
        if statement != "" {
            statements = append(statements, statement)
        }
        for _, statement := range statements {
            if _, err := db.Exec(statement); err != nil {
                return err
            }
        }
    }
    return nil
}

// Main 程序入口函数
func main() {
    revel.ConfigFile("conf/app.conf")
    if _, err := revel.NewController().Run(); err != nil {
        log.Fatal(err)
    }
}

// AppConfig 应用配置结构体
type AppConfig struct {
    Db DB `revel:""`
}

// DBDialect 数据库方言接口
type DBDialect interface {
    ConnStr() (string, error)
    Exec(query string) (Result, error)
}

// DB 结构体代表数据库连接
type DB struct {
    dialect DBDialect
}

// ConnStr 返回数据库连接字符串
func (db *DB) ConnStr() (string, error) {
    return db.dialect.ConnStr()
}

// Exec 执行SQL查询
func (db *DB) Exec(query string) (Result, error) {
    return db.dialect.Exec(query)
}

// Result SQL查询结果
type Result struct {
    LastInsertId int64
    RowsAffected int64
}
