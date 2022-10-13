package ginx

import (
    "strings"

    "github.com/gin-gonic/gin"
    ut "github.com/go-playground/universal-translator"
    val "github.com/go-playground/validator/v10"
)

type ValidError struct {
    Key     string
    Message string
}

type ValidErrors []*ValidError

func (v *ValidError) Error() string {
    return v.Message
}

func (v ValidErrors) Error() string {
    return strings.Join(v.Errors(), ",")
}

func (v ValidErrors) Errors() []string {
    var errs []string
    for _, err := range v {
        errs = append(errs, err.Error())
    }

    return errs
}

//
func BindAndValid(c *gin.Context, v interface{}) (bool, ValidErrors) {
    var errs ValidErrors
    err := c.ShouldBind(v)
    if err != nil {
        v := c.Value("trans")
        trans, _ := v.(ut.Translator)
        verrs, ok := err.(val.ValidationErrors)
        if !ok {
            return false, errs
        }

        vv := removeTopStruct(verrs.Translate(trans))
        for key, value := range vv {
            errs = append(errs, &ValidError{
                Key:     key,
                Message: value,
            })
        }

        return false, errs
    }

    return true, nil
}

// removeTopStruct 去除字段名中的结构体名称标识
// refer from:https://github.com/go-playground/validator/issues/633#issuecomment-654382345
func removeTopStruct(fields map[string]string) map[string]string {
    res := map[string]string{}
    for field, err := range fields {
        res[field[strings.Index(field, ".")+1:]] = err
    }
    return res
}
