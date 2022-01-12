package screwmatrix

type MoveWay int

const (
	Top MoveWay = iota
	Right
	Down
	Left
)

var moveWayArr []MoveWay

func (s *SMatrix) initMoveWays() {
	if s.CWise == RightClockWise {
		moveWayArr = []MoveWay{Right, Down, Left, Top}
	} else {
		moveWayArr = []MoveWay{Down, Right, Top, Left}
	}
}

func (s *SMatrix) move() {
	var row, col int = 0, 0
	var currentMoveWayIndex int = 0
	var currentVal int = 1
	for currentVal <= s.endVal {
		if currentVal == 1 {
			s.Matrix[row][col] = currentVal
			currentVal++
			continue
		}

		way := moveWayArr[currentMoveWayIndex]
		tmpRow, tmpCol := s.getNextRowAndCol(row, col, way)
		if s.checkMoveLegit(tmpRow, tmpCol) {
			row = tmpRow
			col = tmpCol
			s.Matrix[row][col] = currentVal
			currentVal++
		} else {
			if currentMoveWayIndex == len(moveWayArr)-1 { // 走完一個循環
				currentMoveWayIndex = 0 // 重置方向
			} else {
				currentMoveWayIndex++
			}
		}
	}
}

func (s *SMatrix) getNextRowAndCol(row, col int, moveWay MoveWay) (int, int) {
	var nextRow, nextCol int = 0, 0
	switch moveWay {
	case Top:
		nextRow = row - 1
		nextCol = col
		break
	case Right:
		nextRow = row
		nextCol = col + 1
		break
	case Down:
		nextRow = row + 1
		nextCol = col
		break
	case Left:
		nextRow = row
		nextCol = col - 1
		break
	default:
		nextRow = 0
		nextCol = 0
	}
	return nextRow, nextCol
}

func (s *SMatrix) checkMoveLegit(row, col int) bool {
	var maxPos = s.Capacity - 1
	var minPos = 0
	if row > maxPos || col > maxPos {
		return false
	}

	if row < minPos || col < minPos {
		return false
	}

	val := s.Matrix[row][col]
	if val != 0 {
		return false
	}
	return true
}
