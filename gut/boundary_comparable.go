package gut

type BoundaryComparable interface {
	Less(other BoundaryComparable) bool
	LessOrEqual(other BoundaryComparable) bool
	Greater(other BoundaryComparable) bool
	GreaterOrEqual(other BoundaryComparable) bool
}

// int64
type BoundaryInt int64

func (m BoundaryInt) Less(other BoundaryComparable) bool {
	return m < other.(BoundaryInt)
}

func (m BoundaryInt) LessOrEqual(other BoundaryComparable) bool {
	return m <= other.(BoundaryInt)
}

func (m BoundaryInt) Greater(other BoundaryComparable) bool {
	return m > other.(BoundaryInt)
}

func (m BoundaryInt) GreaterOrEqual(other BoundaryComparable) bool {
	return m >= other.(BoundaryInt)
}

// uint64
type BoundaryUInt uint64

func (m BoundaryUInt) Less(other BoundaryComparable) bool {
	return m < other.(BoundaryUInt)
}

func (m BoundaryUInt) LessOrEqual(other BoundaryComparable) bool {
	return m <= other.(BoundaryUInt)
}

func (m BoundaryUInt) Greater(other BoundaryComparable) bool {
	return m > other.(BoundaryUInt)
}

func (m BoundaryUInt) GreaterOrEqual(other BoundaryComparable) bool {
	return m >= other.(BoundaryUInt)
}

// float64
type BoundaryFloat float64

func (m BoundaryFloat) Less(other BoundaryComparable) bool {
	return m < other.(BoundaryFloat)
}

func (m BoundaryFloat) LessOrEqual(other BoundaryComparable) bool {
	return m <= other.(BoundaryFloat)
}

func (m BoundaryFloat) Greater(other BoundaryComparable) bool {
	return m > other.(BoundaryFloat)
}

func (m BoundaryFloat) GreaterOrEqual(other BoundaryComparable) bool {
	return m >= other.(BoundaryFloat)
}

// string
type BoundaryString string

func (m BoundaryString) Less(other BoundaryComparable) bool {
	return m < other.(BoundaryString)
}

func (m BoundaryString) LessOrEqual(other BoundaryComparable) bool {
	return m <= other.(BoundaryString)
}

func (m BoundaryString) Greater(other BoundaryComparable) bool {
	return m > other.(BoundaryString)
}

func (m BoundaryString) GreaterOrEqual(other BoundaryComparable) bool {
	return m >= other.(BoundaryString)
}
