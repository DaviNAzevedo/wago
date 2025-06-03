package webhook

type Referral struct {
	SourceURL    string  `json:"source_url"`              // URL da fonte (ex: URL do anúncio)
	SourceType   string  `json:"source_type"`             // Tipo da fonte (ex: "ad", "post")
	SourceID     string  `json:"source_id"`               // ID da fonte (ex: ID do anúncio)
	Headline     *string `json:"headline,omitempty"`      // Título do anúncio/post
	Body         *string `json:"body,omitempty"`          // Corpo do anúncio/post
	MediaType    *string `json:"media_type,omitempty"`    // Tipo de mídia ("image", "video")
	ImageURL     *string `json:"image_url,omitempty"`     // URL da imagem (se media_type="image")
	VideoURL     *string `json:"video_url,omitempty"`     // URL do vídeo (se media_type="video")
	ThumbnailURL *string `json:"thumbnail_url,omitempty"` // URL da miniatura do vídeo
}
