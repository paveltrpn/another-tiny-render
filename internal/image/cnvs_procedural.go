package image

// Fill already created canvas with checker texture
// chkr_size - size of checker square in pixels
func (cnvs *SCanvas) DrawChecker(chkr_size int) {
	var (
		value uint8
	)

	for row := 0; row < cnvs.width; row++ {
		for col := 0; col < cnvs.height; col++ {
			X := ((row / chkr_size) % 2) == 0
			Y := ((col / chkr_size) % 2) == 0
			// X xor Y in next line!
			if (X || Y) && !(X && Y) {
				value = 255
			} else {
				value = 0
			}

			cnvs.data[((row*3)*cnvs.height+col*3)+0] = value
			cnvs.data[((row*3)*cnvs.height+col*3)+1] = value
			cnvs.data[((row*3)*cnvs.height+col*3)+2] = value
		}
	}
}
