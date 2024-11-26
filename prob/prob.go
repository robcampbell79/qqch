package prob

import(
	"fmt"
	"math/big"
)

func LottoProb(b int, rng int, pow int) {
	nB := big.NewInt(int64(b))
	nR := big.NewInt(int64(rng))
	nB2 := big.NewInt(int64(b))
	nR2 := big.NewInt(int64(rng))
	//powBall := big.NewInt(1)
	pR := big.NewInt(int64(pow))

	nb := facto(nB)
	nr := facto(nR)
	var nrnb = new(big.Int)
	nrnb.Sub(nR2, nB2)
	fnrnb := facto(nrnb)
	var nftb = new(big.Int)
	nftb.Mul(nb, fnrnb)
	var nftr = new(big.Int)
	nftr.Div(nr, nftb)
	var odds = new(big.Int)
	odds.Mul(nftr, pR)

	fmt.Println("odds ", odds)

}

func facto(number *big.Int) (result*big.Int) {
	one := big.NewInt(1)
	on := big.NewInt(0)

	if number.Cmp(on) == -1 {
		result = one
	}
	if number.Cmp(on) == 0 {
		result = one
	} else {
		result = new(big.Int)
		result.Set(number)
		result.Mul(result, facto(number.Sub(number, one)))
	}

	return
}