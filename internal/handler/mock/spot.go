package mock

import "github.com/dena-gohost/gohost-server/gen/api"

func strev(s string) *string {
	return &s
}

var Spots = []*api.Spot{
	{
		Id:          strev("7ed0c28c-890a-4c3e-8c0d-6db93969168c"),
		Name:        strev("油井グランドホテル"),
		Description: strev("千葉県東金市油井にある廃ホテル。正式名称『油井グランドホテル』。心霊スポットとして知られていたが、ある事件を境に全国的に有名となる。「二階にある焦げた部屋が特に危険」という噂がある。"),
		ImageUrl:    strev("https://shinreispot.com/wp-content/uploads/2018/02/yjimage.jpg"),
		Address:     strev("千葉県東金市油井263-1"),
	},
	{
		Id:          strev("e74b267c-ae37-4a8e-a5f2-68252ecc4f86"),
		Name:        strev("正丸峠"),
		Description: strev("飯能市と秩父郡横瀬町の間にある峠です。道幅が狭くカーブも多いため、事故が絶えないようです。\nそのせいか、口裂け女や子供の霊、四つん這いで追いかけてくる霊などの様々な目撃情報が多数。街灯がない山道なので、夜のドライブは充分注意してくださいね。"),
		ImageUrl:    strev("https://www.tohge-project.jp/upload/shop/15-0l.jpg"),
		Address:     strev("埼玉県飯能市・秩父郡横瀬町"),
	},
}
