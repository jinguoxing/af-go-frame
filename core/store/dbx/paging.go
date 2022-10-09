package dbx


type PaginationResult struct {

    // 页码
    Offset int `json:"offset"`
    // 每页的条数
    Limit int `json:"limit"`
    //总行数
    TotalCount int `json:"total_count"`
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







