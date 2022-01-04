package data

import (
	"context"
	"mall-go/app/app/service/internal/biz"
	"mall-go/app/app/service/internal/data/model/theme"

	"github.com/go-kratos/kratos/v2/log"
)

type themeRepo struct {
	data *Data
	log  *log.Helper
}

func NewThemeRepo(data *Data, logger log.Logger) biz.ThemeRepo {
	return &themeRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
func (r *themeRepo) GetThemeWithSpu(ctx context.Context, name string) (t biz.ThemeSpu, err error) {
	po, err := r.data.db.Theme.Query().Where(theme.Name(name)).WithSpu().First(ctx)
	if err != nil {
		return
	}
	spuList := make([]*biz.Spu, 0)
	// for _, v := range po.Edges.Spu {

	// 	spuList = append(spuList, &biz.Spu{
	// 		Id:             v.ID,
	// 		Title:          v.Title,
	// 		Price:          v.Price,
	// 		Subtitle:       v.Subtitle,
	// 		CategoryId:     v.CategoryID,
	// 		RootCategoryId: v.RootCategoryID,
	// 		Img:            v.Img,
	// 		Online:         int32(po.Online),
	// 	})
	// }
	theme := biz.Theme{
		Id:             po.ID,
		Name:           po.Name,
		Title:          po.Title,
		Description:    po.Description,
		EntranceImg:    po.EntranceImg,
		InternalTopImg: po.InternalTopImg,
		TitleImg:       po.TitleImg,
		TplName:        po.TplName,
		Online:         int32(po.Online),
	}
	return biz.ThemeSpu{
		Theme:   theme,
		SpuList: spuList,
	}, nil
}
func (r *themeRepo) GetThemeByNames(ctx context.Context, names []string) (themes []biz.Theme, err error) {
	pos, err := r.data.db.Theme.Query().Where(theme.NameIn(names...)).All(ctx)
	if err != nil {
		return nil, err
	}
	for _, po := range pos {
		themes = append(themes, biz.Theme{
			Id:          po.ID,
			Name:        po.Name,
			Title:       po.Title,
			Description: po.Description,
			EntranceImg: po.EntranceImg,

			InternalTopImg: po.InternalTopImg,
			TitleImg:       po.TitleImg,
			TplName:        po.TplName,
			Online:         int32(po.Online),
		})
	}
	return themes, nil
}