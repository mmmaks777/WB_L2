// Плюсы:
// 1. С помощью паттерна "Фасад" можно сделать простой интерфейс для сложной подсистемы.
// 2. У конечного пользователя нет доступа к реализации системы, он предоставляет именно те методы которые нужны клиенту.
// Минусы:
// 1. Может стать божественным объектом, если фасад начинает обрастать функциональностью, он может превратиться в монолит.
// Применимость:
// 1. Когда требуется упростить взаимодействие с сложной системой, предоставив простой интерфейс.
// 2. Когда система состоит из множества взаимосвязанных классов, и нужно изолировать клиента от этой сложности.
// Примеры:
// 1. Библиотеки для работы с базами данных, предоставляют простой интерфейс для выполнения запросов, скрывая детали подключения и взаимодействия.
// 2. GUI-фреймворки предоставляют фасады для создания окон, кнопок и других элементов, скрывая сложность низкоуровневого программирования.

package main

import "fmt"

type audioPlayer struct{}

func (a *audioPlayer) loadAudio(file string) {
	fmt.Println("Loaoding audio: ", file)
}

func (a *audioPlayer) playAudio() {
	fmt.Println("Playing audio...")
}

type videoPlayer struct{}

func (v *videoPlayer) loadVideo(file string) {
	fmt.Println("Loading video: ", file)
}

func (v *videoPlayer) playVideo() {
	fmt.Println("Playing video...")
}

type imageViewer struct{}

func (i *imageViewer) loadImage(file string) {
	fmt.Println("Loading image: ", file)
}

func (i *imageViewer) displayImage() {
	fmt.Println("Displaying image...")
}

type MediaFacade struct {
	audioPlayer *audioPlayer
	videoPlayer *videoPlayer
	imageViewer *imageViewer
}

func NewMediaFacade() *MediaFacade {
	return &MediaFacade{
		audioPlayer: &audioPlayer{},
		videoPlayer: &videoPlayer{},
		imageViewer: &imageViewer{},
	}
}

func (m *MediaFacade) PlayMedia(file string, mediaType string) {
	switch mediaType {
	case "audio":
		m.audioPlayer.loadAudio(file)
		m.audioPlayer.playAudio()
	case "video":
		m.videoPlayer.loadVideo(file)
		m.videoPlayer.playVideo()
	case "image":
		m.imageViewer.loadImage(file)
		m.imageViewer.displayImage()
	default:
		fmt.Println("Unknown madia type")
	}
}

func main() {
	mediaPlayer := NewMediaFacade()

	mediaPlayer.PlayMedia("song.mp3", "audio")
	mediaPlayer.PlayMedia("video.mp4", "video")
	mediaPlayer.PlayMedia("image.png", "image")
}
