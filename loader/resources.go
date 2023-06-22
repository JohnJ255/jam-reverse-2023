package loader

import (
	"reverse-jam-2023/framework"
	"reverse-jam-2023/models"
)

var CarFileNames = map[framework.ControlType]*ImageResource{
	framework.Player: {
		Rotation: framework.AngleTop,
		Filenames: []string{
			"Car_1_01.png",
			"Car_1_02.png",
			"Car_1_03.png",
			"Car_1_04.png",
			"Car_1_05.png",
		},
	},
	framework.Computer: {
		Rotation: framework.AngleTop,
		Filenames: []string{
			"Car_3_01.png",
			"Car_3_02.png",
			"Car_3_03.png",
			"Car_3_04.png",
			"Car_3_05.png",
		},
	},
}

var TrailerFileNames = map[models.TrailerType]*ImageResource{
	models.TrailerTypeCart: {
		Rotation: framework.AngleRight,
		Filenames: []string{
			"Trailer_1_01.png",
		},
	},
	models.TrailerTypeTrailer: {
		Rotation: framework.AngleRight,
		Filenames: []string{
			"Trailer_2_01.png",
		},
	},
}
