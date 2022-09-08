package errorx

import "bytes"

type (

    errorArray []error

    BatchError struct {
        errs errorArray
    }

)


func(be *BatchError) Add(errs ...error){

    for _, err := range errs {
        if err != nil {
            be.errs = append(be.errs, err)
        }
    }

}

// Err returns an error that represents all errors.
func(be *BatchError) Err() error{

    switch len(be.errs) {
    case 0:
        return nil
    case 1:
        return be.errs[0]
    default:
        return be.errs
    }

}

func(be *BatchError) NotNil() bool {

    return len(be.errs) > 0
}


func (ea errorArray) Error() string{

    var buf bytes.Buffer

    for i:= range ea {
        if i > 0 {
            buf.WriteByte('\n')
        }
        buf.WriteString(ea[i].Error())
    }

    return buf.String()
}