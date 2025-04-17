package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/bxcodec/faker/v3"
)

// --- Main (simplified code with data generation) ---

func main() {
	// Backward compatibility with old golang versions
	rand.Seed(time.Now().UnixNano())

	// Repositories
	userRepo := NewUserRepo()
	postRepo := NewPostRepo()
	commentRepo := NewCommentRepo()

	// Services
	userService := NewUserService(userRepo)
	postService := NewPostService(postRepo)
	commentService := NewCommentService(commentRepo)

	// Facade
	blogFacade := NewBlogFacade(userService, postService, commentService)

	// Generate users
	numUsers := 10
	for i := 1; i <= numUsers; i++ {
		name := faker.Name()
		blogFacade.RegisterUser(i, name)
	}

	// Generate posts
	numPosts := 20
	for i := 1; i <= numPosts; i++ {
		authorID := rand.Intn(numUsers) + 1
		title := faker.Sentence()
		body := faker.Paragraph()
		blogFacade.CreatePost(i, authorID, title, body)
	}

	// Generate comments
	numComments := 50
	for i := 1; i <= numComments; i++ {
		authorID := rand.Intn(numUsers) + 1
		postID := rand.Intn(numPosts) + 1
		text := faker.Sentence()
		blogFacade.AddComment(i, authorID, postID, text)
	}

	// Get all posts with their authors and comments
	allPosts, err := blogFacade.GetAllPosts()
	if err != nil {
		log.Fatal(err)
	}

	// Convert to JSON and print
	jsonAllPosts, err := json.MarshalIndent(allPosts, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonAllPosts))

	// Get post with comments and author
	aggregatedPost, _ := blogFacade.GetPostWithComments(5)
	jsonPost, _ := json.MarshalIndent(aggregatedPost, "", "  ")
	fmt.Println(string(jsonPost))
}