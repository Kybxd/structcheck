# structcheck

Check if a struct has all the things you expect.


# examle

```golang
type ExampleRsp struct {
	Result         *int32      `json:"Result"`
	RetMsg         *string     `json:"RetMsg"`
	Info           *UserInfo   `json:"Info"`
	Friends        []*UserInfo `json:"Friends"`
	LastLoginTime  *int64      `json:"LastLoginTime"`
	LastLogoutTime *int64      `json:"LastLogoutTime"`
	IsOnline       *bool       `json:"IsOnline"`
}

type UserInfo struct {
	Id    *string `json:"Id"`
	Name  *string `json:"Name"`
	Level *uint32 `json:"Level"`
}

rsp, err := examplehttp.Post(exampleUrl, exampleReq)
expected := &ExampleRsp{
    Result: Int32(0),
    RetMsg: String("query user info success"),
    Info: &UserInfo{
        Id:    String("1234567890"),
        Name:  String("kybxd"),
        Level: Uint32(100),
    },
    Friends: []*UserInfo{
        {
            Id:   String("1234567891"),
            Name: String("goodfriend"),
        },
        {
            Id:   String("1234567892"),
            Name: String("newfriend"),
        },
    },
    IsOnline: Bool(true),
}
ok, msg := IsExpected(expected, rsp)
if !ok {
    log.Println(msg)
}

```
