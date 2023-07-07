package xgo

import "math"

func mapValue(v, oMin, oMax, tMin, tMax int) int {
	fv := math.Min(float64(v), float64(oMax))
	fv = math.Max(fv, float64(oMin))
	fv -= float64(oMin)
	r1 := math.Abs(float64(oMax) - float64(oMin))
	r2 := math.Abs(float64(tMax) - float64(tMin))
	fv /= (r1 / r2)
	if tMin < tMax {
		return tMin + int(fv)
	}
	return tMin - int(fv)
}

func generateMessage(v ...byte) []byte {
	b := make([]byte, 6+len(v))
	b[0] = 0x55
	b[1] = 0x00
	b[2] = byte(len(b))
	chksum := int(b[2])
	for i, vl := range v {
		b[3+i] = vl
		chksum += int(vl)
	}
	chkId := len(b) - 3
	b[chkId] = ^byte(chksum & 0xff)
	b[chkId+1] = 0x00
	b[chkId+2] = 0xaa
	return b
}

func Rotate(d RotateEnum, spd int) []byte {
	if d == RotateLeft {
		return generateMessage(0x00, 0x32, byte(mapValue(spd, 0, 100, 128, 255)))
	}
	return generateMessage(0x00, 0x32, byte(mapValue(spd, 0, 100, 128, 0)))

}

func BodyHeight(hgt int) []byte {
	return generateMessage(0x00, 0x35, byte(mapValue(hgt, 0, 100, 0, 255)))
}

func Move(d DirectionEnum, spd int) []byte {
	switch d {
	case DirectionBackward:
		return generateMessage(0x00, 0x30, byte(mapValue(spd, 0, 100, 128, 0)))
	case DirectionLeft:
		return generateMessage(0x00, 0x3, byte(mapValue(spd, 0, 100, 128, 0)))
	case DirectionRight:
		return generateMessage(0x00, 0x30, byte(mapValue(spd, 0, 100, 128, 255)))
	default:
		return generateMessage(0x00, 0x30, byte(mapValue(spd, 0, 100, 128, 255)))
	}
}
