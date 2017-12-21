# Helix - Markdown Blogging Platform

Helix is a Markdown Blogging platform written in Golang and the Gin Framework. Helix was created to be a simple solution to getting a blog up and running with as minimal overhead as possible. Your Helix website is completely themeable and customizable.  

## Demo
[Technopathic](https://blog.technopathic.me)

## Screenshots
![Screenshot](https://technopathic.me/storage/helixScreen.png)
![Screenshot](https://technopathic.me/storage/helixScreen2.png)

## Requirements
* Go 1.9.x

## Getting Started
Be sure to set your GOPATH to this repo before you begin. If you are unfamiliar with GoLang, I highly recommend checking out this [repo](https://github.com/alco/gostart).
Once you've cloned / downloaded Helix, you will navigate into the folder and install all of hte dependencies.
```
git clone https://github.com/Technopathic/Helix
cd Helix
go get -d ./...
```

Helix's ease-of-use allows you to easily customize the theme located in the theme folder. Simply update the theme using HTML5/CSS3/JavaScript to your own liking.

In order to create a new blog post, simply run the command ``` go run new.go ``` which will create two new markdown files located in ```/posts/summary``` and ```/posts/content```. Your post summary file contains all of the metadata for your blog post while the content file should contain the body of your blog post.
Please format the summary markdown file as below:
```
Post Title
Post Date
Post Image
Post Summary
```

To run your server:
```
go run main.go
```

## License
MIT
