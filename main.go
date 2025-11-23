package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	_ "image/png" // registramos el decodificador PNG
	"os"
	"path/filepath"
	"strings"
)

func main() {
	inputDir := flag.String("in", "./inputs", "Carpeta de entrada con PNG")
	outputDir := flag.String("out", "./outputs", "Carpeta de salida para JPG")
	quality := flag.Int("q", 80, "Calidad JPG (1-100, recomendado 70-85)")
	flag.Parse()

	if *quality < 1 {
		*quality = 1
	}
	if *quality > 100 {
		*quality = 100
	}

	if err := os.MkdirAll(*outputDir, 0o755); err != nil {
		fmt.Println("Error creando carpeta de salida:", err)
		return
	}

	err := filepath.Walk(*inputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if !strings.EqualFold(filepath.Ext(info.Name()), ".png") {
			return nil
		}

		relPath, err := filepath.Rel(*inputDir, path)
		if err != nil {
			return err
		}

		// Cambiamos extensión a .jpg
		outRel := strings.TrimSuffix(relPath, filepath.Ext(relPath)) + ".jpg"
		outPath := filepath.Join(*outputDir, outRel)

		if err := os.MkdirAll(filepath.Dir(outPath), 0o755); err != nil {
			return err
		}

		fmt.Printf("Convirtiendo: %s -> %s\n", path, outPath)
		if err := convertPNGToJPG(path, outPath, *quality); err != nil {
			fmt.Println("  Error:", err)
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error general:", err)
	}
}

func convertPNGToJPG(inputPath, outputPath string, quality int) error {
	f, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return err
	}

	bounds := img.Bounds()
	// Creamos un lienzo RGBA con fondo blanco
	dst := image.NewRGBA(bounds)

	// Pintar fondo blanco (para manejar transparencias)
	draw.Draw(dst, bounds, &image.Uniform{C: color.White}, image.Point{}, draw.Src)
	// Dibujar la imagen original encima (respeta alfa si lo hay)
	draw.Draw(dst, bounds, img, bounds.Min, draw.Over)

	// Crear archivo de salida
	outFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	opts := &jpeg.Options{Quality: quality}
	if err := jpeg.Encode(outFile, dst, opts); err != nil {
		return err
	}

	return nil
}
