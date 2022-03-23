// Author: chenqionghe
// Time: 2018-10
// 分页相关操作
package page

import (
	"fmt"
	"github.com/cube-group/golib/types/convert"
	"gorm.io/gorm"
	"html/template"
	"github.com/gin-gonic/gin"
	"math"
	"net/url"
	"strconv"
	"strings"
)

//分页专用结构

//分页工具类
type PageInfo struct {
	Index       uint   `json:"index"`       //当前页码
	Size        uint   `json:"size"`        //分页大小
	Total       uint   `json:"total"`       //总条数
	Count       uint   `json:"count"`       //总页数
	Next        uint   `json:"next"`        //下一页
	Pre         uint   `json:"pre"`         //上一页
	Limit       uint   `json:"limit"`       //limit大小
	Offset      uint   `json:"offset"`      //offset大小
	Enabled     bool   `json:"enabled"`     //是否可用分页
	LimitString string `json:"limitString"` //sql limit 0,1
}

//获取offset
func getOffset(page, pageSize uint) uint {
	return (page - 1) * pageSize
}

//获取limit
func getLimit(pageSize uint) uint {
	return pageSize
}

//获取上一页
func getPre(page uint) uint {
	return page - 1
}

//获取下一页
func getNext(page, pageCount uint) uint {
	if pageCount > 1 && page >= 0 && page < pageCount {
		return page + 1
	}
	return 0
}

//获取总页数
func getCount(totalCount, pageSize uint) uint {
	return uint(math.Ceil(float64(totalCount) / float64(pageSize)))
}

//通过参数获取页码和每页数量
//@param options []interface{} 配置项[page, pageSize]
func PageParams(options []interface{}) (uint, uint) {
	var page uint = 1
	var pageSize uint = 10
	optionsLength := len(options)
	if optionsLength > 0 {
		page = options[0].(uint)
	}
	if optionsLength > 1 {
		pageSize = options[1].(uint)
	}
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 1
	} else if pageSize > 100 {
		pageSize = 100
	}
	return page, pageSize
}

//获取分页统一结构
func List(c *gin.Context, list interface{}, query *gorm.DB) (map[string]interface{}, error) {
	var page, pagesize uint
	var totalCount int64

	page = GetPage(c)
	pagesize = GetPageSize(c)
	query.Model(list).Count(&totalCount)

	p := NewPage(page, pagesize, uint(totalCount))

	if err := query.Limit(int(p.Limit)).Offset(int(p.Offset)).Find(list).Error; err != nil {
		return nil, err
	}
	rendering := p.Rendering(c)
	search := GetParams(c)
	return map[string]interface{}{
		"page":      p,
		"list":      list,
		"search":    search,
		"rendering": template.HTML(rendering),
	}, nil

}

//New出分页工具类
func NewPage(page, pageSize, totalCount uint) PageInfo {
	if page == 0 {
		page = 1
	}
	if pageSize == 0 {
		pageSize = 20
	}
	pageCount := getCount(totalCount, pageSize)
	limit := getLimit(pageSize)
	offset := getOffset(page, pageSize)

	return PageInfo{
		Index:       page,
		Size:        pageSize,
		Total:       totalCount,
		Count:       pageCount,
		Enabled:     pageCount == 1,
		Next:        getNext(page, pageCount),
		Pre:         getPre(page),
		Limit:       limit,
		Offset:      offset,
		LimitString: fmt.Sprintf("%v,%v", limit, offset),
	}
}

//Pages 渲染生成html分页标签
func (p *PageInfo) Rendering(c *gin.Context) string {
	queryParams := c.Request.URL.Query()
	//从当前请求中获取page
	page := queryParams.Get("page")
	if page == "" {
		page = "1"
	}
	//将页码转换成整型，以便计算
	pagenum, _ := strconv.Atoi(page)
	if pagenum == 0 {
		return ""
	}

	//计算总页数
	var totalPageNum = int(math.Ceil(float64(p.Total) / float64(p.Size)))

	//首页链接
	var firstLink string
	//上一页链接
	var prevLink string
	//下一页链接
	var nextLink string
	//末页链接
	var lastLink string
	//中间页码链接
	var pageLinks []string

	//总数目信息
	var totalCount string

	//首页和上一页链接
	if pagenum > 1 {
		firstLink = fmt.Sprintf(`<li><a href="%s">首页</a></li>`, p.pageURL(c, "1"))
		prevLink = fmt.Sprintf(`<li><a href="%s">上一页</a></li>`, p.pageURL(c, strconv.Itoa(pagenum-1)))
	} else {
		firstLink = `<li class="disabled"><a href="#">首页</a></li>`
		prevLink = `<li class="disabled"><a href="#">上一页</a></li>`
	}

	//末页和下一页
	if pagenum < totalPageNum {
		lastLink = fmt.Sprintf(`<li><a href="%s">末页</a></li>`, p.pageURL(c, strconv.Itoa(totalPageNum)))
		nextLink = fmt.Sprintf(`<li><a href="%s">下一页</a></li>`, p.pageURL(c, strconv.Itoa(pagenum+1)))
	} else {
		lastLink = `<li class="disabled"><a href="#">末页</a></li>`
		nextLink = `<li class="disabled"><a href="#">下一页</a></li>`
	}

	//生成中间页码链接
	pageLinks = make([]string, 0, 10)
	startPos := pagenum - 3
	endPos := pagenum + 3
	if startPos < 1 {
		endPos = endPos + int(math.Abs(float64(startPos))) + 1
		startPos = 1
	}
	if endPos > totalPageNum {
		endPos = totalPageNum
	}
	for i := startPos; i <= endPos; i++ {
		var s string
		if i == pagenum {
			s = fmt.Sprintf(`<li class="active"><a href="%s">%d</a></li>`, p.pageURL(c, strconv.Itoa(i)), i)
		} else {
			s = fmt.Sprintf(`<li><a href="%s">%d</a></li>`, p.pageURL(c, strconv.Itoa(i)), i)
		}
		pageLinks = append(pageLinks, s)
	}

	////总条数和总页数
	pageCount, _ := convert.String(p.Count)
	recordTotal, _ := convert.String(p.Total)

	//分页选择框
	/*	var options, selected string
		options = "<option>分页大小</option>"
		for _, v := range []uint{10, 20, 50} {
			if p.Size == v {
				selected = "selected"
			}
			options += fmt.Sprintf(`<option value="%v" %v>%v</option>`, v, selected, v)
		}
		totalCount = fmt.Sprintf(`<li><select id="pagesizeSelect" class="form-control ib w100" style="width:100px">%v</select></li>`, options)
	*/
	totalCount += fmt.Sprintf(`<li class="disabled"><a href="javascript:">共%s页 %s条记录</a></li>`, pageCount, recordTotal)

	return fmt.Sprintf(`<ul class="pagination">%s%s%s%s%s%s</ul>`,
		totalCount, firstLink, prevLink, strings.Join(pageLinks, ""), nextLink, lastLink)
}

//pageURL 生成分页url
func (p *PageInfo) pageURL(c *gin.Context, page string) string {
	//基于当前url新建一个url对象
	//u, _ := url.Parse(p.Request.URL.String())
	u, _ := url.Parse(c.Request.URL.String())
	q := u.Query()
	q.Set("page", page)
	u.RawQuery = q.Encode()
	return u.String()
}
