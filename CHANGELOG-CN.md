# 更新日志

本项目所有重要变更都将记录在此文件中。

## [1.0.0] - 2025-05-26

### 功能特性

- 泛型支持，适用于任意数据类型
- 69 个内置验证规则（字符串、数值、时间、文件、网络、正则等）
- `Validate` - 快速验证，遇错即停
- `ValidateAll` - 收集所有验证错误
- `ValidateStruct` - 结构体字段验证
- 链式 API，支持自定义错误消息（`Errf`）
- 条件验证（`And`、`Or`）
- 依赖验证（`Dependency`、`MutualExclude`）
- 使用 `Field` 辅助函数进行结构体字段验证

### 字符串规则

- `StartWith` / `EndWith` - 前后缀验证
- `ChineseOnly` / `FullWidthOnly` / `HalfWidthOnly` - 字符类型验证
- `UpperCaseOnly` / `LowerCaseOnly` - 大小写验证
- `SpecialChars` / `Contains` / `NotContains` - 子字符串验证

### 数值规则

- `Min` / `Max` / `Between` - 范围验证
- `Positive` / `Negative` / `Even` / `Odd` - 数值属性验证
- `Precision` / `DivisibleBy` / `MultipleOf` - 精度验证

### 时间规则

- `Before` / `After` / `TimeBetween` - 时间范围验证
- `DateFormat` / `TimeFormat` / `DateTimeFormat` - 格式验证
- `Weekend` / `Workday` / `Holiday` - 日期类型验证

### 网络规则

- `IP` / `IPv4` / `IPv6` - IP 地址验证
- `URL` / `Domain` / `Port` - 网络地址验证
- `MACAddress` / `SubnetMask` - 网络配置验证

### 文件规则

- `FileSize` / `FileType` / `FileExtension` / `FileMimeType` - 文件属性验证

### 正则规则

- `IsEmail` / `IsPhone` / `IsIDCard` - 常用格式验证
- `IsBankCard` / `IsPassport` / `IsTaxNumber` - 证件验证
- `IsSocialCredit` / `Regex` - 自定义正则验证

### 其他规则

- `Required` / `NonZero` / `Zero` / `Nil` - 空值验证
- `Len` / `In` / `NotIn` - 长度与集合验证
