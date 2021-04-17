# 2021-hackathon-backend

# 管理员文档

## 权限：修改用户verification字段

说明：

用户verification字段

| value | 说明 |
| --- | --- |
| ""  | 普通用户  |
| "v"  | 认证的组织（学院或社团）  |


## 请求方法（转自用户信息文档）：

### 方法描述：修改用户认证信息

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