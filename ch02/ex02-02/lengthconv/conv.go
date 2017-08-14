package lengthconv

func MToF(m Meter) Foot { return Foot(m / 0.3048) }

func FToM(f Foot) Meter { return Meter(f * 0.3048) }
