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

3. 
Controller -> Get Routes
Service layer -> size and page, to deal with business logics
Repo -> working on resources

type App struct {
	cr CommentsRepository
	mr MoviesRepository
	cs CommentsService
	ms MoviesService
}

cs := NewCommentsService(cr.NewMongoRepository(mongoconnection), cr.NewRedisepository(redisconnection))
ms := NewMoviesService(app.mr, mr.NewMovieRepository(app.config, app.client))

type MoviesService interface {
	r MoviesRepository
}

type CommentsService struct {
	rm CommentsMongoRepository
	rs CommentsRedisRepository
}

func (cs CommentsService) NewCommentsService(r CommentsRepository) CommentsService {
	return &commentsService{r: r}
}

func (cs CommentsService) GetComments(page, size int, ctx context.Context) ([]*models.Comment, error) {
	if page < 0 {
		page = 0
	}

	page = page * size - 10
	if cs.rs.GetComments(page, size, ctx) != nil {
		return cs.rs.GetComments(page, size, ctx)
	}

	comments := cs.r.GetComments(page, size, ctx)
	cs.rs.SetComments(comments, ctx)

	return comments, nil
}