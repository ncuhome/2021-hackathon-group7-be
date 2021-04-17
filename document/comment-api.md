# 2021-hackathon-backend

# 评论API文档

host: https://nspyf.top

port：11000

## 方法描述：用户对活动评论

URL地址：/auth/comment

请求方法：POST

请求体：

| 字段 | 类型 | 说明 |
| ---  | ---  | ---  |
|  content  |  string  |  内容 |
|  activity_id  |  int  |  活动id  |

请求体示例

```
{
    "content":"hello",
    "activity_id":1
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

## 方法描述：分页获取活动下的评论

**说明：每页十条**

URL地址：/activity/comment

请求方法：GET

请求参数：

| 字段 | 说明 |
| ---  | ---  |
|  id  |  活动id |
|  pre |  上一页的最后一个（最小的）评论id，如果是获取第一页则填0  |

请求示例

```
/activity/comment?id=1&pre=0
```

响应体：


| 字段 | 类型 | 说明 |
| ---  | ---  | ---  |
|  code  |  int  |  状态码  |
|  message  |  string  |  状态码描述  |
|  data  |  JSON  |  评论列表，具体字段参见示例  |

响应体示例

```
{
    "code": 0,
    "data": {
        "list": [
            {
                "comment": {
                    "id": 3,
                    "created_at": "2021-04-14T15:02:52.433+08:00",
                    "content": "hello2",
                    "user_id": 1,
                    "activity_id": 1
                },
                "user": {
                    "user_id": 3,
                    "nickname": "你好",
                    "avatar": "https://",
                    "verification": "v"
                }
            },
            {
                "comment": {
                    "id": 2,
                    "created_at": "2021-04-14T15:02:49.091+08:00",
                    "content": "hello1",
                    "user_id": 1,
                    "activity_id": 1
                },
                "user": {
                    "user_id": 3,
                    "nickname": "你好",
                    "avatar": "https://",
                    "verification": "v"
                }
            },
            {
                "comment": {
                    "id": 1,
                    "created_at": "2021-04-14T13:43:02.041+08:00",
                    "content": "hello",
                    "user_id": 2,
                    "activity_id": 1
                },
                "user": {
                    "user_id": 4
                }
            }
        ]
    },
    "message": "成功"
}
```

## 方法描述：分页获取用户的评论

**说明：每页十条**

URL地址：/user/comment

请求方法：GET

请求参数：

| 字段 | 说明 |
| ---  | ---  |
|  id  |  用户id |
|  pre |  上一页的最后一个（最小的）评论id，如果是获取第一页则填0  |

请求示例

```
/user/comment?id=1&pre=0
```

响应体：


| 字段 | 类型 | 说明 |
| ---  | ---  | ---  |
|  code  |  int  |  状态码  |
|  message  |  string  |  状态码描述  |
|  data  |  JSON  |  评论列表，具体字段参见示例  |

响应体示例

```
{
    "code": 0,
    "data": {
        "list": [
            {
                "comment": {
                    "id": 3,
                    "created_at": "2021-04-14T15:02:52.433+08:00",
                    "content": "hello2",
                    "user_id": 1,
                    "activity_id": 1
                },
                "user": {
                    "user_id": 3,
                    "nickname": "你好",
                    "avatar": "https://",
                    "verification": "v"
                }
            },
            {
                "comment": {
                    "id": 2,
                    "created_at": "2021-04-14T15:02:49.091+08:00",
                    "content": "hello1",
                    "user_id": 1,
                    "activity_id": 1
                },
                "user": {
                    "user_id": 3,
                    "nickname": "你好",
                    "avatar": "https://",
                    "verification": "v"
                }
            }
        ]
    },
    "message": "成功"
}
```