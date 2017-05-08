# panosplit

Simple tool to split a panorama into square parts.

## Usage

````
$ panosplit ~/Desktop/IMG_5357.jpg
19:16:29 main.go:55: writing pano_1.jpg
19:16:31 main.go:55: writing pano_2.jpg
19:16:32 main.go:55: writing pano_3.jpg
$ panosplit -h
Usage of panosplit:
  -prefix string
    	output prefix (default "pano_")
  -quality int
    	JPEG output quality 1-100 (default 75)
````
