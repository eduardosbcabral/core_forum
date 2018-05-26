package config

import(
	"reflect"
)

//Generic error handler - WRONG
func ErrorHandler(function interface{}, args ...interface{}) bool {

	v := reflect.ValueOf(function)
    rargs := make([]reflect.Value, len(args))
    for i, a := range args {
        rargs[i] = reflect.ValueOf(a)
    }
    fn := v.Call(rargs)

    errorInterface := reflect.TypeOf((*error)(nil)).Elem()

    for _, a := range fn {
    	if a.Type().Implements(errorInterface) {
    		if !a.IsNil() {
    			return true
    		}
    	}
    }

    return false
}