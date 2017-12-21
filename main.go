package main

import (
    "fmt"
    "io/ioutil"
    "html/template"
    "path/filepath"
    "log"
    "sort"
    "strconv"
    "strings"
    "net/http"
    "github.com/gin-contrib/static"
    "github.com/gin-gonic/gin"
    "gopkg.in/russross/blackfriday.v2"
)

type Post struct {
    Title string
    Date string
    Image string
    Summary string
    Content template.HTML
    File string
}

func main() {
   r := gin.Default()
   r.Use(gin.Logger())
   r.Delims("{{", "}}")

   r.Use(static.Serve("/assets", static.LocalFile("./theme/assets", true)))
   r.LoadHTMLGlob("./theme/*.tmpl.html")

   r.GET("/", func(c *gin.Context) {
       start := 0
       end := 12
       previous := "/"
       next := "/"

       a := []Post{}
       files, err := filepath.Glob("posts/summary/*")
       if err != nil {
           log.Fatal(err)
       }
       if len(files) < end {
           end = len(files)
       } else {
           next = "/page/2"
       }

       files = files[start:end]
       sort.Sort(sort.Reverse(sort.StringSlice(files)))
	
       for _, f := range files {
           file := strings.Replace(f, "posts/summary/", "post/", -1)
           file = strings.Replace(file, ".md", "", -1)
           fileread, _ := ioutil.ReadFile(f)
           lines := strings.Split(string(fileread), "\n")
           title := string(lines[0])
           date := string(lines[1])
           image := string(lines[2])
           summary := string(lines[3])
           content := template.HTML(lines[3])
           a = append(a, Post{title, date, image, summary, content, file})
       }

       c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
           "posts": a,
	   "next": next,
	   "previous": previous,
       })
    })

     r.GET("/page/:page", func(c *gin.Context) {
       page, err := strconv.Atoi(c.Param("page"))
       offset := 12
       start := 0
       end := start + offset * page
       start = end - offset

       previous := "/"
       next := "/"

       if page > 1 {
	   page = page - 1
           previous = "/page/"+ string(page)
       }

       a := []Post{}
       files, err := filepath.Glob("posts/summary/*")
       if err != nil {
           log.Fatal(err)
       }

       if len(files) < end {
           end = len(files)
       } else {
           page = page + 1
           next = "/page/" + string(page)
       }

       files = files[start:end]
       sort.Sort(sort.Reverse(sort.StringSlice(files)))

       for _, f := range files {
           file := strings.Replace(f, "posts/summary/", "post/", -1)
           file = strings.Replace(file, ".md", "", -1)
           fileread, _ := ioutil.ReadFile(f)
           lines := strings.Split(string(fileread), "\n")
           title := string(lines[0])
           date := string(lines[1])
           image := string(lines[2])
           summary := string(lines[3])
           content := template.HTML(lines[3])
           a = append(a, Post{title, date, image, summary, content, file})
       }

       c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
           "posts": a,
	   "next": next,
           "previous": previous,
       })
    })


    r.GET("/post/:postName", func(c *gin.Context) {
        postName := c.Param("postName") 
	sumfile, err := ioutil.ReadFile("./posts/summary/" + postName + ".md")
        confile, err := ioutil.ReadFile("./posts/content/" + postName + ".md")
	previous := "/"
	next := "/"
	files, err := filepath.Glob("posts/summary/*")
	for key, f := range files {
	    
	    if f == "posts/summary/" + postName + ".md" {
                if key != 0 {
		    previous = files[key - 1]
		    previous = strings.Replace(previous, "posts/summary/", "", -1)
		    previous = strings.Replace(previous, ".md", "", -1)
		} else {
		    previous = "/"
		}

		if key < len(files) - 1 {
		   next = files[key + 1]
		   next = strings.Replace(next, "posts/summary/", "", -1)
                   next = strings.Replace(next, ".md", "", -1)
		} else {
		   next = "/"
		}
	    }
	}

        if err != nil {
           fmt.Println(err)
           c.HTML(http.StatusNotFound, "error.tmpl.html", nil)
           return
        }
	lines := strings.Split(string(sumfile), "\n")
	title := string(lines[0])
	date := string(lines[1])
	image := string(lines[2])
	summary := string(lines[3])
	content := template.HTML(blackfriday.Run([]byte(confile)))

        post := Post{Title:title, Date:date, Image:image, Summary:summary, Content:content, File:postName}

        c.HTML(http.StatusOK, "post.tmpl.html", gin.H{
	    "Title": post.Title,
	    "Date": post.Date,
	    "Image": post.Image,
	    "Summary": post.Summary,
	    "Content": post.Content,
	    "File": post.File,
	    "Next": next,
	    "Previous": previous,
        })
    })

    r.Run()
}
