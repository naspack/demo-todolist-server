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
    "id": "number",
    "title": "string",
    "description": "string",
    "completed": false,
    "created_at": "string",
    "updated_at": "string",
    "deleted_at": null
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
        "id": "number",
        "title": "string",
        "description": "string",
        "completed": "boolean",
        "created_at": "string",
        "updated_at": "string",
        "deleted_at": null
    }
]
```

### 获取单个 Todo

-   方法: `POST`
-   路径: `/api/todos/get`
-   描述: 获取指定 ID 的待办事项
-   请求体:

```json
{
    "id": "number"
}
```

-   响应: `200 OK`

```json
{
    "id": "number",
    "title": "string",
    "description": "string",
    "completed": "boolean",
    "created_at": "string",
    "updated_at": "string",
    "deleted_at": null
}
```

### 更新 Todo

-   方法: `POST`
-   路径: `/api/todos/update`
-   描述: 更新指定 ID 的待办事项
-   请求体:

```json
{
    "id": "number",
    "title": "string",
    "description": "string",
    "completed": "boolean"
}
```

-   响应: `200 OK`

```json
{
    "id": "number",
    "title": "string",
    "description": "string",
    "completed": "boolean",
    "created_at": "string",
    "updated_at": "string",
    "deleted_at": null
}
```

### 完成 Todo

-   方法: `POST`
-   路径: `/api/todos/complete`
-   描述: 将指定 ID 的待办事项标记为已完成
-   请求体:

```json
{
    "id": "number"
}
```

-   响应: `200 OK`

```json
{
    "id": "number",
    "title": "string",
    "description": "string",
    "completed": true,
    "created_at": "string",
    "updated_at": "string",
    "deleted_at": null
}
```

### 取消完成 Todo

-   方法: `POST`
-   路径: `/api/todos/uncomplete`
-   描述: 将指定 ID 的待办事项标记为未完成
-   请求体:

```json
{
    "id": "number"
}
```

-   响应: `200 OK`

```json
{
    "id": "number",
    "title": "string",
    "description": "string",
    "completed": false,
    "created_at": "string",
    "updated_at": "string",
    "deleted_at": null
}
```

### 删除 Todo

-   方法: `POST`
-   路径: `/api/todos/delete`
-   描述: 删除指定 ID 的待办事项
-   请求体:

```json
{
    "id": "number"
}
```

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
