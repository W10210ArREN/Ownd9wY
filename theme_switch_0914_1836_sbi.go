// 代码生成时间: 2025-09-14 18:36:39
package controllers

import (
    "encoding/json"
    "github.com/revel/revel"
)

// ThemeController 控制主题切换的逻辑
type ThemeController struct {
    revel.Controller
}

// SwitchAction 实现主题切换功能
func (c ThemeController) SwitchAction(theme string) revel.Result {
    // 检查主题是否有效
    validThemes := []string{"light", "dark"}
    if !contains(validThemes, theme) {
        return c.RenderJSON(revel.Result{
            "error": "Invalid theme",
        })
    }

    // 存储主题到会话
    c.Session["theme"] = theme

    // 返回成功响应
    return c.RenderJSON(revel.Result{
        "message": "Theme switched successfully",
        "theme": theme,
    })
}

// contains 检查字符串切片中是否包含指定的元素
func contains(s []string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}
