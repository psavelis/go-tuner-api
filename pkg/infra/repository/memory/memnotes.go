package memory

import (
	"fmt"

	"github.com/psavelis/go-tuner-api/pkg/domain/entity"
)


type InMemoryNotesRepository struct {
	notes map[int]entity.StandardNote
}

func NewInMemoryNotesRepository() *InMemoryNotesRepository {
	return &InMemoryNotesRepository{
		notes: map[int]entity.StandardNote{},
	}
}

func (r *InMemoryNotesRepository) Find(key int) (entity.StandardNote, error) {
	name := pianoMap[key]

	if name == "" {
		return entity.StandardNote{}, entity.ErrNoteNotFound
	}

	res := entity.StandardNote{ID: fmt.Sprint(key), Name: name, KeyNumber: key}

	return res, nil
}

var pianoMap = map[int]string{
	108: "B8",
	107: "A♯8/B♭8",
	106: "A8",
	105: "G♯8/A♭8",
	104: "G8",
	103: "F♯8/G♭8",
	102: "F8",
	101: "E8",
	100: "D♯8/E♭8",
	99:  "D8",
	98:  "C♯8/D♭8",
	88:  "C8 Eighth octave",
	87:  "B7",
	86:  "A♯7/B♭7",
	85:  "A7",
	84:  "G♯7/A♭7",
	83:  "G7",
	82:  "F♯7/G♭7",
	81:  "F7",
	80:  "E7",
	79:  "D♯7/E♭7",
	78:  "D7",
	77:  "C♯7/D♭7",
	76:  "C7 Double high C",
	75:  "B6",
	74:  "A♯6/B♭6",
	73:  "A6",
	72:  "G♯6/A♭6",
	71:  "G6",
	70:  "F♯6/G♭6",
	69:  "F6",
	68:  "E6",
	67:  "D♯6/E♭6",
	66:  "D6",
	65:  "C♯6/D♭6",
	64:  "C6 Soprano C (High C)",
	63:  "B5",
	62:  "A♯5/B♭5",
	61:  "A5",
	60:  "G♯5/A♭5",
	59:  "G5",
	58:  "F♯5/G♭5",
	57:  "F5",
	56:  "E5",
	55:  "D♯5/E♭5",
	54:  "D5",
	53:  "C♯5/D♭5",
	52:  "C5 Tenor C",
	51:  "B4",
	50:  "A♯4/B♭4",
	49:  "A4",
	48:  "G♯4/A♭4",
	47:  "G4",
	46:  "F♯4/G♭4",
	45:  "F4",
	44:  "E4",
	43:  "D♯4/E♭4",
	42:  "D4",
	41:  "C♯4/D♭4",
	40:  "C4 Middle C",
	39:  "B3",
	38:  "A♯3/B♭3",
	37:  "A3",
	36:  "G♯3/A♭3",
	35:  "G3",
	34:  "F♯3/G♭3",
	33:  "F3",
	32:  "E3",
	31:  "D♯3/E♭3",
	30:  "D3",
	29:  "C♯3/D♭3",
	28:  "C3",
	27:  "B2",
	26:  "A♯2/B♭2",
	25:  "A2",
	24:  "G♯2/A♭2",
	23:  "G",
	22:  "F♯2/G♭",
	21:  "F",
	20:  "E",
	19:  "D♯2/E♭",
	18:  "D",
	17:  "C♯2/D♭",
	16:  "C2 Deep",
	15:  "B",
	14:  "A♯1/B♭",
	13:  "A",
	12:  "G♯1/A♭",
	11:  "G",
	10:  "F♯1/G♭",
	9:   "F",
	8:   "E",
	7:   "D♯1/E♭",
	6:   "D",
	5:   "C♯1/D♭1",
	4:   "C1 Pedal C",
	3:   "B0",
	2:   "A♯0/B♭0",
	1:   "A0",
	97:  "G♯0/A♭0",
	96:  "G0",
	95:  "F♯0/G♭0",
	94:  "F0",
	93:  "E0",
	92:  "D♯0/E♭0",
	91:  "D0",
	90:  "C♯0/D♭0",
	89:  "C0 Double Pedal C",
}
