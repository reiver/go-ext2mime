package ext2mime

func init() {
	// 3d
	Set(".vrml",   "model/vrml")

	// audio
	Set(".mid",    "audio/midi")
	Set(".midi",   "audio/midi")
	Set(".mp3",    "audio/mpeg")
	Set(".oga",    "audio/ogg")
	Set(".wav",    "audio/wav")
	Set(".weba",   "audio/webm")

	// bundle
	Set(".atom",   "application/atom+xml")
	Set(".rss",    "application/rss+xml")
	Set(".ogx",    "application/ogg")
	Set(".tar",    "application/x-tar")

	// compress
	Set(".bz",     "application/x-bzip")
	Set(".bz2",    "application/x-bzip2")
	Set(".gz",     "application/gzip")
	Set(".zip",    "application/zip")

	// data
	Set(".bin",          "application/octet-stream")
	Set(".csv",          "text/csv")
	Set(".ini",          "text/ini")
	Set(".jrd",          "application/jrd+json")
	Set(".json",         "application/json")
	Set(".jsonactivity", "application/activity+json")
	Set(".jsonfeed",     "application/feed+json")
	Set(".jsonld",       "application/ld+json")
	Set(".rdf",          "application/rdf+xml")
	Set(".uris",         "text/uri-list")

	// font
	Set(".otf",    "font/otf")
	Set(".ttf",    "font/ttf")
	Set(".woff",   "font/woff")
	Set(".woff2",  "font/woff2")

	// image
	Set(".apng",   "image/apng")
	Set(".bmp",    "image/bmp")
	Set(".gif",    "image/gif")
	Set(".ico",    "image/vnd.microsoft.icon")
	Set(".ink",    "application/inkml+xml")
	Set(".inkml",  "application/inkml+xml")
	Set(".jpeg",   "image/jpeg")
	Set(".jpg",    "image/jpeg")
	Set(".qoi",    "image/qoi")
	Set(".png",    "image/png")
	Set(".tif",    "image/tiff")
	Set(".tiff",   "image/tiff")
	Set(".webp",   "image/webp")

	// program
	Set(".js",     "text/javascript")
	Set(".jar",    "application/java-archive")
	Set(".mjs",    "text/javascript")
	Set(".xul",    "application/vnd.mozilla.xul+xml")

	// style
	Set(".css",    "text/css")

	// text
	Set(".ansi",   "text/ansi")
	Set(".epub",   "application/epub+zip")
	Set(".gemini", "text/gemini")
	Set(".gmi",    "text/gemini")
	Set(".gmni",   "text/gemini")
	Set(".gpub",   "application/gpub+zip")
	Set(".htm",    "text/html")
	Set(".html",   "text/html")
	Set(".md",     "text/markdown")
	Set(".pdf",    "application/pdf")
	Set(".ps",     "application/postscript")
	Set(".rtf",    "application/rtf")
	Set(".txt",    "text/plain")
	Set(".xhtml",  "application/xhtml+xml")

	// video
	Set(".avi",    "video/x-msvideo")
	Set(".mp4",    "video/mp4")
	Set(".mpeg",   "video/mpeg")
	Set(".ogv",    "video/ogg")
	Set(".smi",    "application/smil+xml")
	Set(".sml",    "application/smil+xml")
	Set(".smil",   "application/smil+xml")
	Set(".webm",   "video/webm")
}
