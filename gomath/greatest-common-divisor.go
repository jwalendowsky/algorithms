package gomath

/**
 * Computes the greatest common denominator between two numbers using Euclid's
 * algorithm.
 */

// NOTE: Go functions can return more than one result.
// NOTE: We can use return values.
// NOTE: We will only export the function starting with the uppercase.
func Gcd(x int, y int) (gcd int) {

  // NOTE: Whenever we explicit a return value, we don't need to declare it,
  // since it was previosly declared.

  if (y ==0 ) {
    // gcd (a,0) = 0
    gcd = x
  } else {
    // gcd (a, b) = gcd (b, a mod b)
    gcd = Gcd(y, x % y)
  }

  // returning the name variable gcd. This is good to return more than one
  // thing.
  return
}
