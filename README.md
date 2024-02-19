# GroupC_Assignment3 README

This is our group project. We built a simple web server in Go that can manage to-do items. You can add, get, update, and delete items with our program.

## Files in Our Project
- **main.go:** This file has all the code for our server and functions.
- **main_test.go:** This file has tests for our server's functions.

### How to Use
To use our project, you need to have Go installed on your computer. You can download Go from here.

1. Clone the Repository

First, clone this project to your local machine using Git:

```git clone https://github.com/ashbirdhiman/groupC_Assignment3.git
cd groupC_Assignment3```

**Running the Server**
- Open a terminal or command prompt.
- Change the directory to where you have our project files.
- Type **'go run main.go'** to start the server.
```go run main.go```

- Now, the server should be running on port 8012. You can use a tool like Postman to interact with it.

### Available Actions
- **Add an item:** Send a POST request to /AddItem with item details.
- **Get all items:** Send a GET request to /GetAllItems.
- **Get one item:** Send a GET request to /GetOneItem/{itemID}.
- **Update an item:** Send a PUT request to /UpdateItem/{itemID} with new details.
- 
- **Delete an item:** Send a DELETE request to /DeleteOneItem/{itemID}.


### Running Tests
To run the tests for our server:

1. Open a terminal or command prompt.
2. Change the directory to our project folder.
3. Type **'go test'** to run the tests.
```go test```
### Team Members
Our team worked together on this project. Here are our members:

Ashbir - 500228410
Jevica - 500218849
Akash - 500218794
Abhinav Mahajan - 500230044
Vinay Chhabra - 500228151
Bilal Nawaz -  500228652
Mohamed Ayan Khatri - 500226334
Simrandeep Singh - 500229180
Nikhil Kaushik - 500223528
Rajkarn kaur - 500226333


Everyone contributed by writing parts of the code and tests.
