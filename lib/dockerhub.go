package lib

import (
	"github.com/mohemohe/fetch"
	"time"
)

type (
	Tags struct {
		Count    int         `json:"count"`
		Next     *string      `json:"next"`
		Previous string `json:"previous"`
		Results  []Result   `json:"results"`
	}
	Result struct {
		Creator             int         `json:"creator"`
		ID                  int         `json:"id"`
		ImageID             interface{} `json:"image_id"`
		Images              []Image    `json:"images"`
		LastUpdated         time.Time   `json:"last_updated"`
		LastUpdater         int         `json:"last_updater"`
		LastUpdaterUsername string      `json:"last_updater_username"`
		Name                string      `json:"name"`
		Repository          int         `json:"repository"`
		FullSize            int         `json:"full_size"`
		V2                  bool        `json:"v2"`
		TagStatus           string      `json:"tag_status"`
		TagLastPulled       time.Time   `json:"tag_last_pulled"`
		TagLastPushed       time.Time   `json:"tag_last_pushed"`
	}
	Image struct {
		Architecture string      `json:"architecture"`
		Features     string      `json:"features"`
		Variant      interface{} `json:"variant"`
		Digest       string      `json:"digest"`
		Os           string      `json:"os"`
		OsFeatures   string      `json:"os_features"`
		OsVersion    interface{} `json:"os_version"`
		Size         int         `json:"size"`
		Status       string      `json:"status"`
		LastPulled   time.Time   `json:"last_pulled"`
		LastPushed   time.Time   `json:"last_pushed"`
	}
)

func ListImageTags(imageName string) []string {
	result := make([]string, 0)
	now := time.Now()

	url := "https://hub.docker.com/v2/repositories/"+ imageName +"/tags"
	for {
		res, err := fetch.Fetch(url)
		if err != nil {
			panic(err)
		}
		if !res.OK {
			panic("invalid response")
		}

		tags := new(Tags)
		if err := res.JSON(tags); err != nil {
			panic(err)
		}
		for _, r := range tags.Results {
			if now.Sub(r.TagLastPulled) > 720 * time.Hour {
				result = append(result, r.Name)
			}
		}

		if tags.Next != nil {
			url = *tags.Next
		} else {
			break
		}
	}

	return result
}
