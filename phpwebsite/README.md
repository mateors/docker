# Simple PHP Website

## How to pull from docker hub
> docker pull mateors/phpwebsite:latest

## How to run from image
> docker run -d -p 8082:80 --name MyPHPWebsite mateors/phpwebsite:latest

![website_homepage][./screenshots/website_homepage.png]

I put together this project while introducing a friend of mine to PHP. I decided to clean it up a bit and put it on Github so anyone new to PHP can have a taste of a very simple and minimal website built with PHP.

This project is meant for absolute beginners. I've intentionally kept it the most minimal possible while introducing some separation of concerns.

## Concepts

The project covers these concepts:

 * PHP variables
 * PHP arrays
 * PHP functions
 * Pretty links (/about) with fallback to query string (?page=about)
 * Basic example of separation of concerns (functionality, content, template)

## Resource 
* [Docker Link](https://hub.docker.com/repository/docker/mateors/phpwebsite/tags?page=1&ordering=last_updated)

## Lisence

MIT
