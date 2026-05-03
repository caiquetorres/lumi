package vm

type tag byte

const (
	tagInt tag = iota
	tagBool
	tagPtr
	tagString
	tagFn
	tagNil
)

const (
	tagShift = 61
	tagMask  = uint64(0b111) << tagShift
	valMask  = ^tagMask
)

func getTag(v uint64) tag {
	return tag(uint64(v) >> tagShift)
}

func encodeInt(v int64) uint64 {
	return (uint64(tagInt) << tagShift) | (uint64(v) & valMask)
}

func encodeBool(b bool) uint64 {
	var v uint64
	if b {
		v = 1
	}
	return (uint64(tagBool) << tagShift) | v
}

func encodeString(addr int64) uint64 {
	return uint64((uint64(tagString) << tagShift) |
		(uint64(addr) & valMask))
}

func encodeFn(addr int64) uint64 {
	return uint64((uint64(tagFn) << tagShift) |
		(uint64(addr) & valMask))
}

func decodeInt(v uint64) int64 {
	payload := uint64(v) & valMask

	if payload&(1<<60) != 0 {
		return int64(payload | tagMask)
	}

	return int64(payload)
}

func decodeBool(v uint64) bool {
	return (uint64(v) & 1) == 1
}

func decodeString(v uint64) int64 {
	return int64(v & valMask)
}

func decodeFn(v uint64) int64 {
	return int64(v & valMask)
}
