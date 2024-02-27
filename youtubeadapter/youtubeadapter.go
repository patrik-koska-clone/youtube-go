package youtubeadapter

import (
	"context"
	"errors"
	"fmt"

	"github.com/patrik-koska-clone/youtube-go/browser"
	"github.com/patrik-koska-clone/youtube-go/config"
	"github.com/patrik-koska-clone/youtube-go/utils"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

const videoKind = "youtube#video"
const channelKind = "youtube#channel"

var (
	contentPart  = []string{"contentDetails"}
	searchParts  = []string{"id", "snippet"}
	playlistPart = []string{"snippet"}
	randIndex    int
)

type YoutubeAdapter struct {
	Client *youtube.Service
}

func New(c config.Config) (*YoutubeAdapter, error) {

	service, err := youtube.NewService(context.Background(), option.WithAPIKey(c.ApiKey))
	if err != nil {
		return nil, fmt.Errorf("error creating new YouTube client: %v", err)
	}

	return &YoutubeAdapter{
		Client: service,
	}, nil

}

func (y YoutubeAdapter) LoadNewVideo(c config.Config, maxResults int64) error {
	var channelIDs []string

	randIndex = utils.ChooseRandomIndex(c.ChannelConfig.ChannelList)
	if randIndex == -1 {
		return errors.New("given list is empty: could not retrieve random index")
	}

	_, channels, err := y.MakeSearchQuery(&c.ChannelConfig.ChannelList[randIndex], maxResults)
	if err != nil {
		return fmt.Errorf("could not run search query\n%v", err)
	}

	for _, id := range channels {
		channelIDs = append(channelIDs, id)
	}

	randIndex = utils.ChooseRandomIndex(channelIDs)

	videos, err := y.GetVideosFromChannels(channelIDs[randIndex], maxResults)
	if err != nil {
		return fmt.Errorf("could not retrieve videos from channel\n%v", err)
	}

	randIndex = utils.ChooseRandomIndex(videos)

	err = browser.LaunchBrowser(videos[randIndex])
	if err != nil {
		return fmt.Errorf("could not launch browser\n%v", err)
	}

	return nil

}

func (y YoutubeAdapter) GetVideosFromChannels(channelID string,
	maxResults int64) ([]string, error) {

	var videoIDs []string

	channelCall := y.Client.Channels.
		List(contentPart).
		Id(channelID).
		MaxResults(1)

	channelResp, err := channelCall.Do()
	if err != nil {
		return nil, fmt.Errorf("error fetching channel details: %v", err)
	}

	if len(channelResp.Items) == 0 {
		return nil, fmt.Errorf("no channel found with ID: %s", channelID)
	}

	uploadsPlaylistID := channelResp.
		Items[0].
		ContentDetails.
		RelatedPlaylists.
		Uploads

	playlistCall := y.Client.PlaylistItems.
		List(playlistPart).
		PlaylistId(uploadsPlaylistID).
		MaxResults(maxResults)

	response, err := playlistCall.Do()
	if err != nil {
		return nil, fmt.Errorf("error fetching playlist items: %v", err)
	}

	for _, item := range response.Items {
		videoIDs = append(videoIDs, item.Snippet.ResourceId.VideoId)
	}

	return videoIDs, nil

}

func (y YoutubeAdapter) MakeSearchQuery(query *string, maxResults int64) (
	map[string]string,
	map[string]string,
	error) {

	call := y.Client.Search.List(searchParts).
		Q(*query).
		MaxResults(maxResults)

	response, err := call.Do()
	if err != nil {
		return nil, nil, fmt.Errorf("could not make search query\n%v", err)
	}

	titlesWithVideoIds := make(map[string]string)
	titlesWithchannelIds := make(map[string]string)

	for _, item := range response.Items {
		switch item.Id.Kind {
		case videoKind:
			titlesWithVideoIds[item.Snippet.Title] = item.Id.VideoId
		case channelKind:
			titlesWithchannelIds[item.Snippet.Title] = item.Id.ChannelId
		}
	}

	return titlesWithVideoIds, titlesWithchannelIds, nil
}
