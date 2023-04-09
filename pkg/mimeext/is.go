package mimeext

func IsVideo(mime string) bool {
	return mime == MP4 || mime == MOV
}

func IsImage(mime string) bool {
	return mime == JPG || mime == PNG
}
