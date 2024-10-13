package common

type GF256 struct {
    Exp 		[]int  // Exponentiation table
    Log 		[]int  // Logarithm table
    Size 		int   // Size of the field (256 for GF(256))
}

type Polynomial struct {
    Coefficients []int // Coefficients of the polynomial (in GF(256))
    Field        *GF256 // The Galois field for arithmetic
}

type ReedSolomonEncoder struct {
    N      		int       // Total number of symbols (data + parity)
    K      		int       // Number of data symbols
    Field  		*GF256    // Galois field
    Generator	*Polynomial // Generator polynomial
}

type ReedSolomonDecoder struct {
    N      int       // Total number of symbols
    K      int       // Number of data symbols
    Field  *GF256    // Galois field
}
