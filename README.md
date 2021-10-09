Pre-Requisites:
	MongoDB
	Go

To Run:
1) Open Terminal and navigate to the folder
2) In terminal type "go run main.go"
3) The program will run and use port 9000 to communicate
4) Open Postman
5) Select Get to retrive data and Post to create new data

User:
For GET use URL 'http://localhost:9000/users/{id}'
For Post use URL 'http:localhost:9000/users'

Sample JSON:
{
'Name':'{name}',
'Email':'{email}',
'Password':'{password}'
}

Post:
For GET use URL 'http://localhost:9000/posts/{id}'
For Post use URL 'http:localhost:9000/posts'

Sample JSON:
{
'Caption':'{caption}',
'Image_URL':'{image url}'
}