**题目：**

我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？



**答案：**

应该Wrap error抛给上层，dao层只需要负责逻辑层面，具体的错误判断应该由上层业务层自己判断并控制处理，dao层只需要如实上报实际情况。

dao:

```go

package dao

import (
	"database/sql"
	"github.com/pkg/errors"
)

var db *sql.DB

type User struct {
	ID    int
	Phone string
}

func GetUserById(id int) (user *User, err error) {
	user = &User{ID: id}
	err = db.QueryRow("select * from user where id=?", id).Scan(&user)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// GO认为是错误，但是上层也需要判断。
			return user, nil
		} else {
			err = errors.Wrap(err, "failed to find user")
			return nil, err
		}
	}
	return
}

```

BLL:
```go
user, err := dao.GetUserById(id)

if errors.Is(err, sql.ErrNoRows} {

}
```