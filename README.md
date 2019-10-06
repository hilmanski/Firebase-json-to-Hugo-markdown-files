# Firebase-json-to-Hugo-markdown-files
convert from firebase data or json to hugo markdown files

## What is this
A Go script to generate markdown files from json file. In my case i exported my firebase data in json files to migrate it to hugo site

## How to export firebase database?
Dalenguyen's [repo export firebase](https://github.com/dalenguyen/firestore-import-export) data

## Watch out
- Make sure you change the type struct with your need
- It will generate markdown files, you still have to move all the markdown files to your hugo content folder

## How the json structure looks like
Sample data below generated from firebase (random id/key)
```
 {    
    "0BVxYRh8MixS0mYEtZ8v":{
            "title":"the titls",
            "subject":"the subject"
    },
    "0BVxYRh8MixasdS0mYEtZ8v":{
            "title":"the titls2",
            "subject":"the subject2"
    }
}
```
