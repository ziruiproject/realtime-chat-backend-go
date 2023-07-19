# REALTIME CHAT API SPECS
## `/api/v1/`

### Unauthorized
```json
{
  "status": "error",
  "error_type": "client",
  "message": "Unauthorized access. Please authenticate to proceed.",
  "code": 401,
  "data": {}
}
```
### Resources Not Found
```json
{
  "status": "error",
  "error_type": "client",
  "message": "Resource not found. Please check the provided identifier.",
  "code": 404,
  "data": {}
}
```
### Server Error
```json
{
    "status" : "error",
    "error_type" : "server",
    "message" : "An unexpected error occurred on the server.",
    "code" : 500,
    "data" : {}
}
```

# Users
### Get All Users `GET /users`
#### Success
```json
{
    "status" : "OK",
    "message" : "Request was Successful",
    "code" : 200,
    "data" : {
        "users" : [
            {
                "id" : "832a5b22-77f5-4c03-8cb4-318fd8192b07",
                "name" : "Yudha Prawira",
                "username" : "@zirui.dev",
                "email" : "example@gmail.com",
                "profile_img" : "/path/to/img.png",
                "create_at" : 1689721714,
                "updated_at" : 1689721714
            },
            {
                "id" : "d5fd602b-a82e-4b80-a53a-f93b180f3e42",
                "name" : "Prawira Sugiharto",
                "username" : "@ziruiproject",
                "email" : "moreexample@gmail.com",
                "profile_img" : "/path/to/img.png",
                "create_at" : 1689721714,
                "updated_at" : 1689721714
            },
        ]
    }
}
```
## Get User by Id `GET /users/id`
#### Success
```json
{
    "status" : "OK",
    "message" : "User retrieved successfully",
    "code" : 200,
    "data" : {
        "id" : "832a5b22-77f5-4c03-8cb4-318fd8192b07",
        "name" : "Yudha Prawira",
        "username" : "@zirui.dev",
        "email" : "example@gmail.com",
        "profile_img" : "/path/to/img.png",
        "create_at" : 1689721714,
        "updated_at" : 1689721714
    }
}
```
## Create User `POST /users`
#### Success
```json
{
    "status" : "OK",
    "message" : "User created successfully",
    "code" : 200,
    "data" : {
        "id" : "832a5b22-77f5-4c03-8cb4-318fd8192b07",
        "name" : "Yudha Prawira",
        "username" : "@zirui.dev",
        "email" : "example@gmail.com",
        "profile_img" : "/path/to/img.png",
        "create_at" : 1689721714,
        "updated_at" : 1689721714
    }
}
```
## Update User `PATCH /users/id`
#### Success
```json
{
    "status" : "OK",
    "message" : "User updated successfully",
    "code" : 200,
    "data" : {
        "id" : "832a5b22-77f5-4c03-8cb4-318fd8192b07",
        "name" : "Yudha Prawira",
        "username" : "@zirui.dev",
        "email" : "example@gmail.com",
        "profile_img" : "/path/to/img.png",
        "create_at" : 1689721714,
        "updated_at" : 1689746996
    }
}
```
#### Validation Error
```json
{
  "status": "error",
  "error_type": "client",
  "message": "Validation failed. Please correct the provided data.",
  "code": 422,
  "data": {
    "errors": {
        "name" : [
            "Required",
            "No weird symbols"
        ],
        "password" : [
            "Required",
            "Min 8 characters"
        ],
        "username" : [
            "Required",
            "Max 20 characters",
            "This username already taken"
        ],
        "email" : [
            "Required",
            "Must be a valid email",
        ]
    }
  }
}
```
#### Duplicate Email Error
```json
{
  "status": "error",
  "error_type": "client",
  "message": "A user with the provided email already exists.",
  "code": 409,
  "data": {}
}
```
#### Duplicate Username Error
```json
{
  "status": "error",
  "error_type": "client",
  "message": "A user with the provided username already exists.",
  "code": 409,
  "data": {}
}
```
## Delete User `PATCH /users/id`
#### Success
```json
{
    "status" : "OK",
    "message" : "User updated successfully",
    "code" : 200,
    "data" : {
        "id" : "832a5b22-77f5-4c03-8cb4-318fd8192b07",
        "name" : "Yudha Prawira",
        "username" : "@zirui.dev",
        "email" : "example@gmail.com",
        "profile_img" : "/path/to/img.png",
        "create_at" : 1689721714,
        "updated_at" : 1689746996
    }
}
```


