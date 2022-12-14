package main

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"regexp"
	"time"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/pkg/errors"

	"github.com/dena-gohost/gohost-server/gen/api"
	environment "github.com/dena-gohost/gohost-server/internal/env"
	"github.com/dena-gohost/gohost-server/internal/service"
)

const (
	registerURL   = "http://localhost:5050/register"
	loginURL      = "http://localhost:5050/login"
	spotsURL      = "http://localhost:5050/spots?date=2022-09-10&limit=3"
	spotURL       = "http://localhost:5050/spots/%s"
	spotEntryURL  = "http://localhost:5050/spots/%s/entry"
	planURL       = "http://localhost:5050/plan"
	planCancelURL = "http://localhost:5050/plan/cancel"
	planFinishURL = "http://localhost:5050/plan/finish"
)

type PlanFinishInfo = struct {
	Like   *bool   `json:"like,omitempty"`
	UserId *string `json:"user_id,omitempty"`
}

var (
	users = []api.User{
		{
			FirstName:    strev("DeNA"),
			LastName:     strev("ハッカソン"),
			UserName:     strev("denason"),
			Email:        strev("hackathon@dena.ac.jp"),
			Password:     strev("passw0rd"),
			UniversityId: strev("06faba8c-5d9f-495c-a9d7-3bc82a040398"),
			BirthDate:    daterev(time.Date(1999, 9, 10, 0, 0, 0, 0, time.Local)),
			Year:         intrev(4),
			GenderId:     strev("1"),
			IconUrl:      strev("https://www.photo-ac.com/assets/img/ai_model_v2/model_6.png"),
			InstagramId:  strev("dummy"),
		},
		{
			FirstName:    strev("でーぬえー"),
			LastName:     strev("programming"),
			UserName:     strev("degramming"),
			Email:        strev("dena@programming.ac.jp"),
			Password:     strev("passw0rd"),
			UniversityId: strev("06faba8c-5d9f-495c-a9d7-3bc82a040398"),
			BirthDate:    daterev(time.Date(2001, 9, 11, 0, 0, 0, 0, time.Local)),
			Year:         intrev(1),
			GenderId:     strev("2"),
			IconUrl:      strev("https://cdn-ak.f.st-hatena.com/images/fotolife/c/color-hiyoko/20190924/20190924145352.jpg"),
			InstagramId:  strev("yummy"),
		},
		{
			FirstName:    strev("ハッカソン"),
			LastName:     strev("dena"),
			UserName:     strev("hackathon"),
			Email:        strev("dena@hackathon.ac.jp"),
			Password:     strev("passw0rd"),
			UniversityId: strev("06faba8c-5d9f-495c-a9d7-3bc82a040398"),
			BirthDate:    daterev(time.Date(2001, 9, 12, 0, 0, 0, 0, time.Local)),
			Year:         intrev(3),
			GenderId:     strev("3"),
			IconUrl:      strev("http://df.iwatobi-sc.com/img/story/09_still_01.jpg"),
			InstagramId:  strev("dummy_insta"),
		},
	}
	date = api.PostSpotsSpotIdEntryJSONRequestBody{
		Date: &openapi_types.Date{
			time.Date(2020, 9, 11, 0, 0, 0, 0, time.Local),
		},
	}
)

func strev(s string) *string {
	return &s
}

func daterev(t time.Time) *openapi_types.Date {
	return &openapi_types.Date{t}
}

func intrev(i int) *int {
	return &i
}

func main() {

	cookies := make([]string, 0)
	env, err := environment.Process()
	if err != nil {
		panic(err)
	}

	for _, user := range users {
		// ユーザー登録
		payload, err := json.Marshal(user)
		if err != nil {
			panic(err)
		}

		_, _ = post(payload, registerURL, nil)

		// ログイン
		rh, _ := post(payload, loginURL, nil)

		r := regexp.MustCompile("Set-Cookie: (.*? )")
		fss := r.FindStringSubmatch(string(rh))
		if len(fss) != 2 {
			panic(errors.Errorf("cannot parse %s", rh))
		}

		cookie := fss[1]
		cookies = append(cookies, cookie)

		// おすすめ心霊スポット
		ba := get(spotsURL, cookie)

		spots := make([]api.Spot, 0, 3)
		if err := json.Unmarshal(ba, &spots); err != nil {
			panic(err)
		}

		if len(spots) == 0 {
			panic(errors.New("len(spots) == 0"))
		}

		spotID := spots[0].Id

		// 詳細情報取得
		_ = get(fmt.Sprintf(spotURL, *spotID), cookie)

		// エントリー
		payload, err = json.Marshal(date)
		if err != nil {
			panic(err)
		}

		_, _ = post(payload, fmt.Sprintf(spotEntryURL, *spotID), &cookie)
	}
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", env.DBUser, env.DBPass, env.DBHost, env.DBPort, env.DBName))
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	txn, err := db.Begin()
	defer txn.Rollback()
	if err != nil {
		panic(err)
	}

	if err := service.Match(ctx, txn); err != nil {
		panic(err)
	}

	if err := txn.Commit(); err != nil {
		panic(err)
	}

	for i, _ := range users {
		_ = get(planURL, cookies[i])
	}

}

func post(payload []byte, url string, cookie *string) ([]byte, *http.Response) {
	client := &http.Client{}

	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer(payload),
	)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	if cookie != nil {
		req.Header.Set("Cookie", *cookie)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	rh, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}

	ba, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(ba))

	return rh, resp
}

func get(url, cookie string) []byte {
	client := &http.Client{}

	req, err := http.NewRequest(
		"GET",
		url,
		nil)

	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", cookie)

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	ba, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(ba))

	return ba
}
