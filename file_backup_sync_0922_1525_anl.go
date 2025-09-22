// 代码生成时间: 2025-09-22 15:25:30
package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "strings"
)

// BackupSyncTool 结构体，用于文件备份和同步
type BackupSyncTool struct {
    sourceDir  string
    backupDir  string
    syncDir    string
    verbose    bool
}

// NewBackupSyncTool 创建新的 BackupSyncTool 实例
func NewBackupSyncTool(sourceDir, backupDir, syncDir string, verbose bool) *BackupSyncTool {
    return &BackupSyncTool{
        sourceDir:  sourceDir,
        backupDir:  backupDir,
        syncDir:    syncDir,
        verbose:    verbose,
    }
}

// Backup 备份源目录到备份目录
func (b *BackupSyncTool) Backup() error {
    return filepath.WalkDir(b.sourceDir, func(path string, d os.DirEntry, err error) error {
        if err != nil {
            return err
        }
        relPath, err := filepath.Rel(b.sourceDir, path)
        if err != nil {
            return err
        }
        destPath := filepath.Join(b.backupDir, relPath)
        if d.IsDir() {
            return os.MkdirAll(destPath, 0755)
        } else {
            content, err := ioutil.ReadFile(path)
            if err != nil {
                return err
            }
            if err := ioutil.WriteFile(destPath, content, 0644); err != nil {
                return err
            }
        }
        return nil
    })
}

// Sync 同步备份目录到同步目录
func (b *BackupSyncTool) Sync() error {
    return filepath.WalkDir(b.backupDir, func(path string, d os.DirEntry, err error) error {
        if err != nil {
            return err
        }
        relPath, err := filepath.Rel(b.backupDir, path)
        if err != nil {
            return err
        }
        destPath := filepath.Join(b.syncDir, relPath)
        if d.IsDir() {
            return os.MkdirAll(destPath, 0755)
        } else {
            content, err := ioutil.ReadFile(path)
            if err != nil {
                return err
            }
            if err := ioutil.WriteFile(destPath, content, 0644); err != nil {
                return err
            }
        }
        return nil
    })
}

// Run 运行备份和同步工具
func (b *BackupSyncTool) Run() error {
    if err := b.Backup(); err != nil {
        return fmt.Errorf("backup failed: %w", err)
    }
    if err := b.Sync(); err != nil {
        return fmt.Errorf("sync failed: %w", err)"
    }
    return nil
}

func main() {
    // 示例用法
    sourceDir := "./source"
    backupDir := "./backup"
    syncDir := "./sync"
    verbose := false
    
    tool := NewBackupSyncTool(sourceDir, backupDir, syncDir, verbose)
    if err := tool.Run(); err != nil {
        log.Fatalf("Error running backup and sync tool: %s", err)
    }
    fmt.Println("Backup and sync completed successfully")
}