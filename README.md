适配于 `github.com/zhongjin616/errors` 错误包的错误码实现。

## Code 设计规范

> **通用** - 所有服务都适用的错误，提高复用性，避免重复造轮子

## 错误描述规范

错误描述包括：对外的错误描述和对内的错误描述两部分。

### 对外的错误描述

- 对外暴露的错误，统一大写开头，结尾不要加`.`
- 对外暴露的错误，要简洁，并能准确说明问题
- 对外暴露的错误说明，应该是 `该怎么做` 而不是 `哪里错了`

### 对内的错误描述

- 告诉用户他们可以做什么，而不是告诉他们不能做什么。
- 当声明一个需求时，用 must 而不是 should。例如，must be greater than 0、must match regex '[a-z]+'。
- 当声明一个格式不对时，用 must not。例如，must not contain。
- 当声明一个动作时用 may not。例如，may not be specified when otherField is empty、only name may be specified。
- 引用文字字符串值时，请在单引号中指示文字。例如，ust not contain '..'。
- 当引用另一个字段名称时，请在反引号中指定该名称。例如，must be greater than request。
- 指定不等时，请使用单词而不是符号。例如，must be less than 256、must be greater than or equal to 0 (不要用 larger than、bigger than、more than、higher than)。
- 指定数字范围时，请尽可能使用包含范围。
- 建议 Go 1.13 以上，error 生成方式为 fmt.Errorf("module xxx: %w", err)。
- 错误描述用小写字母开头，结尾不要加标点符号。

> 错误信息是直接暴露给用户的，不能包含敏感信息

## 错误记录规范

在错误产生的最原始位置调用日志，打印错误信息，其它位置直接返回。

