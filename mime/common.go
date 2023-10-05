package mime

const (
	// Plain content must not contain binary data.
	// Plain is the default type for text content.
	// The content should be human-readable.
	Plain = "text/plain"
	Css   = "text/css"
	Csv   = "text/csv"
	Html  = "text/html"
	Js    = "text/javascript"

	// OctetStream indicates arbitrary binary data.
	// OctetStream is the default type for non-text content.
	// An unknown file type should use this type.
	OctetStream = "application/octet-stream"
	Unknown     = OctetStream
	Json        = "application/json"
	JsonLd      = "application/ld+json"
	Pdf         = "application/pdf"

	Avif      = "image/avif"
	Bmp       = "image/bmp"
	Gif       = "image/gif"
	Jpeg      = "image/jpeg"
	Png       = "image/png"
	Svg       = "image/svg+xml"
	WebmImage = "image/webm"

	Mp3       = "audio/mpeg"
	OggAudio  = "audio/ogg"
	Wav       = "audio/wav"
	WebmAudio = "audio/webm"

	Mp4       = "video/mp4"
	Mpeg      = "video/mpeg"
	OggVideo  = "video/ogg"
	WebmVideo = "video/webm"
)
