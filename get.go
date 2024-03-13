package ext2mime

func Get(fileExtension string) (string, bool) {
	return Default.Get(fileExtension)
}
