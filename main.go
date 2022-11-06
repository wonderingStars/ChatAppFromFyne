package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"net/http"
	"time"
)

type message struct {
	widget.BaseWidget
	text, from string
}

const (
	myName        = "LightFox"
	messageIndent = 20
)

func newMessage(text, name string) *message {
	m := &message{text: text, from: name}
	m.ExtendBaseWidget(m)
	return m
}

type messageRender struct {
	msg *message
	bg  *canvas.Rectangle
	txt *widget.Label
}

func (r *messageRender) messageMinSize(s fyne.Size) fyne.Size {
	fitSize := s.Subtract(fyne.NewSize(messageIndent, 0))
	r.txt.Resize(fitSize.Max(r.txt.MinSize())) // have the wrap code run
	return r.txt.MinSize()
}

func (r *messageRender) MinSize() fyne.Size {
	itemSize := r.messageMinSize(r.msg.Size())
	return itemSize.Add(fyne.NewSize(messageIndent, 0))
}

func (r *messageRender) Layout(s fyne.Size) {
	itemSize := r.messageMinSize(s)
	itemSize = itemSize.Max(fyne.NewSize(
		s.Width-messageIndent, s.Height))
	bgPos := fyne.NewPos(0, 0)
	if r.msg.from == myName {
		r.txt.Alignment = fyne.TextAlignTrailing
		r.bg.FillColor = theme.PrimaryColorNamed(
			theme.ColorBlue)
		bgPos = fyne.NewPos(s.Width-itemSize.Width, 0)
	} else {
		r.txt.Alignment = fyne.TextAlignLeading
		r.bg.FillColor = theme.PrimaryColorNamed(
			theme.ColorGreen)
	}
	r.txt.Move(bgPos)
	r.bg.Resize(itemSize)
	r.bg.Move(bgPos)
}

func (r *messageRender) BackgroundColor() color.Color {
	return color.Transparent
}
func (r *messageRender) Destroy() {
}
func (r *messageRender) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.bg, r.txt}
}
func (r *messageRender) Refresh() {
	r.Refresh()
}

func (m *message) CreateRenderer() fyne.WidgetRenderer {

	text := widget.NewLabel(m.text)
	text.Wrapping = fyne.TextWrapWord

	return &messageRender{msg: m,
		bg: &canvas.Rectangle{}, txt: text}

}

func loadMessages() *fyne.Container {

	var MessagesBoxes = container.NewVBox()

	for i := 0; i < len(chatData); i++ {

		MessagesBoxes.Add(newMessage(chatData[i], userNamesMsg[i]))
	}
	return container.NewVBox(MessagesBoxes)

}

func makeUI(c *http.Client) fyne.CanvasObject {
	list := loadMessages()
	msg := widget.NewEntry()
	send := widget.NewButtonWithIcon("", theme.MailSendIcon(), func() {
		sendMsgToChat(c, msg.Text)
		list.Add(newMessage(msg.Text, myName))

	})

	input := container.NewBorder(nil, nil, nil, send, msg)
	return container.NewBorder(nil, input, nil, nil,
		container.NewVScroll(list))
}

func main() {

	c := GetHttpClient()
	a := app.New()
	w := a.NewWindow("Messages")

	go func() {
		getChat(c)
	}()

	time.Sleep(10 * time.Second)

	w.SetContent(makeUI(c))
	w.Resize(fyne.NewSize(160, 280))
	w.ShowAndRun()

}
