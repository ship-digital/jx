// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	versioned4 "github.com/knative/serving/pkg/client/clientset/versioned"
	"github.com/petergtz/pegomock"
	"reflect"
)

func AnyVersioned4Interface() versioned4.Interface {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(versioned4.Interface))(nil)).Elem()))
	var nullValue versioned4.Interface
	return nullValue
}

func EqVersioned4Interface(value versioned4.Interface) versioned4.Interface {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue versioned4.Interface
	return nullValue
}
