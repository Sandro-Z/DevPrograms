# 临时写的api文档

这是一个模拟学生点菜与商家开户的系统后端

## 各个用户类型与可用范围有

1. 顾客（用户类型1）登录，查看商家，进行点餐，查看当前单
2. 商家（用户类型2）登录，申请开店，发布商品，确认接单，确认结单
2. 管理员（用户类型3）登录，审批开店
3. 基本流程： 商家申请开店=>管理员通过=>商家上架商品及当前剩余量、是否开店=>用户点单=>商家确认接单=>用户查看当前单数目=>商家结单=>交易完成



## 响应

### 当请求成功时，返回以下内容

Response 200

```json
{
  "success": true,
  "data": {
      ...
  },
}
```

  


### 当请求失败时，返回以下内容

Response 200

```json
{
  "success": false,
  "message": "报错信息",
  "code": uint
}
```

  

# 模块

## 模块目录

[用户1模块(用户1所有模块对其他类型用户均可用)](#学生(用户1)模块)

[用户2模块](#user2)

[管理员模块](#admin)

[测试模块](#test)


## 学生(用户1)模块

### 接口详情

除了注册和登录，其他API在未登录的情况下不能操作，返回未登录的信息。

### 接口列表

[用户注册](#用户注册) `POST {base_url}/api/register`

[用户登录](#用户登录) `POST {bause_url}/api/login`

[用户登出](#用户登出) `DELETE {base_url}/api/logout`

[用户查询当前所有的店家](#用户查询当前所有的店家) `GET {base_url}/api/getRestaurants`

[用户查询本店所有菜谱信息](#用户查询本店所有菜谱信息)  `GET {base_url}/api/foods`

[选择一份菜品](#选择一份菜品) `POST {base_url}/api/food/select`

[用户查询餐品详细信息](#用户查询餐品详细信息) `GET {base_url}/api/food/:id`

[获取当前单号信息](#获取当前单号信息) `GET {base_url}/api/foods/status/:id`

[用餐后打分](#用餐后打分) `POST {base_url}/api/foods/like`



#### 用户注册 
`POST {base_url}/api/register`

tips: 
用户在注册时学/工号不能重复，如果重复，则注册失败
注册的同时登录
初始为非管理员（要变成管理员需要修改数据库）
注册，登录和登出使用session处理

Request 
```json
{
  "number": "学号",
  "name": "姓名",
  "password": "密码",
  "type": "我是店家/学生"
}
```

Response
```json
{
  "success": true,
  "data": {
    "userId": 1, //用户id
    "number": "学号", //这一项为字符串
    "name": "姓名" // 这一项为字符串
  }
}
```


#### 用户登录 
`POST {base_url}/api/login`

Request 
```json
{
  "number": "学号",
  "password": "密码"
}
```


Response
```json
{
  "success": true,
  "data": {
    "userId": 1, //用户id
    "number": "学号",
    "name": "姓名"
  }
}
```

#### 用户登出 
`DELETE {base_url}/api/logout`


Response
```json
{
  "success": true,
  "data": {
    "login": true //如果之前处于登录状态，则为true，否则为false
  }
}
```


#### 用户查询当前所有的店家 
`GET {base_url}/api/getRestaurants`

查询参数为页数和每页的限制数量，即返回数据为(page-1) * limit+1条到page * limit条记录，此后的查询和此方式一致

Request 


```json
{
  "page": 1, //第几页(从1开始)
  "limit": 10 //一页几条记录
}
```

Response
```json
{
  "success": true,
  "data": {
    "total": 100, //总数量
    "list": [
      {
        "id": 1, //店家id
        "name": "店名",
        "status": 1, //是否开业
        "address": "nalanala", //店家地址
        "people": 1 // 总用餐人数（结单后计数器增加）
      },
      ···
      ]
  }
}
```



#### 用户查询本店所有菜谱信息 
`GET {base_url}/api/foods`


Request 


```json
{
  "page": 1, //第几页(从1开始)
  "limit": 10, //一页几条记录
  "id": 1 //店子的ID
}
```


Response
```json
{
  "success": true,
  "data": {
    "total": 100, //总数量
    "list": [
      {
        "id": 1, //菜品id
        "name": "菜名",
        "costs": 20, //菜品价格~~USD~~
        "number": 18, //点过该餐品总人数
        "like": 1, // 推荐人数
        "foodless": 2, // 当前剩余餐数
      },
      ···
      ]
  }
}
```


#### 用户查询餐品详细信息 
`GET {base_url}/api/food/:id` 



Response
```json
{
  "success": true,
  "data": {
    "id": 1, //菜品id
    "name": "菜名",
    "costs": 20, //菜品价格~~USD~~
    "number": 18, //点过该餐品总人数
    "like": 1, // 推荐人数
    "time": 2, // 当前剩余餐数
    "describes": "xxx",
  }
}
```


#### 选择一份菜品 
`POST {base_url}/api/food/select`

如果菜品剩余为0或商家未开店(审核未通过)，则不可选择

Request 
```json
{
  "Id": 1,
}
```
Response
```json
{
  "success": true,
  "data": {
    "status": 1, //店家未接单
    "titleId": 1 //当前单号
  }
}
```


#### 获取当前状态信息 
`GET {base_url}/api/foods/status/:id`

Response
```json
{
  "success": true,
  "data": {
    "status": 1
  }
}
```



#### 用餐后打分（仅限结单后） 
`POST {base_url}/api/foods/like`

```json
{
  "titleId": 123, //订单id
  "like": 0/1,
}
```

Response

```json
{
  "success": true,
}
```



## 商家模块

### 接口列表

[上架菜品列表](#上架菜品列表) `POST {base_url}/api/restaurant/food/add`

[修改菜品](#修改菜品) `POST {base_url}/api/restaurant/food/change`

[查看当前点单](#查看当前点单) `GET {base_url}/api/restaurant/checkAll`

[改变订单状态（不可逆）](#改变订单状态) `POST {base_url}/api/restaurant/changeStatus`

[申请开店（仅可开一家店）](#申请开店)`POST {base_url}/api/restaurant/apply`

[查看开店审核状态](#查看开店审核状态) `GET {base_url}/api/restaurant/checkStatus`



#### 上架菜品列表 
`POST {base_url}/api/restaurant/food/add`

Request 
```json
{
  "foodName": 1, 
  "foodLess": 1, //剩余菜品数
  "cost": 114,
  "describe": "xxx"
}
```


Response

```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "name": "xxx",
      "describe": "xxx",
      "cost": "114514",
      "createdAt": "nalanal",
      "updatedAt": "baabwhnarja",
    },
    ...
  ]
}
```




#### 修改菜品 
`POST {base_url}/api/restaurant/food/change`

Request

```json
{
  "foodId":1,
  //修改任意一项
}
```

Response

```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "name": "xxx",
      "describe": "xxx",
      "cost": "114514",
      "foodless": 1,//剩余菜品数
      "createdAt":"222",
      "updatedAt":"2131412414",
    },
    ...
  ]
}
```




#### 查看当前点单
`GET {base_url}/api/restaurant/checkAll`

Request 
```json
{
  "page": 1, //第几页(从1开始)
  "limit": 10, //一页几条记录
  "id": 1 //店子的ID
}
```
Response

```json
{
  "success": true,
  "data": {
    "id":"xxx",
    "foodName":"xxx",
    "foodId": 1,
    "userid": "xxx",
    
  }
}
```




#### 改变订单状态
`POST {base_url}/api/restaurant/changeStatus`

Request

```json
{
  "id":1,
  "changeStatus":"xx"
}
```

Response

```json
{
  "success": true
}
```




#### 申请开店
`POST {base_url}/api/restaurant/apply`

Request
```json
{
  "name":"Xxx",
  "describe":"xxx",
  "address":"xxx",
  "licenceStar": 1-5,
}
```

Response
```json
{
  "success": true,
  "data": {
    "status": //审核中
  }
}

```




#### 查看开店状态 
`GET {base_url}/api/restaurant/checkStatus`

Request

Response
```json
{
  "success": true,
  "data": {
    "status": //通过、未通过、审核中
  }
}
```




## 管理员模块

### 接口列表

[查看申请开店列表](#查看申请开店列表) `GET {base_url}/api/admin/restaurant`

[查询本店所有菜谱信息](#查询本店所有菜谱信息) `GET {base_url}/api/admin/food`

[修改审核状态](#修改审核状态) `POST {base_url}/api/admin/restaurant/check`




### 接口详情




#### 查看申请开店列表
`GET {base_url}/api/admin/restaurant`

Request 
```json
{
  "page": 1, //第几页
  "limit": 1,   //一页几条记录
}
```

Response
```json
{
  "success": true,
  "data": {
    "total": 100, //总数
    "list": [
      {
        "id": 1, //用户的id
        "restaurantName": "xxx",
        "describe": "xx",
        "userId":123123,
        "status":"xx",
      },
      ...
    ]
  }
}
```



#### 查询本店所有菜谱信息 
`GET {base_url}/api/admin/foods`

Request


```json
{
  "Id": 1
  "page": 1, //第几页(从1开始)
  "limit": 10 //一页几条记录
}
```

Response
```json
{
  "success": true,
  "data": {
    "total": 100, //总数量
    "list": [
      {
        "id": 1, //菜品id
        "name": "菜名",
        "costs": 20, //菜品价格~~USD~~
        "number": 18, //点过该餐品总人数
        "like": 1, // 推荐人数
        "time": 2, // 当前剩余餐数
      },
      ···
    ]
  }
}
```




#### 修改审核状态 
`POST {base_url}/api/admin/restaurant/check`

Request 
```json
{
  "Id":"xx",
  "status": "22",
}
```

Response
```json
{
  "success": true
}
```

## 测试模块

### 接口列表

[测试服务是否启动](#测试服务是否启动)

[查询当前登录用户信息](#查询当前登录用户信息)

[查询当前登录商家信息](#查询当前登录商家信息)


### 接口详情



#### 测试服务是否启动 
`GET {base_url}/api`

Request  

Response  
```json
{
  "success": true,
  "data": "一些欢迎语句"
}
```




#### 查询当前登录用户信息 
`GET {base_url}/api/userinfo`

Response  
```json
{
  "success": true,
  "username": "当前登录用户名字"
}
```



#### 查询当前登录商家信息 
`GET {base_url}/api/restaurant`

Response

当当前登录用户是商家时
```json  
{
  "success": true,
  "restaurant": "当前登录商家名字"
}
```

反之
```json
{
  "success": false,
  "error": "当前登录非商家"
}
```