# rePocketable

[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gomods/athens.svg)](https://github.com/gomods/athens)
[![Linux](https://svgshare.com/i/Zhy.svg)](https://svgshare.com/i/Zhy.svg)
[![macOS](https://svgshare.com/i/ZjP.svg)](https://svgshare.com/i/ZjP.svg)
[![Windows](https://svgshare.com/i/ZhY.svg)](https://svgshare.com/i/ZhY.svg)
[![Build](https://github.com/owulveryck/rePocketable/actions/workflows/go.yml/badge.svg)](https://github.com/owulveryck/rePocketable/actions/workflows/go.yml)

This tool and its webpage are under construction.

Best possible option if you want to see what it will eventually do is to run a cli tool such as to epub:

```shell
go run cmd/toEpub/*.go https://whateverpageyouwanttoread/
```

## Hacktoberfest

This is a toy project, but I am more and more relying on it. I think that hacktoberfest is a good opportunity to turn this project into a product.
I will write a contributing guide soon; meanwhile if you want to participate the urgent matters are:

- Writing a proper vision: discussing it into an issue and submitting a PR to mention in in the README
- Writing autonomous end-to-end tests: grabbing a sample page, running an `httptest` server, running a toEpub code and analysing the result
- Writing a proper documentation
- Adding a contribution guide
- sky is the limit, discuss in issues and submit PR once issues are discussed :D 

## Features

The internal libraries (used by the CLI) are implemeting those features:

- Webpage fetching and pre-processing
  - preprocessing and sanitization of figures to fetch the correct image from responsive and/or javascript tags (Medium and Toward datascience)
  - experimental feature to turn LaTeX figures into pictures ([github.com/go-latex/latex](github.com/go-latex/latex))
  - extraction of the content based on the ARC90 readility project ([github.com/cixtor/readability](github.com/cixtor/readability))
- Opengraph processing to extract meta informations ([github.com/dyatlov/go-opengraph](github.com/dyatlov/go-opengraph))
  - Generation of a cover picture with the front image of the website, the title and the author of the artible
  - Generation of a first chapter with meta data such as the publication date
- epub generation ([github.com/bmaupin/go-epub](github.com/bmaupin/go-epub))
- experimental getpocket integration
  - reading the article lists and generating epubs from the list
  - a daemon mode that will eventually runs on a ereader device to sync the list (heavy WIP)

## Configurations

Those configuration may influence various internal libraries.

| KEY                             | TYPE        | DEFAULT    | REQUIRED    | DESCRIPTION    |
|---------------------------------|-------------|------------|-------------|----------------|
| DOWNLOADER_LIVENESS_CHECK       | Duration    | 5m         | true        |                |
| DOWNLOADER_PROBE_TIMEOUT        | Duration    | 60m        | true        |                |
| DOWNLOADER_HTTP_TIMEOUT         | Duration    | 10s        | true        |                |
| DOWNLOADER_TRANSPORT_TIMEOUT    | Duration    | 5s         | true        |                |

Those configuration are used for cli using the pocket integration

| KEY                        | TYPE        | DEFAULT                         | REQUIRED    | DESCRIPTION                                                        |
|----------------------------|-------------|---------------------------------|-------------|--------------------------------------------------------------------|
| POCKET_UPDATE_FREQUENCY    | Duration    | 1h                              | true        | How often to query getPocket                                       |
| POCKET_HEALTH_CHECK        | Duration    | 30s                             | true        |                                                                    |
| POCKET_POCKET_URL          | String      | https://getpocket.com/v3/get    | true        |                                                                    |
| POCKET_CONSUMER_KEY        | String      |                                 | true        | See https://getpocket.com/developer/apps/ to get a consumer key    |
| POCKET_USERNAME            | String      |                                 |             | The pocket username (will try to fetch it if not found)            |
| POCKET_TOKEN               | String      |                                 |             | The access token, will try to fetch it if not found or invalid     |
