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

**说明：如果要获取下一页，参数pre=12**

```
{
    "code": 0,
    "data": {
        "list": [
            {
                "id": 21,
                "created_at": "2021-04-14T15:03:26.274+08:00",
                "content": "hello12",
                "user_id": 1,
                "activity_id": 1
            },
            {
                "id": 20,
                "created_at": "2021-04-14T15:03:25.798+08:00",
                "content": "hello12",
                "user_id": 1,
                "activity_id": 1
            },
            {
                "id": 19,
                "created_at": "2021-04-14T15:03:25.388+08:00",
                "content": "hello12",
                "user_id": 1,
                "activity_id": 1
            },
            {
                "id": 18,
                "created_at": "2021-04-14T15:03:24.917+08:00",
                "content": "hello12",
                "user_id": 1,
                "activity_id": 1
            },
            {
                "id": 17,
                "created_at": "2021-04-14T15:03:24.432+08:00",
                "content": "hello12",
                "user_id": 1,
                "activity_id": 1
            },
            {
                "id": 16,
                "created_at": "2021-04-14T15:03:23.938+08:00",
                "content": "hello12",
                "user_id": 1,
                "activity_id": 1
            },
            {
                "id": 15,
                "created_at": "2021-04-14T15:03:23.457+08:00",
                "content": "hello12",
                "user_id": 1,
                "activity_id": 1
            },
            {
                "id": 14,
                "created_at": "2021-04-14T15:03:22.978+08:00",
                "content": "hello12",
                "user_id": 1,
                "activity_id": 1
            },
            {
                "id": 13,
                "created_at": "2021-04-14T15:03:22.564+08:00",
                "content": "hello12",
                "user_id": 1,
                "activity_id": 1
            },
            {
                "id": 12,
                "created_at": "2021-04-14T15:03:19.689+08:00",
                "content": "hello11",
                "user_id": 1,
                "activity_id": 1
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

**说明：如果要获取下一页，参数pre=12**

```
{
    "code": 0,
    "data": {
        "list": [
            {
                "id": 21,
                "created_at": "2021-04-14T15:03:26.274+08:00",
                "content": "hello12",
                "user_id": 1,
                "activity_id": 1
            },
            {
                "id": 20,
                "created_at": "2021-04-14T15:03:25.798+08:00",
                "content": "hello12",
                "user_id": 1,
                "activity_id": 1
            },
            {
                "id": 19,
                "created_at": "2021-04-14T15:03:25.388+08:00",
                "content": "hello12",
                "user_id": 1,
                "activity_id": 1
            },
            {
                "id": 18,
                "created_at": "2021-04-14T15:03:24.917+08:00",
                "content": "hello12",
                "user_id": 1,
                "activity_id": 1
            },
            {
                "id": 17,
                "created_at": "2021-04-14T15:03:24.432+08:00",
                "content": "hello12",
                "user_id": 1,
                "activity_id": 1
            },
            {
                "id": 16,
                "created_at": "2021-04-14T15:03:23.938+08:00",
                "content": "hello12",
                "user_id": 1,
                "activity_id": 1
            },
            {
                "id": 15,
                "created_at": "2021-04-14T15:03:23.457+08:00",
                "content": "hello12",
                "user_id": 1,
                "activity_id": 1
            },
            {
                "id": 14,
                "created_at": "2021-04-14T15:03:22.978+08:00",
                "content": "hello12",
                "user_id": 1,
                "activity_id": 1
            },
            {
                "id": 13,
                "created_at": "2021-04-14T15:03:22.564+08:00",
                "content": "hello12",
                "user_id": 1,
                "activity_id": 1
            },
            {
                "id": 12,
                "created_at": "2021-04-14T15:03:19.689+08:00",
                "content": "hello11",
                "user_id": 1,
                "activity_id": 1
            }
        ]
    },
    "message": "成功"
}
```