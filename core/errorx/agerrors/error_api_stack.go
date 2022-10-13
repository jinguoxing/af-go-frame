package agerrors

// Cause returns the root cause error of `err`.
func Cause(err error) error {
    if err == nil {
        return nil
    }
    if e, ok := err.(ICause); ok {
        return e.Cause()
    }
    if e, ok := err.(IUnwrap); ok {
        return Cause(e.Unwrap())
    }
    return err
}

// Stack returns the stack callers as string.
// It returns the error string directly if the `err` does not support stacks.
func Stack(err error) string {
    if err == nil {
        return ""
    }
    if e, ok := err.(IStack); ok {
        return e.Stack()
    }
    return err.Error()
}

// Current creates and returns the current level error.
// It returns nil if current level error is nil.
func Current(err error) error {
    if err == nil {
        return nil
    }
    if e, ok := err.(ICurrent); ok {
        return e.Current()
    }
    return err
}

// Unwrap returns the next level error.
// It returns nil if current level error or the next level error is nil.
func Unwrap(err error) error {
    if err == nil {
        return nil
    }
    if e, ok := err.(IUnwrap); ok {
        return e.Unwrap()
    }
    return nil
}

// HasStack checks and reports whether `err` implemented interface `gerror.IStack`.
func HasStack(err error) bool {
    _, ok := err.(IStack)
    return ok
}

// Equal reports whether current error `err` equals to error `target`.
// Please note that, in default comparison logic for `Error`,
// the errors are considered the same if both the `code` and `text` of them are the same.
func Equal(err, target error) bool {
    if err == target {
        return true
    }
    if e, ok := err.(IEqual); ok {
        return e.Equal(target)
    }
    if e, ok := target.(IEqual); ok {
        return e.Equal(err)
    }
    return false
}

// Is reports whether current error `err` has error `target` in its chaining errors.
// It is just for implements for stdlib errors.Is from Go version 1.17.
func Is(err, target error) bool {
    if e, ok := err.(IIs); ok {
        return e.Is(target)
    }
    return false
}

// HasError is alias of Is, which more easily understanding semantics.
func HasError(err, target error) bool {
    return Is(err, target)
}
