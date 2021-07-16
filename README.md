# spotify-rest
A REST Wrapper for Spotify Web API

## Deployed Example on:
https://spotify-rest.up.railway.app

## Requirements:
- Go 1.15

## How to Create Your own:
- Clone this repository
- Go to [Spotify Developer's Dashboard](https://developer.spotify.com/dashboard/applications) and create a new app
- Create a new .env file using .env.example given
- Fill CLIENT_ID and CLIENT_SECRET with your app's Client ID and Client Secret
- Run ```go run ./server.go```
- Done!

## Available Endpoints

### ```/artist```
#### Data: Returns the query result's artist details and their albums
#### Parameters:
- ```query```: the artist's name

### ```/album```
#### Data: Returns the album's details and it's tracks
#### Parameters:
- ```id```: the albums' id

### ```/track```
#### Data: Returns the track's details
#### Parameters:
- ```id```: the track's id
