package services

// import (
// 	"fmt"

// 	"gocv.io/x/gocv"
// )

// // TODO: Work on this function
// func CaptureWebCamImage() {
// 	webcam, err := gocv.OpenVideoCapture(0)
// 	if err != nil {
// 		fmt.Printf("Erro ao abrir a webcam: %v\n", err)
// 		return
// 	}
// 	defer webcam.Close()

// 	img := gocv.NewMat()
// 	defer img.Close()

// 	if ok := webcam.Read(&img); !ok {
// 		fmt.Printf("Erro ao capturar imagem da webcam\n")
// 		return
// 	}
// 	if img.Empty() {
// 		fmt.Println("Imagem capturada está vazia")
// 		return
// 	}

// 	filename := "captured_image.png"

// 	file, err := os.Create(filename)
// 	if err != nil {
// 		fmt.Printf("Erro ao criar arquivo: %v\n", err)
// 		return
// 	}
// 	defer file.Close()

// 	imgToSave, err := img.ToImage()
// 	if err != nil {
// 		fmt.Printf("Erro ao converter imagem: %v\n", err)
// 		return
// 	}

// 	err = png.Encode(file, imgToSave)
// 	if err != nil {
// 		fmt.Printf("Erro ao salvar imagem: %v\n", err)
// 		return
// 	}

// 	fmt.Printf("Imagem capturada e salva em: %s\n", filename)

// 	return filename
// }
