package Domain

type DirectionType string

var Direction = struct {
	N DirectionType
	E DirectionType
	W DirectionType
	S DirectionType
}{
	N: "N",
	E: "E",
	W: "W",
	S: "S",
}
