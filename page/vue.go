package page

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//返回标准vue所需分页结构
func Vue(c *gin.Context, list interface{}, query *gorm.DB) (gin.H, error) {
	res := gin.H{
		"page":     1,
		"pageSize": 10,
		"total":    0,
	}
	var page uint = 1
	var pagesize uint = 10
	var totalCount int64 = 0
	if c != nil {
		page = GetPage(c)
		pagesize = GetPageSize(c)
	}

	if err := query.Model(list).Count(&totalCount).Error; err != nil {
		return res, err
	}
	res["total"] = totalCount
	p := NewPage(page, pagesize, uint(totalCount))
	if err := query.Limit(int(p.Limit)).Offset(int(p.Offset)).Find(list).Error; err != nil {
		return nil, err
	}
	return res, nil
}
