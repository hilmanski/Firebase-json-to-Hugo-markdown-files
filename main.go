package main

import (
    "fmt"
    "os"
    "io"
    "io/ioutil"
    "encoding/json"
    "strings"
     "time"
)

type Blogs map[string]Blog

//feel free to edit to match you json needs
type Blog struct{
    Title string `json:"title"`
    Slug string `json:"slug"`
    Subject string `json:"desc"`
    Tag string `json:"tag"`
    Time `json:"created_at"`
}

type Time struct{
    Seconds int `json:"_seconds"`
}

func main() {
    // Open jsonFile
    jsonFile, err := os.Open("./tak-blog.json")

    if err != nil {
        fmt.Println(err)
    }
    defer jsonFile.Close()

    //convert to byte
    jsonByte, _ := ioutil.ReadAll(jsonFile)

    var blogs Blogs
    json.Unmarshal(jsonByte, &blogs)

    for _, blog := range blogs {
        //metaData add double quote
        title := "\"" + blog.Title + "\""
        description := "\"" + createMetaSubject(blog.Subject) + "\""
        tag   := "\"" + blog.Tag + "\""
        created_at := "\"" + convertTime(blog.Time.Seconds) + "\""

        slug  := blog.Slug
        subject := blog.Subject

        //markdown file structure
        text := "+++\n" +
                "title = "+ title + "\n" +
                "description = "+ description + "\n" +
                "date = "+ created_at + "\n" +
                "tag = "+ tag + "\n" +
                "+++ \n" +
                subject

        if err := WriteStringToFile(slug +".md", text); err != nil {
          panic(err)
        }
    }
}

func WriteStringToFile(filepath, s string) error {
	fo, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer fo.Close()

	_, err = io.Copy(fo, strings.NewReader(s))
	if err != nil {
		return err
	}
	return nil
}

func convertTime(seconds int) string {
    timeRaw := time.Unix(int64(seconds), 0)
    timeNew := strings.Split(timeRaw.String(), " ")
    date := timeNew[0]
    return date +"T"+ timeNew[1] + "+00:00"
}

func createMetaSubject(text string) string {
    text = strings.Replace(text, "\n", "", -1)
    text = strings.Replace(text, "\"", "'", -1) //prevent if any double quote in meta
    if len(text) < 120 {
        return text
    } else {
        rune := []rune(text)
        return string(rune[0:120])
    }
}
