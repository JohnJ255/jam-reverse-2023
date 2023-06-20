package loader

import "reverse-jam-2023/helper"

type ImageResource struct {
	Filenames []string
	Rotation  helper.Radian
}

func (f *ImageResource) GetFileNames() []string {
	return f.Filenames
}

func (f *ImageResource) GetBaseAngle() helper.Radian {
	return f.Rotation
}
