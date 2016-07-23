package gomath

const INITIAL_GUESS = 10.0
const ACCEPTABLE_MARGIN = 0.000000000001

/**
 * Computes the square root of a number using the Newton formula.
 */

func Sqrt(value float64) (x float64) {
	x = INITIAL_GUESS                // The first step of the Newton's formula is to define the guess for the square root.
	valueInAcceptableMargin := false // NOTE: we define local var with the := notation.

	// After picking a guess to the sqrt, let's iterate until it is good enough.
	// NOTE: The while in golang is a for with only the condition and no semicolons.
	for !valueInAcceptableMargin {

		// Let's calculate the next value of x by using Newton's formula: f(x1) = x - f(x0)/f'(x0)
		// Where f(x0) = x2 - value and f'(x0) = 2 * x
		x = x - ((x*x - value) / (2 * x))

		// Let's stop when it is good enough, in otherwords when x2 approximates to the value below the acceptable
		// margin
		valueInAcceptableMargin = (x*x-value <= ACCEPTABLE_MARGIN)
	}

	return
}
