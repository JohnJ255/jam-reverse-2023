package game

import "reverse-jam-2023/framework"

type LevelAbout struct {
}

func (l *LevelAbout) Fill(level *LevelManager) {

}

func (l *LevelAbout) GetSize() framework.Size {
	return framework.Size{800, 600}
}
