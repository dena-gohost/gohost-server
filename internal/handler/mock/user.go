package mock

import (
	"time"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"

	"github.com/dena-gohost/gohost-server/gen/api"
)

func daterev(t time.Time) *openapi_types.Date {
	return &openapi_types.Date{t}
}

func intrev(i int) *int {
	return &i
}

var Users = []api.User{
	{
		Id:           strev("aeb4238a-0582-408f-a68e-fc3f8163fc87"),
		FirstName:    strev("DeNA"),
		LastName:     strev("ハッカソン"),
		UserName:     strev("denason"),
		Email:        strev("hackathon@dena.ac.jp"),
		UniversityId: strev("06faba8c-5d9f-495c-a9d7-3bc82a040398"),
		BirthDate:    daterev(time.Date(1999, 9, 10, 0, 0, 0, 0, time.Local)),
		Year:         intrev(4),
		GenderId:     strev("1"),
		IconUrl:      strev("https://www.photo-ac.com/assets/img/ai_model_v2/model_6.png"),
		InstagramId:  strev("dummy"),
	},
	{
		Id:           strev("1f4e3db9-e6ec-4028-befe-4f92c2bb2051"),
		FirstName:    strev("でーぬえー"),
		LastName:     strev("programming"),
		UserName:     strev("degramming"),
		Email:        strev("dena@programming.ac.jp"),
		UniversityId: strev("06faba8c-5d9f-495c-a9d7-3bc82a040398"),
		BirthDate:    daterev(time.Date(2001, 9, 11, 0, 0, 0, 0, time.Local)),
		Year:         intrev(1),
		GenderId:     strev("2"),
		IconUrl:      strev("https://cdn-ak.f.st-hatena.com/images/fotolife/c/color-hiyoko/20190924/20190924145352.jpg"),
		InstagramId:  strev("yummy"),
	},
}

var Genders = []api.Gender{
	{
		strev("1"),
		strev("男"),
	},
	{
		strev("2"),
		strev("女"),
	},
	{
		strev("3"),
		strev("その他"),
	},
}
