package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func check(err error) {  
    if err != nil {  
        panic(err)  
    }  
}

func adcionImgConc(i , j, aux int, img image.Image,img2 image.Image, wImg image.RGBA){
  for x := 0; x < i; x++ { 
      for y := 0; y < j; y++ { 
          pixel := img.At(x, y) 
          originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
          pixel2 := img2.At(x, y) 
          originalColor2 := color.RGBAModel.Convert(pixel2).(color.RGBA)
          r := float64(originalColor.R) + float64(originalColor2.R)
          g := float64(originalColor.G) + float64(originalColor2.G)
          b := float64(originalColor.B) + float64(originalColor2.B)
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
  imgPath := "img.jpg"
  f, err := os.Open(imgPath) 
  check(err) 
  defer f.Close()
  imgPath2 := "img2.jpg"
  f2, err1 := os.Open(imgPath2) 
  check(err1) 
  defer f2.Close()

  img, _, err := image.Decode(f)
  img2, _, err1 := image.Decode(f2)

  size := img.Bounds().Size() 
  rect := image.Rect(0, 0, size.X, size.Y)
  wImg := image.NewRGBA(rect)

  start := time.Now()
  go adcionImgConc(size.X / 2, size.Y, 0, img,img2, *wImg)
  go adcionImgConc(size.X, size.Y, size.X / 2, img,img2, *wImg)

  elapsed := time.Since(start)
  log.Printf("Adicion de Imagenes Concurrente tomo %s", elapsed)

  var wait int //para que no termine la gorutine
	fmt.Scanln(&wait)

  ext := filepath.Ext(imgPath)  
  name := strings.TrimSuffix(filepath.Base(imgPath), ext)  
  newImagePath := fmt.Sprintf("%s/%s_adicioncon%s", filepath.Dir(imgPath), name, ext)  
  fg, err := os.Create(newImagePath)  
  defer fg.Close()  
  check(err)  
  err = jpeg.Encode(fg, wImg, nil)  
  check(err)  
}