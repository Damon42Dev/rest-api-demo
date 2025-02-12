Hi Damon :)

Your task is to create a JSON REST API in Golang to handle CRUD operations with MongoDB of your choice.

Requirement

- All the project should be containerized with Docker. -> ✔️
- MongoDB should be seeded. For seeding, you can use [this open dataset](https://github.com/neelabalan/mongodb-sample-dataset/tree/main/sample_mflix) -> ✔️
- The CRUD operations on the dataset will involve
    - Getting movies and comments
    - Creating, Updating and Deleting comments

- You can use any library for development of the API and database manipulation. ECAL use Gin for API development.
- Follow [clean architecture principles](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) for organising the code with a clear sepeartion of the layers.
- Write tests for the CRUD endpoints.

Happy coding :)

TODOs: 
1. 
comment of movies needs to be connected to Movie. Passing the movie id in the route

movies/:id/comments -> it shows all the comments of a specific move(id)

movies/:movie_id/comments/:comment_id

2. Inetrface