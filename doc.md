# Todo List API 文档

## 基础信息
- 基础URL: `http://localhost:8080`
- 所有需要认证的接口都需要在请求头中携带 `Authorization: Bearer <token>`

## 用户相关接口

### 注册用户
- 方法: `POST`
- 路径: `/api/register`
- 描述: 创建新用户账号
- 请求体:
```json
{
    "username": "string",
    "password": "string"
}
```
- 响应:
```json
{
    "token": "string",
    "user": {
        "id": "number",
        "username": "string"
    }
}
```

### 用户登录
- 方法: `POST`
- 路径: `/api/login`
- 描述: 用户登录获取token
- 请求体:
```json
{
    "username": "string",
    "password": "string"
}
```
- 响应:
```json
{
    "token": "string",
    "user": {
        "id": "number",
        "username": "string"
    }
}
```

## Todo相关接口

### 创建Todo
- 方法: `POST`
- 路径: `/api/todos`
- 描述: 创建新的待办事项
- 认证: 需要
- 请求体:
```json
{
    "title": "string",
    "description": "string"
}
```
- 响应:
```json
{
    "id": "number",
    "title": "string",
    "description": "string",
    "completed": false,
    "created_at": "string",
    "updated_at": "string"
}
```

### 获取Todo列表
- 方法: `GET`
- 路径: `/api/todos`
- 描述: 获取当前用户的所有待办事项
- 认证: 需要
- 响应:
```json
[
    {
        "id": "number",
        "title": "string",
        "description": "string",
        "completed": "boolean",
        "created_at": "string",
        "updated_at": "string"
    }
]
```

### 获取单个Todo
- 方法: `GET`
- 路径: `/api/todos/:id`
- 描述: 获取指定ID的待办事项
- 认证: 需要
- 响应:
```json
{
    "id": "number",
    "title": "string",
    "description": "string",
    "completed": "boolean",
    "created_at": "string",
    "updated_at": "string"
}
```

### 更新Todo
- 方法: `PUT`
- 路径: `/api/todos/:id`
- 描述: 更新指定ID的待办事项
- 认证: 需要
- 请求体:
```json
{
    "title": "string",
    "description": "string",
    "completed": "boolean"
}
```
- 响应:
```json
{
    "id": "number",
    "title": "string",
    "description": "string",
    "completed": "boolean",
    "created_at": "string",
    "updated_at": "string"
}
```

### 删除Todo
- 方法: `DELETE`
- 路径: `/api/todos/:id`
- 描述: 删除指定ID的待办事项
- 认证: 需要
- 响应: HTTP 204 No Content

## 错误响应
所有接口在发生错误时会返回以下格式的响应：
```json
{
    "error": "错误信息描述"
}
```

常见HTTP状态码：
- 200: 请求成功
- 201: 创建成功
- 204: 删除成功
- 400: 请求参数错误
- 401: 未授权或token无效
- 403: 禁止访问
- 404: 资源不存在
- 500: 服务器内部错误