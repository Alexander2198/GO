# "Rest two numbers" project in GO

This is a simple **GO** project.The project consist in Rest two numbers and calculate the result. The project is dockerized, so you can easily run it on any system that has **Docker** installed.

## Requirements
- git installed on your system. [GIT installation instructions](https://git-scm.com/downloads)
- Docker installed on your system. [Docker installation instructions](https://docs.docker.com/get-docker/)

## Setup and Execution
## 1 Clone the github repository
From git we must create a folder with any name where we are going to clone the repository.
Once we have defined the address where we are going to clone
we place the following command

- git clone https://github.com/Alexander2198/GO

## 2 Build the Docker image (DockerHub)
From the CMD of your Windows system we must go to the address where the repository was downloaded and place 
 * cd GO

then place the following command:
* docker build -t go-rest-app .

## 3 Run the container
* docker-compose up --build

and start in http://localhost:8080

## 2nd FORM to run the project (DOCKER HUB)  PENDIENTE 
we create a folder in cmd where the project will be saved
- inside the folder run the following command (creates the docker image)
* docker pull alexanderc7777/go-rest-app
- finally we put:
* docker run -p 8000:80 --name app alexanderc7777/go-rest-app
# ready our project will be running at http://localhost:8080


