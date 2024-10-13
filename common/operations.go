package common

/*
Operations for GF256
*/

func (gf *GF256) Add(a,b int)int{
	return a^b
}

func (gf *GF256) Mul(a,b int)int{
	if a == 0 || b == 0 {
		return 0
	}
	log_a := gf.Log[a]
	log_b := gf.Log[b]
	log_result := (log_a + log_b) % 255
	return gf.Exp[log_result]
}

func (gf * GF256) Div(a,b int)int{
	if b == 0 {
        panic("Division by zero in GF(256)")
    }
    if a == 0 {
        return 0 // Dividing zero yields zero
    }
    log_a := gf.Log[a]
    log_b := gf.Log[b]
    log_result := (log_a - log_b + 255) % 255 // Ensure the result is positive
    return gf.Exp[log_result]
}

func (gf *GF256) Pow(a,b int)int{
	if a == 0{
		return 0
	}
	log_a := gf.Log[a]
	log_result := (log_a * b) % 255
	return gf.Exp[log_result]
}

/*
Operations for Polynomial 
*/

func (p *Polynomial) Add(other *Polynomial) *Polynomial {
	var result Polynomial	

	maxLen := max(len(p.Coefficients),len(other.Coefficients))

	resultCoeff := make([]int,maxLen)

	for i := 0; i < maxLen; i++ {
		var a,b int
		if i<len(p.Coefficients) {
			a = p.Coefficients[i]
		}
		if i<len(other.Coefficients) {
			b = other.Coefficients[i]
		}
		resultCoeff[i] = result.Field.Add(a,b)
	}

	result.Coefficients = resultCoeff

	return &result
}

func (p *Polynomial) Mul(other *Polynomial) *Polynomial {
	var result Polynomial
	resultLen := len(p.Coefficients) + len(other.Coefficients) - 1
	resultCoeff := make([]int,resultLen)

	for i := 0; i < len(p.Coefficients); i++ {
        for j := 0; j < len(other.Coefficients); j++ {
            // Multiply the coefficients using GF(256) multiplication
            resultCoeff[i+j] ^= result.Field.Mul(p.Coefficients[i], other.Coefficients[j])
        }
    }
	result.Coefficients = resultCoeff
	return &result
}

/*
returns quotient and remainder
*/
func (p *Polynomial) Div(other *Polynomial)(*Polynomial,*Polynomial) {
	var quotient,remainder Polynomial

	quotientCoef := make([]int, len(p.Coefficients))
    remainderCoef := make([]int, len(p.Coefficients))
    copy(remainderCoef, p.Coefficients)

    // Perform the division
    for len(remainderCoef) >= len(other.Coefficients) {
        // Get the leading term of remainder and divisor
        leadingTermR := remainderCoef[0] // First element of remainder
        leadingTermD := other.Coefficients[0] // First element of divisor

        // Divide leading terms (GF division)
        quotientTerm := quotient.Field.Div(leadingTermR, leadingTermD)
        
        // Create a new quotient term polynomial
        termPoly := make([]int, len(remainderCoef))
        termPoly[0] = quotientTerm

        // Subtract (termPoly * other) from the remainder
        for i := 0; i < len(other.Coefficients); i++ {
            remainderCoef[i] ^= remainder.Field.Mul(quotientTerm, other.Coefficients[i])
        }
        
        // Update the quotient polynomial
        quotientCoef[len(remainderCoef)-len(other.Coefficients)] = quotientTerm
        
        // Remove leading zeroes from the remainder
        remainderCoef = trimLeadingZeros(remainderCoef)
    }
	quotient.Coefficients = quotientCoef 
	remainder.Coefficients = remainderCoef

	return &quotient,&remainder
}

