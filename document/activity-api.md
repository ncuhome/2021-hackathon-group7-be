# 2021-hackathon-backend

# 活动API文档

## 方法描述：用户创建新活动

URL地址：/auth/activity

请求方法：POST

请求体：

| 字段 | 类型 | 说明 |
| ---  | ---  | ---  |
|  title  |  string  |  活动题目 |
|  content  |  string  |  活动内容  |
|  start_time  |  int（时间戳）  |  开始时间 |
|  end_time  |  int（时间戳）  |  结束时间 |
|  place  |  string  |  活动举办地 |

请求体示例

```
{
    "title":"activity",
    "content":"this is an activity",
    "start_time":1641505804,
    "end_time":1641506804,
    "place":"classroom1",
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

## 方法描述：获取全部活动


URL地址：/activities/all

请求方法：GET

请求参数：

| 字段   | 说明 |
| ---    | ---  |
|  pre  |  上一页的最后一个（最小的）评论id，第一页则填0  |

请求示例

```
/activities/all？pre=1
```

响应体：


| 字段 | 类型 | 说明 |
| ---  | ---  | ---  |
|  code  |  int  |  状态码  |
|  message  |  string  |  状态码描述  |
|  data  |  JSON  |  具体看样例  |

响应体示例



```
{
    "code": 0,
    "data": [
        {
            "ID": 6,
            "CreatedAt": "2021-04-15T14:42:35.08+08:00",
            "UpdatedAt": "2021-04-15T14:42:35.08+08:00",
            "DeletedAt": null,
            "Title": "test1888me",
            "Content": "sd54",
            "UserId": "35",
            "StartTime": "16400058046",
            "EndTime": "1641505804",
            "Place": "w",
            "Digest": ""
        },
        {
            "ID": 5,
            "CreatedAt": "2021-04-13T23:06:28.466+08:00",
            "UpdatedAt": "2021-04-13T23:06:28.466+08:00",
            "DeletedAt": null,
            "Title": "test1888me",
            "Content": "sd54",
            "UserId": "35",
            "StartTime": "1640005804",
            "EndTime": "1641505804",
            "Place": "w",
            "Digest": ""
        },
        {
            "ID": 7,
            "CreatedAt": "2021-04-15T14:53:20.571+08:00",
            "UpdatedAt": "2021-04-15T14:53:20.571+08:00",
            "DeletedAt": null,
            "Title": "test1888me",
            "Content": "sd54",
            "UserId": "35",
            "StartTime": "1640",
            "EndTime": "1641505804",
            "Place": "w",
            "Digest": ""
        },
        {
            "ID": 8,
            "CreatedAt": "2021-04-15T14:55:46.444+08:00",
            "UpdatedAt": "2021-04-15T14:55:46.444+08:00",
            "DeletedAt": null,
            "Title": "test1888megggg",
            "Content": "sd54",
            "UserId": "35",
            "StartTime": "1640",
            "EndTime": "1641505804",
            "Place": "w",
            "Digest": ""
        },
        {
            "ID": 1,
            "CreatedAt": "2021-04-13T17:47:35.606+08:00",
            "UpdatedAt": "2021-04-13T17:47:35.606+08:00",
            "DeletedAt": null,
            "Title": "test1",
            "Content": "sdfsfdgd165665asdfs",
            "UserId": "35",
            "StartTime": "1213",
            "EndTime": "1111",
            "Place": "this",
            "Digest": ""
        },
        {
            "ID": 2,
            "CreatedAt": "2021-04-13T18:22:16.43+08:00",
            "UpdatedAt": "2021-04-13T18:22:16.43+08:00",
            "DeletedAt": null,
            "Title": "test1888",
            "Content": "sd54",
            "UserId": "35",
            "StartTime": "1213",
            "EndTime": "1111",
            "Place": "this",
            "Digest": ""
        },
        {
            "ID": 3,
            "CreatedAt": "2021-04-13T22:53:40.287+08:00",
            "UpdatedAt": "2021-04-13T22:53:40.287+08:00",
            "DeletedAt": null,
            "Title": "test1888me",
            "Content": "sd54",
            "UserId": "35",
            "StartTime": "1213",
            "EndTime": "1389045004",
            "Place": "w",
            "Digest": ""
        },
        {
            "ID": 4,
            "CreatedAt": "2021-04-13T22:53:52.205+08:00",
            "UpdatedAt": "2021-04-13T22:53:52.205+08:00",
            "DeletedAt": null,
            "Title": "test1888me",
            "Content": "sd54",
            "UserId": "35",
            "StartTime": "1213",
            "EndTime": "1641505804",
            "Place": "w",
            "Digest": ""
        }
    ],
    "message": "成功"
}
```

## 方法描述：获取活动详情


URL地址：/activities/detail

请求方法：GET

请求参数：

| 字段 | 说明 |
| ---  | ---  |
|  id  |  活动id |


请求示例

```
/user/detail?id=1
```

响应体：


| 字段 | 类型 | 说明 |
| ---  | ---  | ---  |
|  code  |  int  |  状态码  |
|  message  |  string  |  状态码描述  |
|  data  |  JSON  |  具体看样例  |

响应体示例


```
{
    "code": 0,
    "data": {
        "Content": "sdfsfdgd165665asdfs",
        "CreatedAt": "2021-04-13T17:47:35.606+08:00",
        "DeletedAt": null,
        "Digest": "",
        "EndTime": "1111",
        "ID": 1,
        "Place": "this",
        "StartTime": "1213",
        "Title": "test1",
        "UpdatedAt": "2021-04-13T17:47:35.606+08:00",
        "UserId": "35"
    },
    "message": "成功"
}
```

## 方法描述：根据活动地点获取活动


URL地址：/activities/place

请求方法：GET

请求参数：

| 字段 | 说明 |
| ---  | ---  |
|  place  |  活动地点 |
|  pre    | 上一页最小id,第一页填0|

请求示例

```
/activity/place?place=1 pre=0
```

响应体：


| 字段 | 类型 | 说明 |
| ---  | ---  | ---  |
|  code  |  int  |  状态码  |
|  message  |  string  |  状态码描述  |
|  data  |  JSON  |  具体看样例  |

响应体示例


```
{
    "code": 0,
    "data": [
        {
            "ID": 6,
            "CreatedAt": "2021-04-15T14:42:35.08+08:00",
            "UpdatedAt": "2021-04-15T14:42:35.08+08:00",
            "DeletedAt": null,
            "Title": "test1888me",
            "Content": "sd54",
            "UserId": "35",
            "StartTime": "16400058046",
            "EndTime": "1641505804",
            "Place": "w",
            "Digest": ""
        },
        {
            "ID": 5,
            "CreatedAt": "2021-04-13T23:06:28.466+08:00",
            "UpdatedAt": "2021-04-13T23:06:28.466+08:00",
            "DeletedAt": null,
            "Title": "test1888me",
            "Content": "sd54",
            "UserId": "35",
            "StartTime": "1640005804",
            "EndTime": "1641505804",
            "Place": "w",
            "Digest": ""
        },
        {
            "ID": 7,
            "CreatedAt": "2021-04-15T14:53:20.571+08:00",
            "UpdatedAt": "2021-04-15T14:53:20.571+08:00",
            "DeletedAt": null,
            "Title": "test1888me",
            "Content": "sd54",
            "UserId": "35",
            "StartTime": "1640",
            "EndTime": "1641505804",
            "Place": "w",
            "Digest": ""
        },
        {
            "ID": 8,
            "CreatedAt": "2021-04-15T14:55:46.444+08:00",
            "UpdatedAt": "2021-04-15T14:55:46.444+08:00",
            "DeletedAt": null,
            "Title": "test1888megggg",
            "Content": "sd54",
            "UserId": "35",
            "StartTime": "1640",
            "EndTime": "1641505804",
            "Place": "w",
            "Digest": ""
        },
        {
            "ID": 4,
            "CreatedAt": "2021-04-13T22:53:52.205+08:00",
            "UpdatedAt": "2021-04-13T22:53:52.205+08:00",
            "DeletedAt": null,
            "Title": "test1888me",
            "Content": "sd54",
            "UserId": "35",
            "StartTime": "1213",
            "EndTime": "1641505804",
            "Place": "w",
            "Digest": ""
        }
    ],
    "message": "成功"
}
```

## 方法描述：根据举办者获取活动

URL地址：/activities/host

请求方法：GET

请求体：

| Key | Value | 说明 |
| ---  | ---  | ---  |
|  host  |  int  | 主办方id |
|  pre   |  int  | 上一页最小id，第一页为0|

请求体示例

```
/activities/host？host=1 pre=0
```

响应体：

| 字段 | 类型 | 说明 |
| ---  | ---  | ---  |
|  code  |  int  |  状态码  |
|  message  |  string  |  状态码描述  |
|  data  |  JSON  |  具体看样例  |

响应体示例

```
{
    "code": 0,
    "data": [
        {
            "Content": "sd5446568498698649868554638458684567",
            "CreatedAt": "2021-04-16T12:08:51.556+08:00",
            "DeletedAt": null,
            "Digest": "sd5446568498698649868554638458684567",
            "EndTime": "1641505804",
            "ID": 9,
            "Place": "w",
            "StartTime": "1640",
            "Title": "test1888megggg",
            "UpdatedAt": "2021-04-16T12:08:51.556+08:00",
            "UserId": "77"
        },
        {
            "Content": "sd5446568498698649868554638458684567dddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddd",
            "CreatedAt": "2021-04-16T12:09:20.217+08:00",
            "DeletedAt": null,
            "Digest": "sd5446568498698649868554638458684567dddddddddddddddddddddddd",
            "EndTime": "1641505804",
            "ID": 10,
            "Place": "w",
            "StartTime": "1640",
            "Title": "test1888megggg",
            "UpdatedAt": "2021-04-16T12:09:20.217+08:00",
            "UserId": "77"
        }
    ],
    "message": "成功"
}
```