package main

import (  
  "os"
  "image"
  "image/color"
  "time"
  "log";
	"image/jpeg"
  _ "image/png"
  "path/filepath"
  "fmt"
  "strings"
)

func check(err error) {  
    if err != nil {  
        panic(err)  
    }  
}

func invertir(i , j, aux int, img image.Image, wImg image.RGBA){
  for x := 0; x < i; x++ { 
    for y := 0; y < j; y++ { 
      pixel := img.At(x, y)
      originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
      r := 255-uint8(originalColor.R)
      g := 255-uint8(originalColor.G)
      b := 255-uint8(originalColor.B)
      if r>255{
        r = 255
      } else if(r<0){
        r = 0
      } 
      if g>255{
        g = 255
      }else if(g<0){
        g = 0
      }
      if b>255{
        b = 255
      }else if(b<0){
        b = 0
      }
      c := color.RGBA{ 
        R: uint8(r), G: uint8(g), B: uint8(b), A: originalColor.A, 
      } 
      wImg.Set(x, y, c) 
    } 
  } 
}

func main() {  
  imgPath := "lena.jpg"
  f, err := os.Open(imgPath) 
  if err != nil {
		panic(err)
	}
	defer f.Close() 

  img, _, err := image.Decode(f)
  size := img.Bounds().Size()  
  rect := image.Rect(0, 0, size.X, size.Y)  
  wImg := image.NewRGBA(rect)

  start := time.Now()
  go invertir(size.X / 2, size.Y, 0, img, *wImg)
  go invertir(size.X, size.Y, size.X / 2, img, *wImg)

  elapsed := time.Since(start)
  
  log.Printf("Invertir los colores de la imagen tomo %s", elapsed) 

  ext := filepath.Ext(imgPath)  
  name := strings.TrimSuffix(filepath.Base(imgPath), ext)  
  newImagePath := fmt.Sprintf("%s/%s_invertido%s", filepath.Dir(imgPath), name, ext)  
  fg, err := os.Create(newImagePath)  
  defer fg.Close()  
  check(err)  
  err = jpeg.Encode(fg, wImg, nil)  
  check(err)
  var wait int //para que no termine la gorutine
	fmt.Scanln(&wait)
}