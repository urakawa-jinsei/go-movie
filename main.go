package main

import (
	"gocv.io/x/gocv"
)

func main() {
	// 入力動画のパス
	inputFile := "input.mp4"

	// 出力動画のパス
	outputFile := "output.mp4"

	// 入力動画のオープン
	input, _ := gocv.VideoCaptureFile(inputFile)
	defer input.Close()

	// 出力動画の設定
	width := int(input.Get(gocv.VideoCaptureFrameWidth))
	height := int(input.Get(gocv.VideoCaptureFrameHeight))
	fps := int(input.Get(gocv.VideoCaptureFPS))
	output, _ := gocv.VideoWriterFile(outputFile, "avc1", float64(fps), width, height, true)
	defer output.Close()

	// フレームごとの処理
	frame := gocv.NewMat()
	defer frame.Close()

	for {
		if ok := input.Read(&frame); !ok {
			break
		}

		// フレームに対する処理
		// ここに任意の処理を追加してください

		// 処理されたフレームの書き込み
		output.Write(frame)
	}
}
