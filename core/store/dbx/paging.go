package dbx


type PaginationResult struct {

    Offset int `json:"offset"`  // 页码
    Limit int `json:"limit"` // 每页的条数
    TotalCount int `json:"total_count"` //总数据条数
}


func (p *PaginationResult) GetOffset() int {
    offset := 0
    if p.Offset > 0 {
        offset = (p.Offset - 1) * p.Limit
    }

    return offset
}


func (p *PaginationResult) TotalPage() int {

    if p.TotalCount == 0 || p.Limit == 0 {
        return 0
    }

    totalPage := int(p.TotalCount) / p.Limit
    if int(p.TotalCount)%p.Limit > 0 {
        totalPage = totalPage + 1
    }
    return totalPage
}


type PaginationParam struct {
    Pagination bool `form:"-"`
    OnlyCount  bool `form:"-"`
    Offset    int  `form:"Offset,default=1"`
    Limit   int  `form:"limit,default=20" binding:"max=100"`
}

func (a PaginationParam) GetCurrent() int {
    return a.Offset
}

func (a PaginationParam) GetPageSize() int {
    pageSize := a.Limit
    if a.Limit <= 0 {
        pageSize = 20
    }
    return pageSize
}







