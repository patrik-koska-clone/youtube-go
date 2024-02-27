package desktop

import (
	"image/color"
	"log"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/patrik-koska-clone/youtube-go/pkg/config"
	"github.com/patrik-koska-clone/youtube-go/pkg/utils"
	"github.com/patrik-koska-clone/youtube-go/pkg/youtubeadapter"
)

const smilingEmoji = "\U0001F60A"

type IconContent struct{}

func (i IconContent) Name() string {
	return "YouTube-logo"
}

func (i IconContent) Content() []byte {
	content, err := os.ReadFile("static/content/youtube-logo-2431.png")
	if err != nil {
		log.Fatalf("could not open static contents\n%v", err)
	}

	return content
}

func OpenConsole(y youtubeadapter.YoutubeAdapter, c config.Config) {

	YTChannelApp := app.New()

	var buttonIcon IconContent

	launchWindow := YTChannelApp.NewWindow("YouTube-Go - youtube channels and stuff")

	label, input := CreateWidgets()

	obj := CreateAnimation()

	maxResultsButton, launchButton, versionButton := CreateButtons(&y, buttonIcon, &c, label, input)

	launchTab, maxResultsTab, versionTab := CreateTabs(launchButton, maxResultsButton, versionButton, input)

	tabs := container.NewAppTabs(
		container.NewTabItem("Videos", launchTab),
		container.NewTabItem("Max Results", maxResultsTab),
		container.NewTabItem("Version", versionTab),
	)

	tabs.SetTabLocation(container.TabLocationLeading)

	launchWindow.SetContent(container.NewVBox(
		label,
		tabs,
		obj,
	))

	launchWindow.Resize(fyne.NewSize(850, 400))
	launchWindow.ShowAndRun()
}

func CreateButtons(y *youtubeadapter.YoutubeAdapter,
	icon IconContent,
	c *config.Config,
	label *widget.Label,
	input *widget.Entry) (*widget.Button, *widget.Button, *widget.Button) {

	var (
		maxResults int64
		err        error
	)

	maxResultsButton := widget.NewButton("Apply", func() {

		maxResults, err = utils.ConvertStrToInt64(input.Text)
		if err != nil {
			log.Fatalf("could not parse ui input\n%v", err)
		}

		label.SetText("Maximum set to " + input.Text)
	})

	launchButton := widget.NewButtonWithIcon("Launch video", icon, func() {
		if maxResults == 0 {
			log.Println("Max results not set or set to 0")
			return
		}
		label.SetText("Video launched.")
		err := y.LoadNewVideo(*c, maxResults)
		if err != nil {
			log.Fatalf("console crashed\n%v", err)
		}
	})

	versionButton := widget.NewButton("Version", func() {

		label.SetText("version: " + c.Version)
	})

	return maxResultsButton, launchButton, versionButton
}

func CreateAnimation() *canvas.Rectangle {
	obj := canvas.NewRectangle(color.Black)
	obj.SetMinSize(fyne.NewSize(50, 160))

	red := color.NRGBA{R: 0xff, A: 0xff}
	blue := color.NRGBA{B: 0xff, A: 0xff}

	canvas.NewColorRGBAAnimation(red, blue, time.Second*2, func(color color.Color) {
		obj.FillColor = color
		canvas.Refresh(obj)
	}).Start()

	return obj
}

func CreateWidgets() (*widget.Label, *widget.Entry) {

	label := widget.NewLabelWithStyle("Recommendations? No thanks "+smilingEmoji, fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter maximum number of search results..")

	return label, input
}

func CreateTabs(launchButton *widget.Button,
	maxResultsButton *widget.Button,
	versionButton *widget.Button,
	input *widget.Entry) (*fyne.Container, *fyne.Container, *fyne.Container) {

	icon := canvas.NewImageFromFile("static/content/gopher-youtube.png")
	icon.FillMode = canvas.ImageFillContain
	icon.SetMinSize(fyne.NewSize(400, 400))

	VideoTab := container.NewVBox(
		icon,
		launchButton,
	)

	MaxResultsTab := container.NewVBox(
		input,
		maxResultsButton,
	)

	VersionTab := container.NewVBox(
		versionButton,
	)

	return VideoTab, MaxResultsTab, VersionTab
}
