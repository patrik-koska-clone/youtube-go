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
	"github.com/patrik-koska-clone/youtube-go/config"
	"github.com/patrik-koska-clone/youtube-go/utils"
	"github.com/patrik-koska-clone/youtube-go/youtubeadapter"
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
	var (
		icon            IconContent
		maxResultsInput *string
		maxResults      *int64
		err             error
	)

	YTChannelApp := app.New()

	launchWindow := YTChannelApp.NewWindow("YouTube Channels - Only which matter")

	obj := canvas.NewRectangle(color.Black)
	obj.SetMinSize(fyne.NewSize(500, 500))

	red := color.NRGBA{R: 0xff, A: 0xff}
	blue := color.NRGBA{B: 0xff, A: 0xff}

	canvas.NewColorRGBAAnimation(red, blue, time.Second*2, func(color color.Color) {
		obj.FillColor = color
		canvas.Refresh(obj)
	}).Start()

	label := widget.NewLabel("Recommendations? No thanks " + smilingEmoji)

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter maximum number of search results..")

	maxResultsButton := widget.NewButton("Save maximum", func() {
		maxResultsInput = &input.Text
		maxResults, err = utils.ConvertStrToInt64(maxResultsInput)
		if err != nil {
			log.Fatalf("could not parse ui input\n%v", err)
		}
	})

	launchButton := widget.NewButtonWithIcon("Launch video", icon, func() {
		label.SetText("opening...")
		err := y.LoadNewVideo(c, *maxResults)
		if err != nil {
			log.Fatalf("console crashed\n%v", err)
		}
	})

	launchWindow.SetContent(container.NewVBox(
		label,
		launchButton,
		input,
		maxResultsButton,
		obj,
	))

	launchWindow.Resize(fyne.NewSize(500, 500))
	launchWindow.ShowAndRun()
}
