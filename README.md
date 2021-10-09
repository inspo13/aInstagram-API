# aInstagram-API


#In this repository, I have Designed and Developed an HTTP JSON API for the basic version of aInstagram.
#It is capable of doing the following operations-
#Create a User
Should be a POST request
Use JSON request body
URL should be ‘/users'
Get a user using id
Should be a GET request
The id should be in the URL parameter
URL should be ‘/users/<id here>’
Create a Post
Should be a POST request
Use JSON request body
URL should be ‘/posts'
Get a post using id
Should be a GET request
Id should be in the URL parameter
URL should be ‘/posts/<id here>’
List all posts of a user
Should be a GET request
URL should be ‘/posts/users/<Id here>'
 Project Folder Explanation:
The name of the main folder is aInstagram-API
@models - consists of post.go and user.go which has the basic structure of variables used like user name, id, etc.
@controllers - consists of methods. Basically HTTP methods.
@main.go - Is the main file of the project. 
About the technologies used to create the project-
Language - Golang
Database- Mongo DB
API Tester- Postman
