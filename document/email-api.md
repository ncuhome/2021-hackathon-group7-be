# 2021-hackathon-backend

# 邮箱API文档s

host: https://nspyf.top

port：11000

## 方法描述：发送绑定邮箱的验证码

URL地址：/auth/email/binding-key

请求方法：POST

请求体：

| 字段 | 类型 | 说明 |
| ---  | ---  | ---  |
|  email  |  string  |  需绑定的邮箱,required |

请求体示例

```
{
    "email":"316851756@qq.com"
}
```

响应体：

| 字段 | 类型 | 说明 |
| ---  | ---  | ---  |
|  code  |  int  |  状态码  |
|  message  |  string  |  状态码描述  |

响应体示例

```
{
    "code": 0,
    "message": "成功"
}
```

## 方法描述：绑定邮箱

URL地址：/auth/email/binding

请求方法：POST

请求体：

| 字段 | 类型 | 说明 |
| ---  | ---  | ---  |
|  email  |  string  |  需绑定的邮箱,required |
|  key  |  string  |  验证码,required |

请求体示例

```
{
    "email":"316851756@qq.com",
    "key":"304793"
}
```

响应体：

| 字段 | 类型 | 说明 |
| ---  | ---  | ---  |
|  code  |  int  |  状态码  |
|  message  |  string  |  状态码描述  |

响应体示例

```
{
    "code": 0,
    "message": "成功"
}
```

## 方法描述：解除邮箱绑定

URL地址：/auth/email/binding

请求方法：DELETE

请求体：无

响应体：

| 字段 | 类型 | 说明 |
| ---  | ---  | ---  |
|  code  |  int  |  状态码  |
|  message  |  string  |  状态码描述  |

响应体示例

```
{
    "code": 0,
    "message": "成功"
}
```

## 方法描述：发送通过邮箱重设密码的邮件

URL地址：/email/password-key

请求方法：POST

请求体示例

```
{
    "email":"316851756@qq.com"
}
```

响应体：

| 字段 | 类型 | 说明 |
| ---  | ---  | ---  |
|  code  |  int  |  状态码  |
|  message  |  string  |  状态码描述  |

响应体示例

```
{
    "code": 0,
    "message": "成功"
}
```

## 方法描述：通过邮箱重设密码

URL地址：/email/password

请求方法：POST

请求体：

| 字段 | 类型 | 说明 |
| ---  | ---  | ---  |
|  key  |  string  |  邮件的重设密码链接地址的key参数 |
|  new_password  |  string  |  新密码,required,min=8,max=32 |

请求体示例

```
{
    "key":"dc6cb5de17c2d6d9185611fbcc07bbb2",
    "new_password":"12345678c"
}
```

响应体：

| 字段 | 类型 | 说明 |
| ---  | ---  | ---  |
|  code  |  int  |  状态码  |
|  message  |  string  |  状态码描述  |

响应体示例

```
{
    "code": 0,
    "message": "成功"
}
```

## 方法描述：获取绑定的邮箱

URL地址：/auth/email

请求方法：GET

响应体：

| 字段 | 类型 | 说明 |
| ---  | ---  | ---  |
|  code  |  int  |  状态码  |
|  message  |  string  |  状态码描述  |
|  data  |  JSON  |  具体看示例  |

响应体示例

```
{
    "code": 0,
    "message": "成功"
    "data": {
        "email": "316851756@qq.com"
    }
}
```