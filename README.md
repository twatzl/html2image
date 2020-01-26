# html2image
A golang wrapper to render html to image. This is a wrapper arount the wkhtmltoimage program (https://wkhtmltopdf.org/) which needs to be installed to work.

This is a modified version to be used as import in other applications.

## Usage

### as package

The package html2image contains the necessary functions for calling wkhtmltoimage.

Additionally the package contains the functions to handle http request and convert them to images.

For usage examples you can have a look at the code.

### as program

#### directly render image

1. render png

http://127.0.0.1:8080/to/img.png?url=http://www.google.com

2. render jpg

http://127.0.0.1:8080/to/img.jpg?url=http://www.google.com

#### render image and return image url info

1. render image and return json

http://127.0.0.1:8080/api/v1/to/img.json?url=http://www.google.com&format=jpg

2. show image url from the json result

http://127.0.0.1:8080/show/img/{your image url}

#### More Params In Url
```shell
 html: the html content to render, if url has set, this param will ignore
 width: the html page width
 height: the html page height
 quality: the image quality
 
```





