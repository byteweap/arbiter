<p align="center">
  <img src="arbiter.png" alt="Arbiter Logo" width="200"/>
</p>

# Arbiter

一个强大而灵活的 Go 数据验证框架。

[English](README.md) | 简体中文

## 概述

Arbiter 是一个 Go 的全面的数据验证框架，提供丰富的验证规则和灵活的验证机制。它支持基本数据类型、字符串、数字、时间、文件等的验证，并内置支持结构体字段验证。

## 特性

- 支持任意数据类型的泛型验证
- 丰富的内置验证规则
- 支持自定义验证规则
- 结构体字段验证
- 链式 API
- 自定义错误消息
- 条件验证
- 依赖验证

## 安装

```bash
go get github.com/byteweap/arbiter
```

## 快速开始

```go
package main

import (
    "fmt"
    "github.com/byteweap/arbiter"
    "github.com/byteweap/arbiter/rule"
)

type Person struct {
    Name  string
    Age   int
    Email string
}

func main() {
    person := &Person{
        Name:  "张三",
        Age:   30,
        Email: "zhangsan@example.com",
    }

    err := arbiter.ValidateStruct(person, "Person 不能为空",
        rule.Field(&person.Name,
            rule.Length[string](2, 50).Errf("姓名长度必须在2-50之间"),
            rule.Required[string]().Errf("姓名不能为空"),
        ),
        rule.Field(&person.Age,
            rule.Min(0),
            rule.Max(120),
        ),
        rule.Field(&person.Email,
            rule.Required[string]().Errf("邮箱格式不正确")
            rule.IsEmail().Errf("邮箱格式错误"),
        ),
    )

    if err != nil {
        fmt.Printf("验证错误: %v\n", err)
    }
}
```

## 核心组件

### 1. 验证器 (Arbiter)

主要的验证函数：

```go
// Validate 对单个值应用多个规则
err := Validate("hello",
    rule.Length[string](3, 10).Errf("字符串格式不正确"),
    // ...
)

// ValidateWithErrs 收集所有验证错误
err := ValidateWithErrs("hello",
    rule.Required[string]().Errf("不能为空")
    rule.Length(3, 10).Errf("格式不正确"),
)

// ValidateStruct 验证结构体及其字段
err := ValidateStruct(person, "Person 不能为空",
    rule.Field(&person.Name, ...),
    rule.Field(&person.Age, ...),
)
```

### 2. 字段验证器

用于验证结构体字段：

```go
// 创建字段验证规则
nameRule := rule.Field(&person.Name,
    rule.Length(2, 50),
    rule.String().Errf("姓名不能为空"),
)
```

### 3. 验证规则

#### 字符串规则
- `StartWith`: 验证字符串前缀
- `EndWith`: 验证字符串后缀
- `OnlyChinese`: 验证中文字符
- `OnlyFullWidth`: 验证全角字符
- `OnlyHalfWidth`: 验证半角字符
- `OnlyUpperCase`: 验证大写字母
- `OnlyLowerCase`: 验证小写字母
- `SpecialChars`: 验证特殊字符
- `Contains`: 验证子串存在
- `NotContains`: 验证子串不存在
- ...
  
#### 数字规则
- `Min`: 最小值
- `Max`: 最大值
- `Between`: 范围验证
- `Positive`: 正数验证
- `Negative`: 负数验证
- `Even`: 偶数验证
- `Odd`: 奇数验证
- `Precision`: 小数精度验证
- ...

#### 时间规则
- `Before`: 早于指定时间
- `After`: 晚于指定时间
- `Between`: 时间范围验证
- ...

#### 文件规则
- `Size`: 文件大小验证
- `Extension`: 文件扩展名验证
- ...

#### 网络规则
- `IP`: IP 地址验证
- `URL`: URL 验证
- `Email`: 邮箱验证
- ...

#### 正则表达式 (采用预编译提高性能)
- `IsEmail`: 邮箱验证
- `IsPhone`: 手机号验证
- `Regex`: 自定义正则表达式验证
- ...

## 最佳实践

### 1. 错误处理

```go
// 收集第一个错误
err := Validate(value,
    rule1,
    rule2,
    rule3,
)
if err != nil {
    // 处理第一个错误
}
```

```go
// 收集所有验证错误
errs := ValidateWithErrs(value,
    rule1,
    rule2,
    rule3,
)
if len(errs) > 0 {
    // 处理多个错误
}
```

### 2. 结构体验证

```go
type User struct {
    Username string
    Password string
    Email    string
}

func (u *User) Validate() error {
   return ValidateStruct(user, "User 不能为空",
        rule.Field(&user.Username,
            rule.Length(3, 20).Errf("用户名不符合规则"),
            // ...
        ),
        rule.Field(&user.Password,
            rule.Length(8, 50),
            rule.SpecialChars(true).Errf("密码必须包含特殊字符"),
            // ...
        ),
        rule.Field(&user.Email,
            rule.IsEmail().Errf("邮箱格式不正确"),
            // ...
        ),
        // ...
    )
}


```

### 3. 规则组合

```go
// 使用 AND 组合规则
rule := rule.And(
    rule.Length(3, 10).Errf("字符串格式不正确"),
)

// 使用 OR 组合规则
rule := rule.Or(
    rule.IsEmail(),
    rule.URL(),
)
```

### 4. 自定义规则

```go
// 创建自定义规则
type CustomRule struct {
    err error
}

func (r *CustomRule) Validate(value string) error {
    if value == "" {
        return r.err
    }
    return nil
}

// 使用自定义规则
rule := &CustomRule{err: errors.New("自定义错误")}
err := Validate("", rule)
```

## 贡献指南

1. Fork 本仓库
2. 创建特性分支 (`git checkout -b feature/amazing-feature`)
3. 提交更改 (`git commit -m '添加一些很棒的特性'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 提交 Pull Request

## 测试

运行测试套件：

```bash
# 运行所有测试
go test -v ./...

# 运行带覆盖率的测试
go test -v -cover ./...
go test -v -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## 许可证

[待添加许可证详情]

## 版本历史

[待添加版本历史]

## 为什么选择 Arbiter？

Arbiter 意为"仲裁者"或"判断者"，旨在成为您代码的验证权威。我们选择这个名字是因为：

1. **全面的验证**：就像法官评估证据一样，Arbiter 根据定义的规则彻底验证您的数据。
2. **类型安全**：基于 Go 的类型系统构建，提供编译时类型检查和泛型支持。
3. **灵活的规则**：支持内置和自定义验证规则，让您能够精确定义什么是有效数据。
4. **高性能**：优化的验证性能，最小化内存分配。
5. **开发体验**：直观的 API 设计，支持链式规则和清晰的错误消息。

### 相比基于 Tag 的验证库的优势

虽然基于 tag 的验证库很流行，但 Arbiter 采用了不同的方法，具有以下关键优势：

1. **类型安全和 IDE 支持**
   - 基于 tag：验证规则定义在字符串标签中，无法进行类型检查，缺乏 IDE 支持
   - Arbiter：规则使用 Go 代码定义，提供完整的类型安全和 IDE 功能（自动完成、重构等）

2. **运行时性能**
   - 基于 tag：需要在运行时解析标签并使用反射进行验证
   - Arbiter：直接函数调用，最小化反射使用，性能更好

3. **灵活性和可维护性**
   - 基于 tag：复杂的验证规则在标签中难以阅读和维护
   - Arbiter：规则就是普通的 Go 代码，便于组织和重用

4. **调试和测试**
   - 基于 tag：验证规则的错误只能在运行时发现
   - Arbiter：验证逻辑可以像普通代码一样进行单元测试和调试

5. **自定义规则**
   - 基于 tag：自定义验证器通常需要注册和反射
   - Arbiter：自定义规则只是 Go 接口，实现和使用都很简单

6. **条件验证**
   - 基于 tag：在标签中难以表达复杂的条件
   - Arbiter：可以使用 Go 的全部功能进行条件逻辑

7. **使用场景**
   不管是web框架(如 Gin、Echo、Iris 等), 还是rpc框架(如 grpc、go-kratos、go-zero等), 总之 Arbiter 可使用于任何Go项目需要的地方

在构建现代应用程序时，数据验证至关重要。Arbiter 提供了一个健壮、类型安全且可扩展的解决方案，能够随您的应用程序一起成长。
