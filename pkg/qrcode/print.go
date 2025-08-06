package qrcode

import (
	"bytes"
	"github.com/skip2/go-qrcode"
	"os"
)

const (
	frontColor      = "\033[48;5;0m  \033[0m"
	backgroundColor = "\033[48;5;7m  \033[0m"
	level           = qrcode.Low
)

func Print(content string) error {
	qr, err := qrcode.New(content, level)
	if err != nil {
		return err
	}

	bitmap := qr.Bitmap()
	output := bytes.NewBuffer([]byte{})
	for ir, row := range bitmap {
		lr := len(row)

		if ir == 0 || ir == 1 || ir == 2 ||
			ir == lr-1 || ir == lr-2 || ir == lr-3 {
			continue
		}

		for ic, col := range row {
			lc := len(bitmap)
			if ic == 0 || ic == 1 || ic == 2 ||
				ic == lc-1 || ic == lc-2 || ic == lc-3 {
				continue
			}
			if col {
				output.WriteString(frontColor)
			} else {
				output.WriteString(backgroundColor)
			}
		}
		output.WriteByte('\n')
	}
	if _, err = output.WriteTo(os.Stdout); err != nil {
		return err
	}
	return nil
}
