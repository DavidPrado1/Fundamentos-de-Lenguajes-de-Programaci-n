package main

import (  
    "os"; "image"; "image/color" ;"strconv"; "time"; "log";
	_ "image/jpeg";"bufio"
)


func histograma(path string,path2 string){
  var arr []int
  for j := 0; j <= 255; j++{
    arr = append(arr, 0)
  }
	file, err := os.Open(path)

			//handle errors while opening
	if err != nil {
	  log.Fatalf("Error when opening file: %s", err)
	}

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan(){
		sv := fileScanner.Text()
		if sv, err := strconv.Atoi(sv); err == nil {
			arr[sv]++;
		}
	}
  file.Close()
  var file1, err1 = os.OpenFile(path2, os.O_RDWR, 0644)
  if err1 != nil {
        log.Fatalf("%s", err1)
  }
  defer file1.Close()
  for p := 0; p < len(arr); p++ {
    if p == len(arr)-1{
      _, err1 = file1.WriteString(strconv.Itoa(arr[p]))
      if err1 != nil {
        log.Fatalf("%s", err1)
      }
    }else{
      _, err1 = file1.WriteString(strconv.Itoa(arr[p])+"\n")
      if err1 != nil {
        log.Fatalf("%s", err1)
      }
    }
  }
  err1 = file1.Sync()
  if err1 != nil {
        log.Fatalf("%s", err1)
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
  x1 := img.Bounds().Dx()
  y1 := img.Bounds().Dy()

  start := time.Now()
  var file, err1 = os.OpenFile("r.txt", os.O_RDWR, 0644)
  if err1 != nil {
        log.Fatalf("%s", err1)
  }
  defer file.Close()
  var file2, err2 = os.OpenFile("g.txt", os.O_RDWR, 0644)
  if err2 != nil {
        log.Fatalf("%s", err2)
  }
  defer file.Close()
  var file3, err3 = os.OpenFile("b.txt", os.O_RDWR, 0644)
  if err3 != nil {
        log.Fatalf("%s", err3)
  }
  defer file.Close()
  for x := 0; x < x1; x++ { 
    for y := 0; y < y1; y++ { 
      pixel := img.At(x, y)
      if(x == x1 - 1 && y == y1 - 1 ){
        originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
        _, err1 = file.WriteString(strconv.Itoa(int(originalColor.R)))
        if err1 != nil {
          log.Fatalf("%s", err1)
        }
        _, err2 = file2.WriteString(strconv.Itoa(int(originalColor.G)))
        if err2 != nil {
          log.Fatalf("%s", err2)
        }
        _, err3 = file3.WriteString(strconv.Itoa(int(originalColor.B)))
        if err3 != nil {
          log.Fatalf("%s", err3)
        }
      }else{
        originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
        _, err1 = file.WriteString(strconv.Itoa(int(originalColor.R))+"\n")
        if err1 != nil {
          log.Fatalf("%s", err1)
        }
        _, err2 = file2.WriteString(strconv.Itoa(int(originalColor.G))+"\n")
        if err2 != nil {
          log.Fatalf("%s", err2)
        }
        _, err3 = file3.WriteString(strconv.Itoa(int(originalColor.B))+"\n")
        if err3 != nil {
          log.Fatalf("%s", err3)
        }
      }
    } 
  } 
  err1 = file.Sync()
  if err1 != nil {
        log.Fatalf("%s", err1)
  }
  err2 = file.Sync()
  if err2 != nil {
        log.Fatalf("%s", err2)
  }
  err3 = file.Sync()
  if err3 != nil {
        log.Fatalf("%s", err3)
  }

  elapsed := time.Since(start)
  histograma("r.txt","r2.txt")
  histograma("g.txt","g2.txt")
  histograma("b.txt","b2.txt")
  
  log.Printf("Recoleccion de datos tomo %s", elapsed) 
}