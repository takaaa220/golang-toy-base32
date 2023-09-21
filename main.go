package main

const S = "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567"

func Main(input []byte) string {
	result := []byte{}

	i := 0
	for cursor := 0; i < len(input); cursor = (cursor + 5) % 8 {
		var value byte

		switch cursor {
		case 0, 1, 2, 3:
			// cursorの位置から5ビットを取得

			value = (input[i] >> (3 - cursor)) & 0x1F

			if cursor == 3 {
				i++
			}
		case 4, 5, 6, 7:
			// input[i]の下位ビットとinput[i+1]の上位ビットを結合

			var bottom byte
			if i < len(input)-1 {
				bottom = input[i+1] >> (11 - cursor)
			}

			top := (input[i] & (0xff >> cursor))

			value = (top << (cursor - 3)) | bottom

			i++
		}

		result = append(result, S[value])
	}

	return string(result)
}
