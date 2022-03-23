package ginutil

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/cube-group/golib/types/convert"
	"gorm.io/gorm"
	"math"
)

type PageNation struct {
	Page       int
	PageSize   int
	PageOffset int
	TotalPage  int
	TotalCount int
}

//直接返回layui的map
func GetLayuiPage(c *gin.Context, query *gorm.DB, list interface{}, elementId, location string) (gin.H, error) {
	pn := GetPageNation(c, 0)
	var count int64
	if err := query.Count(&count).Limit(pn.PageSize).Offset(pn.PageOffset).Find(list).Error; err != nil {
		return nil, err
	}
	js := `
layui.laypage.render({
                elem: '%s',
                curr: %d,
                count: %d, 
                layout: ['count', 'prev', 'page', 'next', 'limit', 'refresh', 'skip'],
                limits: [10, 20, 50, 100],
                limit: %d,
                jump: function (obj, first) {
                    if (!first) {
                        var jumpUrl = '%s';
                        var jumpPage = 'page='+obj.curr+'&pageSize='+obj.limit;
                        jumpUrl += (jumpUrl.indexOf('?') >= 0 ? ('&' + jumpPage) : ('?' + jumpPage));
						jumpUrl += '%s';
                        window.location = jumpUrl;
                    }
                }
            });
`
	return gin.H{
		"pages": fmt.Sprintf(js, elementId, pn.Page, count, pn.PageSize, location, GetNoPageQuery(c)),
		"list":  list,
	}, nil
}

//纯算法
//从上下文中获取分页信息
func GetPageNation(c *gin.Context, totalCount int) *PageNation {
	pn := new(PageNation)
	pn.Page = convert.MustInt(Input(c, "page"))
	if pn.Page <= 0 {
		pn.Page = 1
	}
	pn.PageSize = convert.MustInt(Input(c, "pageSize"))
	if pn.PageSize <= 0 {
		pn.PageSize = 10
	} else if pn.PageSize > 100 {
		pn.PageSize = 100
	}
	pn.PageOffset = (pn.Page - 1) * pn.PageSize

	if totalCount > 0 {
		pn.TotalPage = int(math.Ceil(float64(totalCount) / float64(pn.PageSize)))
		pn.TotalCount = totalCount
		if pn.Page > pn.TotalPage {
			pn.Page = pn.TotalPage
		} else if pn.Page < 1 {
			pn.Page = 1
		}
	}
	return pn
}

type PaginationOut struct {
	Total    int64
	Page     int
	PageSize int
	List     interface{}
}

func Pagination(db *gorm.DB, page, pageSize int, out interface{}) (
	res *PaginationOut, err error) {
	// out 传给find的值
	if page == 0 {
		page = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}
	res = new(PaginationOut)
	res.Page = page
	res.PageSize = pageSize
	res.List = out
	err = db.Model(out).Count(&res.Total).Offset(pageSize * (page - 1)).Limit(pageSize).Find(out).Error
	return
}
