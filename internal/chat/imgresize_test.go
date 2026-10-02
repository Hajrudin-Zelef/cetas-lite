package chat

import (
	"bytes"
	"image/jpeg"
	"testing"
)

func TestResizeJPEGBoundsLongSide(t *testing.T) {
	src, err := encodePNG(2048, 1024)
	if err != nil {
		t.Fatal(err)
	}
	out, err := resizeJPEG(src, 1024, 80)
	if err != nil {
		t.Fatal(err)
	}
	if len(out) >= len(src) {
		t.Fatalf("l'image redimensionnee doit etre plus petite (%d vs %d)", len(out), len(src))
	}
	img, err := jpeg.Decode(bytes.NewReader(out))
	if err != nil {
		t.Fatalf("sortie JPEG invalide: %v", err)
	}
	b := img.Bounds()
	if b.Dx() != 1024 || b.Dy() != 512 {
		t.Fatalf("dimensions = %dx%d, attendu 1024x512", b.Dx(), b.Dy())
	}
}

func TestResizeJPEGNoUpscale(t *testing.T) {
	src, err := encodePNG(100, 50)
	if err != nil {
		t.Fatal(err)
	}
	out, err := resizeJPEG(src, 1024, 80)
	if err != nil {
		t.Fatal(err)
	}
	img, _ := jpeg.Decode(bytes.NewReader(out))
	if b := img.Bounds(); b.Dx() != 100 || b.Dy() != 50 {
		t.Fatalf("pas d'agrandissement attendu, obtenu %dx%d", b.Dx(), b.Dy())
	}
}

func TestResizeJPEGInvalid(t *testing.T) {
	if _, err := resizeJPEG([]byte("not an image"), 1024, 80); err == nil {
		t.Fatal("une donnee non-image doit echouer")
	}
}
