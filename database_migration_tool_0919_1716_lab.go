// 代码生成时间: 2025-09-19 17:16:45
package main

import (
    "fmt"
    "os"
    "path/filepath"

    "github.com/revel/revel"
    "github.com/robfig/revel/modules/db/sqlite"
)

// MigrationTool 结构体用于数据库迁移工具
type MigrationTool struct {
    // 包含 Revel控制器
    *revel.Controller
}

// NewMigrationTool 初始化数据库迁移工具
func NewMigrationTool() *MigrationTool {
    return &MigrationTool{
        Controller: &revel.Controller{
            App: revel.MainApp,
        },
    }
}

// RunMigrations 运行数据库迁移
func (tool *MigrationTool) RunMigrations() error {
    // 检测迁移目录是否存在
    migrationsDir := filepath.Join(revel.BasePath, "migrations")
    if _, err := os.Stat(migrationsDir); os.IsNotExist(err) {
        return fmt.Errorf("migrations directory does not exist: %s", migrationsDir)
    }

    // 连接数据库
    db, err := sqlite.Open("path/to/database.db?_foreign_keys=true")
    if err != nil {
        return fmt.Errorf("failed to connect to the database: %s", err)
    }
    defer db.Close()

    // 初始化迁移
    if err := db.InitDB(); err != nil {
        return fmt.Errorf("failed to initialize the database: %s", err)
    }

    // 获取迁移文件列表
    files, err := os.ReadDir(migrationsDir)
    if err != nil {
        return fmt.Errorf("failed to read migrations directory: %s", err)
    }

    for _, file := range files {
        if !file.IsDir() && filepath.Ext(file.Name()) == ".sql" {
            // 运行迁移文件
            migrationPath := filepath.Join(migrationsDir, file.Name())
            if err := db.RunMigration(migrationPath); err != nil {
                return fmt.Errorf("failed to run migration: %s, error: %s", migrationPath, err)
            }
        }
    }

    return nil
}

func main() {
    // 初始化Revel框架
    revel.Init()
    defer revel.Close()

    // 创建数据库迁移工具
    migrationTool := NewMigrationTool()

    // 运行数据库迁移
    if err := migrationTool.RunMigrations(); err != nil {
        revel.ERROR.Fatalln("Database migration failed: ", err)
    } else {
        fmt.Println("Database migration successful")
    }
}