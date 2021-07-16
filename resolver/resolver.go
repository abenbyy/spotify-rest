package resolver

import (
	"encoding/json"
	"github.com/abenbyy/spotify-rest/auth"
	"github.com/abenbyy/spotify-rest/model"
	"io/ioutil"
	"log"
	"net/http"
	url2 "net/url"
)

func ParseString(i interface{}) (string){
	if i != nil {return i.(string)}
	return ""
}

func QueryArtist(name string) (*model.Artist, error) {
	auth.ValidateToken()
	url := "https://api.spotify.com/v1/search?q=" + url2.QueryEscape(name) + "&type=artist&limit=1&offset=0"

	bearer := "Bearer " + auth.ACCESS_TOKEN

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", bearer)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}

	var result map[string]interface{}
	json.Unmarshal(body, &result)
	artist := result["artists"].(map[string]interface{})["items"].([]interface{})[0].(map[string]interface{})
	image := artist["images"].([]interface{})[0].(map[string]interface{})["url"]

	url2 := "https://api.spotify.com/v1/artists/" + artist["id"].(string) + "/albums"
	req, err = http.NewRequest("GET", url2, nil)
	req.Header.Add("Authorization", bearer)

	resp, err = client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}

	var result2 map[string]interface{}
	json.Unmarshal(body, &result2)

	albumsJSON := result2["items"].([]interface{})
	var albums []*model.Album

	for i := 0; i < len(albumsJSON); i++ {
		album := albumsJSON[i].(map[string]interface{})
		albumIMG := album["images"].([]interface{})[0].(map[string]interface{})["url"]
		albums = append(albums, &model.Album{
			ID:    album["id"].(string),
			Name:  album["name"].(string),
			Image: albumIMG.(string),
		})
	}

	res := &model.Artist{
		ID:     artist["id"].(string),
		Name:   artist["name"].(string),
		Image:  image.(string),
		Albums: albums,
	}

	return res, nil
}

func GetAlbum(id string) (*model.Album){
	auth.ValidateToken()
	url:="https://api.spotify.com/v1/albums/"+id
	bearer := "Bearer " + auth.ACCESS_TOKEN

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", bearer)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}

	var result map[string]interface{}
	json.Unmarshal(body, &result)

	return &model.Album{
		ID:     result["id"].(string),
		Name:   result["name"].(string),
		Image:  result["images"].([]interface{})[0].(map[string]interface{})["url"].(string),
		Tracks: GetAlbumTracks(result["id"].(string)),
	}
}

func GetAlbumTracks(id string) ([]*model.Track){
	auth.ValidateToken()
	url:= "https://api.spotify.com/v1/albums/" + id
	bearer := "Bearer " + auth.ACCESS_TOKEN

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", bearer)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}

	var result map[string]interface{}
	json.Unmarshal(body, &result)

	tracksJSON := result["tracks"].(map[string]interface{})["items"].([]interface{})
	var tracks []*model.Track

	for i:=0 ; i< len(tracksJSON) ; i++{
		track := tracksJSON[i].(map[string]interface{})
		tracks = append(tracks, &model.Track{
			ID:         track["id"].(string),
			Name:       track["name"].(string),
			PreviewURL: ParseString(track["preview_url"]),
		})
	}

	return tracks
}

func GetTrack(id string) *model.Track{
	auth.ValidateToken()
	url:= "https://api.spotify.com/v1/tracks/" + id
	bearer := "Bearer " + auth.ACCESS_TOKEN

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", bearer)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}

	var result map[string]interface{}
	json.Unmarshal(body, &result)
	return &model.Track{
		ID:         result["id"].(string),
		Name:       result["name"].(string),
		PreviewURL: ParseString(result["preview_url"]),
	}
}