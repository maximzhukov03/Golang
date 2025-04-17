package main

import (
	"errors"
	"sync"
)

// UserRepository интерфейс для операций с пользователями
type UserRepository interface {
	Save(id int, name string)
	FindByID(id int) (string, error)
}

// PostRepository интерфейс для операций с постами
type PostRepository interface {
	Save(post *Post)
	FindByID(id int) (*Post, error)
	FindAll() []*Post
}

// CommentRepository интерфейс для операций с комментариями
type CommentRepository interface {
	Save(comment *Comment)
	FindByPostID(postID int) []*Comment
}

// UserServicer интерфейс для операций с пользователями
type UserServicer interface {
	RegisterUser(id int, name string)
	FindUserByID(id int) (*User, error)
}

// PostServicer интерфейс для операций с постами
type PostServicer interface {
	CreatePost(id int, authorID int, title, body string) *Post
	FindPostByID(id int) (*Post, error)
	FindAllPosts() []*Post
}

// CommentServicer интерфейс для операций с комментариями
type CommentServicer interface {
	AddComment(id, authorID, postID int, text string) *Comment
	FindCommentsByPostID(postID int) []*Comment
}

// BlogService интерфейс для агрегирования бизнес-логики блога
type BlogService interface {
	RegisterUser(id int, name string)
	CreatePost(id, authorID int, title, body string) (*AggregatedPost, error)
	AddComment(id, authorID, postID int, text string) (*AggregatedComment, error)
	GetPostWithComments(postID int) (*AggregatedPost, error)
	GetAllPosts() ([]*ListingPost, error)
}

// --- Основные сущности ---

// User struct
type User struct {
	ID   int
	Name string
}

// Post struct
type Post struct { //Пост включает идент. поста, идент. автора, заголовок и текст поста.
	ID       int
	AuthorID int
	Title    string
	Body     string
}

// Comment struct
type Comment struct { //состоит из идент, идент автора, идент поста и текста комментария.
	ID       int
	AuthorID int
	PostID   int
	Text     string
}

// --- Repositories ---

// UserRepo - память для хранения данных пользователей 
// (хранит только идентификаторы и имена)
type UserRepo struct {
	mu    sync.Mutex
	users map[int]string
}

func NewUserRepo() *UserRepo { // Создание Памяти для хранения данных пользователя
	return &UserRepo{
		users: make(map[int]string),
	}
}

func (r *UserRepo) Save(id int, name string) { // Сохранение в памяти USERA
	r.mu.Lock()
	defer r.mu.Unlock()
	r.users[id] = name
}

func (r *UserRepo) FindByID(id int) (string, error) { // Найти USERA в памяти по ID 
	r.mu.Lock()
	defer r.mu.Unlock()
	if name, exists := r.users[id]; exists {
		return name, nil
	}
	return "", errors.New("user not found")
}

// PostRepo - память для постов
type PostRepo struct {
	mu    sync.Mutex
	posts map[int]*Post
}

func NewPostRepo() *PostRepo { // Создание Памяти для хранения данных постов
	return &PostRepo{
		posts: make(map[int]*Post),
	}
}

func (r *PostRepo) Save(post *Post) { // Сохранение в памяти постов
	r.mu.Lock()
	defer r.mu.Unlock()
	r.posts[post.ID] = post
}

func (r *PostRepo) FindByID(id int) (*Post, error) { // Найти Пост в памяти по ID 
	r.mu.Lock()
	defer r.mu.Unlock()
	if post, exists := r.posts[id]; exists {
		return post, nil
	}
	return nil, errors.New("post not found")
}

func (r *PostRepo) FindAll() []*Post { // Найти все посты
	r.mu.Lock()
	defer r.mu.Unlock()
	posts := make([]*Post, 0, len(r.posts))
	for _, post := range r.posts {
		posts = append(posts, post)
	}
	return posts
}

// CommentRepo - память для комментариев
type CommentRepo struct {
	mu       sync.Mutex
	comments map[int]*Comment
}

func NewCommentRepo() *CommentRepo { // Создать память для комментариев
	return &CommentRepo{
		comments: make(map[int]*Comment),
	}
}

func (r *CommentRepo) Save(comment *Comment) { // Сохранить в память для комментариев
	r.mu.Lock()
	defer r.mu.Unlock()
	r.comments[comment.ID] = comment
}

func (r *CommentRepo) FindByPostID(postID int) []*Comment { // Найти комментарий в посте по ID 
	r.mu.Lock()
	defer r.mu.Unlock()
	comments := make([]*Comment, 0)
	for _, comment := range r.comments {
		if comment.PostID == postID {
			comments = append(comments, comment)
		}
	}
	return comments
}

// --- Интерфейсы сервисов ---

// UserService обрабатывает операции, связанные с пользователем
type UserService struct {
	repo UserRepository
}

func NewUserService(repo *UserRepo) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) RegisterUser(id int, name string) {
	s.repo.Save(id, name)
}

func (s *UserService) FindUserByID(id int) (*User, error) {
	name, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return &User{ID: id, Name: name}, nil
}

// PostService handles post-related operations
type PostService struct {
	repo PostRepository
}

func NewPostService(repo *PostRepo) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) CreatePost(id int, authorID int, title, body string) *Post {
	post := &Post{ID: id, AuthorID: authorID, Title: title, Body: body}
	s.repo.Save(post)
	return post
}

func (s *PostService) FindPostByID(id int) (*Post, error) {
	return s.repo.FindByID(id)
}

func (s *PostService) FindAllPosts() []*Post {
	return s.repo.FindAll()
}

// CommentService handles comment-related operations
type CommentService struct {
	repo CommentRepository
}

func NewCommentService(repo *CommentRepo) *CommentService {
	return &CommentService{repo: repo}
}

func (s *CommentService) AddComment(id, authorID, postID int, text string) *Comment {
	comment := &Comment{ID: id, AuthorID: authorID, PostID: postID, Text: text}
	s.repo.Save(comment)
	return comment
}

func (s *CommentService) FindCommentsByPostID(postID int) []*Comment {
	return s.repo.FindByPostID(postID)
}

// --- Агрегированные сущности для блога ---

// AggregatedPost представляет собой пост с автором и комментариями.
type AggregatedPost struct {
	ID       int                  `json:"id"`
	Title    string               `json:"title"`
	Body     string               `json:"body"`
	Author   *User                `json:"author"`
	Comments []*AggregatedComment `json:"comments"`
}

type ListingPost struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	Body          string `json:"body"`
	Author        *User  `json:"author"`
	CommentsCount int    `json:"comments_count"`
}

// AggregatedComment представляет собой комментарий с информацией об авторе
type AggregatedComment struct {
	ID     int    `json:"id"`
	Text   string `json:"text"`
	Author *User  `json:"author"`
}

// AggregatedUserPosts представляет пользователя и его сообщения
type AggregatedUserPosts struct {
	User  *User   `json:"user"`
	Posts []*Post `json:"posts"`
}

// --- Blog ---

// Блог агрегирует бизнес-логику
type Blog struct {
	userService    UserServicer
	postService    PostServicer
	commentService CommentServicer
}

func (blog *Blog) RegisterUser(id int, name string){
	blog.userService.RegisterUser(id, name)
}

func (blog *Blog) CreatePost(id, authorID int, title, body string) (*AggregatedPost, error){
	post := blog.postService.CreatePost(id, authorID, title, body)
	user, err := blog.userService.FindUserByID(authorID)
	
	if err != nil{
		return nil, err
	}

	agregatedPost := &AggregatedPost{
		ID: post.ID,
		Title: post.Title,
		Body: post.Body,
		Author: user,
		Comments: []*AggregatedComment{},
	}
	return agregatedPost, err
}

func (blog *Blog) AddComment(id, authorID, postID int, text string) (*AggregatedComment, error){
	comment := blog.commentService.AddComment(id, authorID, postID, text)
	author, err := blog.userService.FindUserByID(authorID)
	if err != nil {
		return nil, err
	}
	return &AggregatedComment{ID: comment.ID, Text: comment.Text, Author: author}, nil
}

func (blog *Blog)GetPostWithComments(postID int) (*AggregatedPost, error){
	post, err := blog.postService.FindPostByID(postID)
	if err != nil{
		return nil, err
	}
	authorID := post.AuthorID
	author, err := blog.userService.FindUserByID(authorID)
	if err != nil{
		return nil, err
	}
	comments := blog.commentService.FindCommentsByPostID(postID)

	aggregatedComments := make([]*AggregatedComment, 0, len(comments))

	for _, comment := range comments {
		author, _ := blog.userService.FindUserByID(comment.AuthorID)
		aggregatedComments = append(aggregatedComments, &AggregatedComment{
			ID:     comment.ID,
			Text:   comment.Text,
			Author: author,
		})
	}

	return &AggregatedPost{
		ID:       post.ID,
		Title:    post.Title,
		Body:     post.Body,
		Author:   author,
		Comments: aggregatedComments,
	}, nil
}

func (blog *Blog) GetAllPosts() ([]*ListingPost, error){
	posts := blog.postService.FindAllPosts()
	var listPosts []*ListingPost 
	for _, post := range posts{
		author, _ := blog.userService.FindUserByID(post.AuthorID)
		listPost := &ListingPost{
			ID: post.ID,
			Title: post.Title,    
			Body: post.Body,
			Author: author,
			CommentsCount: len(blog.commentService.FindCommentsByPostID(post.ID)),
		}
		listPosts = append(listPosts, listPost)
	}
	
	return listPosts, nil
}



func NewBlogFacade(userService *UserService, postService *PostService, commentService *CommentService) BlogService {
	return &Blog{
		userService:    userService,
		postService:    postService,
		commentService: commentService,
	}
}
