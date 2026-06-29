package eval

import (
	"fmt"
	"github.com/Kuniwak/name/strokes"
)

type Rank byte

const (
	Unknown     Rank = 255
	DaiDaiKichi Rank = 4
	DaiKichi    Rank = 3
	Kichi       Rank = 2
	Kyo         Rank = 1
	DaiKyo      Rank = 0
)

func (r Rank) String() string {
	switch r {
	case DaiDaiKichi:
		return "大大吉"
	case DaiKichi:
		return "大吉"
	case Kichi:
		return "吉"
	case Kyo:
		return "凶"
	case DaiKyo:
		return "大凶"
	default:
		return "不明"
	}
}

func StrokesToRank(strokes byte) Rank {
	switch strokes {
	case 1:
		return DaiKichi
	case 2:
		return DaiKyo
	case 3:
		return DaiKichi
	case 4:
		return DaiKyo
	case 5:
		return DaiKichi
	case 6:
		return DaiKichi
	case 7:
		return Kichi
	case 8:
		return Kichi
	case 9:
		return DaiKyo
	case 10:
		return DaiKyo
	case 11:
		return DaiKichi
	case 12:
		return DaiKyo
	case 13:
		return DaiKichi
	case 14:
		return Kyo
	case 15:
		return DaiDaiKichi
	case 16:
		return DaiKichi
	case 17:
		return Kichi
	case 18:
		return Kichi
	case 19:
		return DaiKyo
	case 20:
		return DaiKyo
	case 21:
		return DaiKichi
	case 22:
		return Kyo
	case 23:
		return DaiKichi
	case 24:
		return DaiDaiKichi
	case 25:
		return Kichi
	case 26:
		return Kichi
	case 27:
		return Kyo
	case 28:
		return Kyo
	case 29:
		return DaiKichi
	case 30:
		return Kyo
	case 31:
		return DaiDaiKichi
	case 32:
		return DaiKichi
	case 33:
		return DaiKichi
	case 34:
		return DaiKyo
	case 35:
		return DaiKichi
	case 36:
		return DaiKyo
	case 37:
		return DaiKichi
	case 38:
		return Kichi
	case 39:
		return DaiKichi
	case 40:
		return Kyo
	case 41:
		return DaiKichi
	case 42:
		return Kyo
	case 43:
		return Kyo
	case 44:
		return DaiKyo
	case 45:
		return DaiKichi
	case 46:
		return Kyo
	case 47:
		return DaiKichi
	case 48:
		return Kichi
	case 49:
		return Kyo
	case 50:
		return Kyo
	case 51:
		return Kyo
	case 52:
		return DaiKichi
	case 53:
		return Kyo
	case 54:
		return DaiKyo
	case 55:
		return Kyo
	case 56:
		return Kyo
	case 57:
		return Kichi
	case 58:
		return Kichi
	case 59:
		return Kyo
	case 60:
		return DaiKyo
	case 61:
		return DaiKichi
	case 62:
		return DaiKyo
	case 63:
		return DaiKichi
	case 64:
		return DaiKyo
	case 65:
		return DaiKichi
	case 66:
		return DaiKyo
	case 67:
		return DaiKichi
	case 68:
		return DaiKichi
	case 69:
		return DaiKyo
	case 70:
		return DaiKyo
	case 71:
		return Kichi
	case 72:
		return Kyo
	case 73:
		return Kichi
	case 74:
		return Kyo
	case 75:
		return Kichi
	case 76:
		return DaiKyo
	case 77:
		return Kichi
	case 78:
		return Kichi
	case 79:
		return DaiKyo
	case 80:
		return DaiKyo
	case 81:
		return DaiKichi
	case 82:
		return DaiKichi
	case 83:
		return DaiKyo
	case 84:
		return DaiKichi
	case 85:
		return DaiKyo
	case 86:
		return DaiKichi
	case 87:
		return DaiKichi
	case 88:
		return Kichi
	case 89:
		return Kichi
	case 90:
		return DaiKyo
	case 91:
		return DaiKyo
	case 92:
		return DaiKichi
	case 93:
		return DaiKyo
	case 94:
		return DaiKichi
	case 95:
		return Kyo
	case 96:
		return DaiDaiKichi
	case 97:
		return DaiKichi
	case 98:
		return Kichi
	case 99:
		return Kichi
	case 100:
		return DaiKyo
	}
	return Unknown
}

type Result struct {
	Tenkaku Rank
	Jinkaku Rank
	Chikaku Rank
	Gaikaku Rank
	Sokaku  Rank
}

func (r Result) HasUnknown() bool {
	return r.Tenkaku == Unknown || r.Jinkaku == Unknown || r.Chikaku == Unknown || r.Gaikaku == Unknown || r.Sokaku == Unknown
}

func (r Result) Total() byte {
	if r.HasUnknown() {
		return 0
	}
	return byte(r.Tenkaku) + byte(r.Jinkaku) + byte(r.Chikaku) + byte(r.Gaikaku) + byte(r.Sokaku)
}

func (r Result) String() string {
	return fmt.Sprintf("Result{Tenkaku: %s, Jinkaku: %s, Chikaku: %s, Gaikaku: %s, Sokaku: %s}", r.Tenkaku.String(), r.Jinkaku.String(), r.Chikaku.String(), r.Gaikaku.String(), r.Sokaku.String())
}

func Evaluate(familyName, givenName []rune, strokesFunc strokes.Func) (Result, error) {
	tenkakuStrokes, err := Tenkaku(familyName, strokesFunc)
	if err != nil {
		return Result{}, err
	}
	tenkaku := StrokesToRank(tenkakuStrokes)

	jinkakuStrokes, err := Jinkaku(familyName, givenName, strokesFunc)
	if err != nil {
		return Result{}, err
	}
	jinkaku := StrokesToRank(jinkakuStrokes)

	chikakuStrokes, err := Chikaku(givenName, strokesFunc)
	if err != nil {
		return Result{}, err
	}
	chikaku := StrokesToRank(chikakuStrokes)

	gaikakuStrokes, err := Gaikaku(familyName, givenName, strokesFunc)
	if err != nil {
		return Result{}, err
	}
	gaikaku := StrokesToRank(gaikakuStrokes)

	sokakuStrokes, err := Sokaku(familyName, givenName, strokesFunc)
	if err != nil {
		return Result{}, err
	}
	sokaku := StrokesToRank(sokakuStrokes)

	return Result{
		Tenkaku: tenkaku,
		Jinkaku: jinkaku,
		Chikaku: chikaku,
		Gaikaku: gaikaku,
		Sokaku:  sokaku,
	}, nil
}

func Tenkaku(familyName []rune, strokesFunc strokes.Func) (byte, error) {
	return strokes.Sum(familyName, strokesFunc)
}

func Jinkaku(familyName, givenName []rune, strokesFunc strokes.Func) (byte, error) {
	return strokes.Add(familyName[len(familyName)-1], givenName[0], strokesFunc)
}

func Chikaku(givenName []rune, strokesFunc strokes.Func) (byte, error) {
	return strokes.Sum(givenName, strokesFunc)
}

func Gaikaku(familyName, givenName []rune, strokesFunc strokes.Func) (byte, error) {
	c1, err := strokesFunc(familyName[0])
	if err != nil {
		return 0, err
	}

	c2, err := strokesFunc(givenName[len(givenName)-1])
	if err != nil {
		return 0, err
	}

	var n1 byte
	if len(familyName) == 1 {
		n1 = c1 + 1
	} else {
		n1 = c1
	}
	var n2 byte
	if len(givenName) == 1 {
		n2 = c2 + 1
	} else {
		n2 = c2
	}
	return n1 + n2, nil
}

func Sokaku(familyName, givenName []rune, strokesFunc strokes.Func) (byte, error) {
	c1, err := strokes.Sum(familyName, strokesFunc)
	if err != nil {
		return 0, err
	}
	c2, err := strokes.Sum(givenName, strokesFunc)
	if err != nil {
		return 0, err
	}
	return c1 + c2, nil
}
