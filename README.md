# Webhook Printer

This service creates temporary Webhook URLs and displays the sent content. If the URL is entered in services, the transmitted data is displayed.

![GitHub](https://img.shields.io/github/license/pkuebler/webhook-printer?style=for-the-badge)
[![Travis (.org)](https://img.shields.io/travis/pkuebler/webhook-printer?style=for-the-badge)](https://travis-ci.org/github/PKuebler/webhook-printer)
[![Docker Image Version (latest semver)](https://img.shields.io/docker/v/pkuebler/webhook-printer?style=for-the-badge)](https://hub.docker.com/repository/docker/pkuebler/webhook-printer)
[![MicroBadger Layers](https://img.shields.io/microbadger/layers/pkuebler/webhook-printer?style=for-the-badge)](https://hub.docker.com/repository/docker/pkuebler/webhook-printer)
[![Docker Image Size (tag)](https://img.shields.io/docker/image-size/pkuebler/webhook-printer/latest?style=for-the-badge)](https://hub.docker.com/repository/docker/pkuebler/webhook-printer)
[![Docker Automated build](https://img.shields.io/docker/cloud/automated/pkuebler/webhook-printer?style=for-the-badge)](https://hub.docker.com/repository/docker/pkuebler/webhook-printer)
[![Docker Build Status](https://img.shields.io/docker/cloud/build/pkuebler/webhook-printer?style=for-the-badge)](https://hub.docker.com/repository/docker/pkuebler/webhook-printer)

![Usecase](https://raw.githubusercontent.com/PKuebler/webhook-printer/master/screenshot.png)

## Features

- Create dynamic Webhook Endpoints
- Show endpoints

This service does not store data persistently and deletes all data as soon as the browser connection is interrupted.

## Use

Start the docker container or binary. The service listen on `:8080` and serve the website.

## License

MIT