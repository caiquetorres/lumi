package vm

type tag byte

const (
	tagInt tag = iota
	tagBool
	tagPtr
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

func isInt(v uint64) bool {
	return getTag(v) == tagInt
}

func isBool(v uint64) bool {
	return getTag(v) == tagBool
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
