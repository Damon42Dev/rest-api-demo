package mock

import (
	"context"
	"errors"
	"example/rest-api-demo/src/models"
	"time"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Test data
var (
	TestMovieID1, _   = primitive.ObjectIDFromHex("60c72b2f9b1e8a5d6c8b4567")
	TestMovieID2, _   = primitive.ObjectIDFromHex("60c72b2f9b1e8a5d6c8b4568")
	TestCommentID1, _ = primitive.ObjectIDFromHex("60c72b2f9b1e8a5d6c8b4567")
	TestCommentID2, _ = primitive.ObjectIDFromHex("60c72b2f9b1e8a5d6c8b4568")
)

var (
	TestMovies = []*models.Movie{
		{
			ID:        TestMovieID1,
			Title:     "Test Movie 1",
			Plot:      "This is a test movie 1",
			Directors: []string{"Test Director 1"},
			Released:  time.Now(),
			Year:      2024,
			Rated:     "PG-13",
			Runtime:   120,
			Type:      "movie",
		},
		{
			ID:        TestMovieID2,
			Title:     "Test Movie 2",
			Plot:      "This is a test movie 2",
			Directors: []string{"Test Director 2"},
			Released:  time.Now(),
			Year:      2024,
			Rated:     "PG-13",
			Runtime:   120,
			Type:      "movie",
		},
	}

	TestComments = []*models.Comment{
		{
			ID:      TestCommentID1,
			Date:    primitive.NewDateTimeFromTime(time.Now()),
			Email:   "test1@example.com",
			MovieID: primitive.NewObjectID(),
			Name:    "Test User 1",
			Text:    "This is a test comment 1",
		},
		{
			ID:      TestCommentID2,
			Date:    primitive.NewDateTimeFromTime(time.Now()),
			Email:   "test2@example.com",
			MovieID: primitive.NewObjectID(),
			Name:    "Test User 2",
			Text:    "This is a test comment 2",
		},
	}
)

// MockMoviesRepository implements MoviesRepository interface
type MockMoviesRepository struct {
	mock.Mock
}

func (m *MockMoviesRepository) GetMovies(findOptions *options.FindOptions, ctx context.Context) ([]*models.Movie, error) {
	if findOptions != nil {
		// Handle pagination
		skip := findOptions.Skip
		limit := findOptions.Limit

		if limit != nil {
			// Convert to int64 for comparison
			limitVal := *limit
			if limitVal < int64(len(TestMovies)) {
				// If skip is nil, start from beginning
				start := int64(0)
				if skip != nil {
					start = *skip
				}
				end := start + limitVal
				if end > int64(len(TestMovies)) {
					end = int64(len(TestMovies))
				}
				return TestMovies[start:end], nil
			}
		}
	}
	return TestMovies, nil
}

func (m *MockMoviesRepository) GetMovieByID(id string, ctx context.Context) (*models.Movie, error) {
	for _, movie := range TestMovies {
		if movie.ID.Hex() == id {
			return movie, nil
		}
	}
	return nil, errors.New("movie not found")
}

func (m *MockMoviesRepository) CreateMovie(movie *models.Movie, ctx context.Context) (string, error) {
	TestMovies = append(TestMovies, movie)
	return movie.ID.Hex(), nil
}

func (m *MockMoviesRepository) UpdateMovie(id string, movie *models.Movie, ctx context.Context) error {
	for i, m := range TestMovies {
		if m.ID.Hex() == id {
			TestMovies[i] = movie
			return nil
		}
	}
	return errors.New("movie not found")
}

func (m *MockMoviesRepository) DeleteMovie(id string, ctx context.Context) error {
	for i, movie := range TestMovies {
		if movie.ID.Hex() == id {
			TestMovies = append(TestMovies[:i], TestMovies[i+1:]...)
			return nil
		}
	}
	return errors.New("movie not found")
}

// MockCommentsRepository implements CommentsRepository interface
type MockCommentsRepository struct {
	mock.Mock
}

func (m *MockCommentsRepository) GetComments(findOptions *options.FindOptions, ctx context.Context) ([]*models.Comment, error) {
	return TestComments, nil
}

func (m *MockCommentsRepository) GetCommentByID(id string, ctx context.Context) (*models.Comment, error) {
	for _, comment := range TestComments {
		if comment.ID.Hex() == id {
			return comment, nil
		}
	}
	return nil, errors.New("comment not found")
}

func (m *MockCommentsRepository) CreateComment(comment *models.Comment, ctx context.Context) (string, error) {
	TestComments = append(TestComments, comment)
	return comment.ID.Hex(), nil
}

func (m *MockCommentsRepository) UpdateComment(id string, comment *models.Comment, ctx context.Context) error {
	for i, c := range TestComments {
		if c.ID.Hex() == id {
			TestComments[i] = comment
			return nil
		}
	}
	return errors.New("comment not found")
}

func (m *MockCommentsRepository) DeleteComment(id string, ctx context.Context) error {
	for i, comment := range TestComments {
		if comment.ID.Hex() == id {
			TestComments = append(TestComments[:i], TestComments[i+1:]...)
			return nil
		}
	}
	return errors.New("comment not found")
}

func (m *MockCommentsRepository) DeleteCommentByID(id string, ctx context.Context) error {
	return m.DeleteComment(id, ctx)
}

func (m *MockCommentsRepository) GetCommentsForMovie(findOptions *options.FindOptions, movieID string, ctx context.Context) ([]*models.Comment, error) {
	var movieComments []*models.Comment
	for _, comment := range TestComments {
		if comment.MovieID.Hex() == movieID {
			movieComments = append(movieComments, comment)
		}
	}
	return movieComments, nil
}

func (m *MockCommentsRepository) UpdateCommentByID(id string, updateData bson.M, ctx context.Context) error {
	for i, comment := range TestComments {
		if comment.ID.Hex() == id {
			// Update fields based on updateData
			if name, ok := updateData["name"].(string); ok {
				TestComments[i].Name = name
			}
			if email, ok := updateData["email"].(string); ok {
				TestComments[i].Email = email
			}
			if text, ok := updateData["text"].(string); ok {
				TestComments[i].Text = text
			}
			return nil
		}
	}
	return errors.New("comment not found")
}

// NewMockMoviesRepository creates a new mock movies repository
func NewMockMoviesRepository() *MockMoviesRepository {
	return &MockMoviesRepository{}
}

// NewMockCommentsRepository creates a new mock comments repository
func NewMockCommentsRepository() *MockCommentsRepository {
	return &MockCommentsRepository{}
}
