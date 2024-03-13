package ext2mime

func Set(fileExtension string, mimeType string) error {
	return Default.Set(fileExtension, mimeType)
}
