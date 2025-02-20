# Todo List API 文档

## 基础信息

-   基础 URL: `http://localhost:8080`

## Todo 相关接口

### 创建 Todo

-   方法: `POST`
-   路径: `/api/todos`
-   描述: 创建新的待办事项
-   请求体:

```json
{
    "title": "string",
    "description": "string"
}
```

-   响应: `201 Created`

```json
{
    "ID": "number",
    "title": "string",
    "description": "string",
    "completed": false,
    "CreatedAt": "string",
    "UpdatedAt": "string",
    "DeletedAt": null
}
```

### 获取 Todo 列表

-   方法: `GET`
-   路径: `/api/todos`
-   描述: 获取所有待办事项
-   响应: `200 OK`

```json
[
    {
        "ID": "number",
        "title": "string",
        "description": "string",
        "completed": "boolean",
        "CreatedAt": "string",
        "UpdatedAt": "string",
        "DeletedAt": null
    }
]
```

### 获取单个 Todo

-   方法: `GET`
-   路径: `/api/todos/:id`
-   描述: 获取指定 ID 的待办事项
-   响应: `200 OK`

```json
{
    "ID": "number",
    "title": "string",
    "description": "string",
    "completed": "boolean",
    "CreatedAt": "string",
    "UpdatedAt": "string",
    "DeletedAt": null
}
```

### 更新 Todo

-   方法: `PUT`
-   路径: `/api/todos/:id`
-   描述: 更新指定 ID 的待办事项
-   请求体:

```json
{
    "title": "string",
    "description": "string",
    "completed": "boolean"
}
```

-   响应: `200 OK`

```json
{
    "ID": "number",
    "title": "string",
    "description": "string",
    "completed": "boolean",
    "CreatedAt": "string",
    "UpdatedAt": "string",
    "DeletedAt": null
}
```

### 删除 Todo

-   方法: `DELETE`
-   路径: `/api/todos/:id`
-   描述: 删除指定 ID 的待办事项
-   响应: `204 No Content`

## 错误响应

所有接口在发生错误时会返回以下格式的响应：

```json
{
    "error": "错误信息描述"
}
```

常见 HTTP 状态码：

-   201: 创建成功
-   200: 请求成功
-   204: 删除成功
-   400: 请求参数错误
-   404: 资源不存在
-   500: 服务器内部错误
