package chat

import (
	"bytes"
	"errors"
	"image"
	"image/jpeg"
	"image/png"

	"golang.org/x/image/draw"

	// Decodeurs d'entree (l'encodage de sortie est JPEG).
	_ "golang.org/x/image/webp"
	_ "image/gif"
)

// P1 — redimensionnement des images avant encodage base64. Le grand cote est
// borne ; l'image est re-encodee en JPEG. Pur Go (CGO_ENABLED=0).
const (
	visionImageMaxPx    = 1024
	visionJPEGQuality   = 80
	attachThumbMaxPx    = 256
	attachThumbQuality  = 75
	imageResizeDisabled = 0
)

// resizeJPEG decode data, borne le grand cote a maxPx (aucun agrandissement)
// et re-encode en JPEG. Si maxPx <= 0, aucune borne de taille n'est
// appliquee, mais l'image est tout de meme re-encodee.
func resizeJPEG(data []byte, maxPx, quality int) ([]byte, error) {
	src, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	dst := src
	if maxPx > 0 {
		b := src.Bounds()
		w, h := b.Dx(), b.Dy()
		if w > maxPx || h > maxPx {
			scale := float64(maxPx) / float64(w)
			if h > w {
				scale = float64(maxPx) / float64(h)
			}
			nw := int(float64(w)*scale + 0.5)
			nh := int(float64(h)*scale + 0.5)
			if nw < 1 {
				nw = 1
			}
			if nh < 1 {
				nh = 1
			}
			rgba := image.NewRGBA(image.Rect(0, 0, nw, nh))
			draw.CatmullRom.Scale(rgba, rgba.Bounds(), src, b, draw.Over, nil)
			dst = rgba
		}
	}
	var buf bytes.Buffer
	if err := jpeg.Encode(&buf, dst, &jpeg.Options{Quality: quality}); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// decodePNGSize : utilitaire interne (tests) — non exporte.
func decodePNGSize(data []byte) (int, int, error) {
	cfg, _, err := image.DecodeConfig(bytes.NewReader(data))
	if err != nil {
		return 0, 0, err
	}
	return cfg.Width, cfg.Height, nil
}

var errNoImage = errors.New("image illisible")

// encodePNG : utilitaire de test (images synthetiques).
func encodePNG(w, h int) ([]byte, error) {
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	for i := range img.Pix {
		img.Pix[i] = 0x40
	}
	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
