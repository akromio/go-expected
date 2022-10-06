// Package expected provides an assertion library.
//
// You can use it as follows:
//
//  1. You must create an expected object at the end of the test:
//
//     expected := expected(t)
//
//  2. You can perform the assertions with this expected such as, for example:
//
//     expected(value).EqualTo(anotherValue)
//     expected(value).GreaterThan(another).LessThan(another)
package expected
