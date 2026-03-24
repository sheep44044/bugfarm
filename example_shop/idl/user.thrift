namespace go example.shop.user

include "base.thrift"

struct user {
    1: i64 user_id
    2: string user_name
}

struct RegisterReq {
    1: string user_name
    2: string password
}

struct RegisterResp {
    255: base.BaseResp baseResp
}

struct LoginReq {
    1: string user_name
    2: string password
}

struct LoginResp {
    255: base.BaseResp baseResp
}

service UserService {
    RegisterResp Register (1: RegisterReq req)
    LoginResp Login (1: LoginReq req)
}