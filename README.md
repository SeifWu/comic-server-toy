## Base

```
- app
  - assets
  - controllers
    - api/v1/manager
    - api/v1/pubilc
  - middleware
  - models
  - utils
- config
  - initializers
  - routes
- global

```

## Features

- [x] 发送邮件
- [x] 邮件验证码 Redis 缓存
- [x] 邮箱方式注册

## TODO

- [ ] JWT
- [ ] validator v10

## Response

### Success

```
  {
    "success": true,
    "data": {},
    "meta": {},
  }
```

### fail

```
  {
    "success": false,
    "code": 40001,
    "msg": "error message",
  }
```
