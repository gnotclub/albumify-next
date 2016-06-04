# This software is a Work In Progress

# albumify-next

## Imgur-like album creator

Albumify lets you create albums like imgur does, with a twist: you decide what
image host to use. You can also use different hosts for different images.

### Getting started

* Download and build the code

    ```bash
    go get github.com/gnotclub/albumify-next
    ```

* Copy `config.json` somewhere and edit it with the right data

* Run the program as

    ```bash
    cd $GOPATH/src/github.com/gnotclub/albumify-next
    albumify-next -config <path to config file>
    ```
