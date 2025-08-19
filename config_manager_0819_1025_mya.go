// 代码生成时间: 2025-08-19 10:25:10
package main

import (
    "encoding/json"
    "fmt"
    "os"
    "path/filepath"
    "revel"
)

// ConfigManager 结构体，用于管理配置文件
type ConfigManager struct {
    configPath string
    configs    map[string]interface{}
}

// NewConfigManager 创建一个新的 ConfigManager 实例
func NewConfigManager(path string) *ConfigManager {
    return &ConfigManager{
        configPath: path,
        configs:    make(map[string]interface{}),
    }
}

// LoadConfig 从指定路径加载配置文件
func (cm *ConfigManager) LoadConfig() error {
    dir, err := os.Open(cm.configPath)
    if err != nil {
        return err
    }
    defer dir.Close()

    filenames, err := dir.Readdirnames(-1)
    if err != nil {
        return err
    }

    for _, filename := range filenames {
        file, err := os.Open(filepath.Join(cm.configPath, filename))
        if err != nil {
            continue // 跳过无法打开的文件
        }
        defer file.Close()

        var config interface{}
        if err := json.NewDecoder(file).Decode(&config); err != nil {
            fmt.Printf("Error decoding config file %s: %s
", filename, err)
            continue // 跳过无法解码的文件
        }

        cm.configs[filename] = config
    }

    return nil
}

// GetConfig 获取指定名称的配置
func (cm *ConfigManager) GetConfig(name string) (interface{}, error) {
    config, ok := cm.configs[name]
    if !ok {
        return nil, fmt.Errorf("config '%s' not found", name)
    }
    return config, nil
}

func main() {
    // Revel 初始化
    revel.OnAppStart(func() {
        // 创建配置管理器实例
        configManager := NewConfigManager("./configs")
        // 加载配置文件
        if err := configManager.LoadConfig(); err != nil {
            revel.ERROR.Printf("Failed to load configs: %v", err)
            return
        }
    })

    // 启动 Revel 应用
    revel.Run(
        revel.RunMode("dev"),
        // 其他配置...
    )
}