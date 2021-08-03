# 2021-hackathon-backend

# 图片API文档

host: https://nspyf.top

port：11000

## 方法描述：上传图片

URL地址：/picture

请求方法：POST

请求体：

**form-data**

| Key | Value | 说明 |
| ---  | ---  | ---  |
|  file  |    |  上传的图片 |

请求体示例

本地测试时使用的例子:

```
curl --location --request POST 'localhost:21000/picture' --form file=@/home/nspyf/下载/git.jpg
```

响应体：

**说明：成功上传文件后响应体里有文件路径**

| 字段 | 类型 | 说明 |
| ---  | ---  | ---  |
|  code  |  int  |  状态码  |
|  message  |  string  |  状态码描述  |
|  data  |  JSON  |  含文件名，具体看样例  |

响应体示例

```
{
    "code":0,
    "data":{
        "file":"https://nspyf.oss-cn-shanghai.aliyuncs.com/caf495c9caf495c9git.jpg"
    },
    "message":"成功"
}
```