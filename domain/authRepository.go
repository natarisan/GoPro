package domain

import(
	"encoding/json"
	"fmt"
	"github.com/natarisan/gop-libs/logger"
	"net/http"
	"net/url"
)

type AuthRepository interface {
	IsAuthorized(token string, routeName string, vars map[string]string) bool
}

type RemoteAuthRepository struct {
}

func (r RemoteAuthRepository) IsAuthorized(token string, routeName string, vars map[string]string) bool {

	u := buildVerifyURL(token, routeName, vars)

	if response, err := http.Get(u); err != nil {
		fmt.Println("Error while sending..." + err.Error())
		return false
	} else {
		m := map[int]bool{}
		if err = json.NewDecoder(response.Body).Decode(&m); err != nil {
			logger.Error("サーバーからのレスポンスをデコードしているときにエラー:" + err.Error())
			return false
		}
		return m[1000]
	}
}

//URLを作成する
func buildVerifyURL(token string, routeName string, vars map[string]string) string{
    u := url.URL{Host: "localhost:6001", Path: "/auth/verify", Scheme: "http"}
	q := u.Query()
	q.Add("token", token)
	q.Add("routeName", routeName)
	for k, v := range vars {
		q.Add(k,v)
	}
	u.RawQuery = q.Encode()
	return u.String()
}

func NewAuthRepository() RemoteAuthRepository {
	return RemoteAuthRepository{}
}