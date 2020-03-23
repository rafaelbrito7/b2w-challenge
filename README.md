
<h2 align="center">
  <img alt="Fastfeet" src="https://res.cloudinary.com/rafaelbrito7/image/upload/v1584980195/b2w-logo.png" />  
  </br>
  </br>
  Junior Software Engineer Challenge
</h2>

### :closed_book: How to run the project

#### Step 1: Clone the repository

``` git clone git@github.com:rafaelbrito7/challenge-b2w.git ```

#### Step 2: Run the application

Inside the root folder of the project, run:

``` docker-compose up --build ```

This command will initiate the environment and the application is listening to the **localhost:8000**
   
### :rocket: How to use the API
  
  First of all, it's necessary to create a planet to populate the database:
  
  - Make a POST request to **/api/planets**, passing the following object structure:
    ``` 
    {
      "name": "Hoth",
      "climate": "frozen",
      "terrain": "tundra, ice caves, mountain ranges"
    } 
    ```
    
  After creating the planet, it's possible to make requests to other methods to manipulate data:
    
  - Make a GET request to **/api/totalPlanets** to get all planets registered in the database
  
  - Make a GET request to **/api/planets/{id}**, passing a query param to get an especific planet by it's ID:
    ###### For example:
    ```
    localhost:8000/api/planets/5e78dec985472b7cf50fc1b2
    ```
  - Make a GET request to **/api/planets?name=Example**, passing a query string to get an especific planet by it's name:
    ###### For example:
    ``` 
    localhost:8000/api/planets?name=Hoth 
    ```
  - Make a DELETE request to **/api/planets/{id}**, passing a query param to delete an especific planet by it's id:
    ###### For example:
    ```
    localhost:8000/api/planets/5e78dec985472b7cf50fc1b2
    ```
    
### :computer: Technologies

This project was developed in order to become a Junior Software Engineer at B2W with the following technologies:

-  [Docker](https://www.docker.com/)
-  [Golang](https://golang.org/)
-  [Goland IDE](https://www.jetbrains.com/go/)
-  [Gorilla Mux](https://github.com/gorilla/mux)
-  [Mongo GO Driver](https://github.com/mongodb/mongo-go-driver)

### :memo: License
This project is under the MIT license.
