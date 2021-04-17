# 2021-hackathon-backend

# 用户信息API文档

host: https://nspyf.top

port：11000

## 方法描述：修改用户信息

URL地址：/auth/user-info

请求方法：PUT

请求体：

| 字段 | 类型 | 说明 |
| ---  | ---  | ---  |
|  nickname  |  string  |  昵称/组织名 |
|  avatar  |  string  |  头像链接地址  |
|  digest  |  string  |  摘要/简介  |

**注意，未更改的值要填写原来的值提交**

请求体示例

```
{
	"nickname":"汤姆",
	"avatar": "http://...",
	"digest":"一只猫"
}
```

响应体：


| 字段 | 类型 | 说明 |
| ---  | ---  | ---  |
|  code  |  int  |  状态码  |
|  message  |  string  |  状态码描述  |
|  data  |  JSON  |  更新的数据  |

响应体示例

```
{
    "code": 0,
    "data": {
        "avatar": "http://...",
        "digest": "一只猫",
        "nickname": "汤姆"
    },
    "message": "成功"
}
```

## 方法描述：获取用户信息

URL地址：/user-info

请求方法：GET

请求参数

| 字段 | 说明 |
| ---  | ---  |
|  id  |  用户id |

请求示例

```
/user-info?id=3
```

响应体：


| 字段 | 类型 | 说明 |
| ---  | ---  | ---  |
|  code  |  int  |  状态码  |
|  message  |  string  |  状态码描述  |
|  data  |  JSON  |  用户信息,verification字段为空字符串则未认证  |

响应体示例

```
{
    "code": 0,
    "data": {
        "avatar": "http://...",
        "digest": "一只猫",
        "nickname": "汤姆"，
        "verification": ""
    },
    "message": "成功"
}
```

## 方法描述：修改用户认证信息

**说明：请求头携带的令牌必须是管理员用户的令牌**

URL地址：/auth/verification

请求方法：PUT

请求体：

| 字段 | 类型 | 说明 |
| ---  | ---  | ---  |
|  id  |  int  |  操作的用户 |
|  verification  |  string  |  认证信息(为空则取消认证)  |

请求体示例

```
{
    "id":3,
    "verification":"v"
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

## 方法描述：分页获取认证用户

**说明：每页十条**

URL地址：/verification/user

请求方法：GET

请求参数：

| 字段 | 说明 |
| ---  | ---  |
|  pre |  上一页的最后一个（最小的）用户id，如果是获取第一页则填0  |

请求示例

```
/verification/user?pre=0
```

响应体：


| 字段 | 类型 | 说明 |
| ---  | ---  | ---  |
|  code  |  int  |  状态码  |
|  message  |  string  |  状态码描述  |
|  data  |  JSON  |  具体字段参见示例  |

响应体示例

```
{
    "code": 0,
    "data": {
        "list": [
            {
                "user_id": 3,
                "nickname": "你好",
                "avatar": "123",
                "Digest": "测试AAAAA",
                "verification": "v"
            },
            {
                "user_id": 4,
                "verification": "v"
            }
        ]
    },
    "message": "成功"
}
```