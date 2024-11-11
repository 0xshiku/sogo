package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"sogo/internal/store"
)

var usernames = []string{
	"alice", "bob", "charlie", "david", "emily",
	"frank", "grace", "harry", "ivy", "jack",
	"kate", "liam", "mia", "noah", "olive",
	"penelope", "quinn", "riley", "sam", "taylor",
	"uma", "vincent", "willow", "xavier", "yara",
	"zeke", "alex", "bella", "charlie", "david",
	"emily", "frank", "grace", "harry", "ivy",
	"jack", "kate", "liam", "mia", "noah",
	"olive", "penelope", "quinn", "riley", "sam",
}

var titles = []string{
	"Slicing Through Golang: A Beginner's Guide",
	"Understanding Slices in Golang: A Deep Dive",
	"Mastering Slices in Golang: A Comprehensive Tutorial",
	"Unraveling the Mystery of Golang Slices",
	"Slices in Golang: A Simple Yet Powerful Data Structure",
	"Advanced Techniques for Working with Golang Slices",
	"Optimizing Your Golang Code with Slices",
	"Leveraging Slices for Efficient Data Manipulation in Golang",
	"Common Pitfalls and Best Practices for Golang Slices",
	"Exploring the Inner Workings of Golang Slices",
	"Building Dynamic Data Structures with Golang Slices",
	"Implementing Custom Data Structures Using Golang Slices",
	"Real-world Applications of Golang Slices",
	"Practical Tips for Using Slices in Golang Projects",
	"Common Use Cases for Golang Slices",
	"The Magic of Slices: A Golang Developer's Secret Weapon",
	"Slicing Your Way to Efficient Golang Code",
	"Unleashing the Power of Slices in Golang",
	"Slices: The Swiss Army Knife of Golang Data Structures",
	"A Slice of Golang: A Tasty Treat for Developers",
}

var contents = []string{
	"Hello, world!",
	"How are you today?",
	"The quick brown fox jumps over the lazy dog.",
	"A stitch in time saves nine.",
	"Early to bed, early to rise, makes a man healthy, wealthy, and wise.",
	"Practice makes perfect.",
	"Where there's a will, there's a way.",
	"A penny saved is a penny earned.",
	"You reap what you sow.",
	"Don't count your chickens before they hatch.",
	"The early bird gets the worm.",
	"All good things must come to an end.",
	"Actions speak louder than words.",
	"Time flies when you're having fun.",
	"It's never too late to learn.",
	"The only way to do great work is to love what you do.",
	"The future belongs to those who believe in the beauty of their dreams.",
	"Believe you can and you're halfway there.",
	"The best way to predict the future is to create it.",
	"In three words I can sum up everything I've learned about life: it goes on.",
}

var tags = []string{
	"golang", "programming", "development", "tutorial", "guide",
	"tips", "tricks", "best practices", "code examples", "open source",
	"software engineering", "web development", "data science", "machine learning",
	"artificial intelligence", "cloud computing", "cybersecurity", "database",
	"networking", "devops",
}

var comments = []string{
	"This is a random comment.",
	"Hello, world!",
	"Code is poetry.",
	"Why is this here?",
	"I'm so tired.",
	"This code is a mess.",
	"I love coding!",
	"I hate debugging.",
	"This is a really long comment that doesn't really say anything.",
	"I'm hungry.",
	"I wonder what this does.",
	"This is a great example of how to do something.",
	"I'm not sure why this is necessary.",
	"I'm going to regret this later.",
	"This is a good place to start.",
	"I'm so confused.",
	"I'm going to need a coffee.",
	"This is a really cool feature.",
	"I'm not sure if this is the right way to do this.",
	"I'm going to break this down into smaller steps.",
}

func Seed(store store.Storage, db *sql.DB) {
	ctx := context.Background()

	users := generateUsers(100)
	tx, _ := db.BeginTx(ctx, nil)

	for _, user := range users {
		if err := store.Users.Create(ctx, tx, user); err != nil {
			log.Println("Error creating user:", err)
			return
		}
	}

	posts := generatePosts(200, users)
	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Println("Error creating post:", err)
			return
		}
	}

	comments := generateComments(500, users, posts)
	for _, comment := range comments {
		if err := store.Comments.Create(ctx, comment); err != nil {
			log.Println("Error creating comment:", err)
			return
		}
	}

	log.Println("Seeding Complete")
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)

	for i := 0; i < num; i++ {
		users[i] = &store.User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Email:    usernames[i%len(usernames)] + fmt.Sprintf("%d", i) + "@example.com",
			RoleID:   1,
		}
	}

	return users
}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)
	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]

		posts[i] = &store.Post{
			UserID:  user.ID,
			Title:   titles[rand.Intn(len(titles))],
			Content: contents[rand.Intn(len(contents))],
			Tags: []string{
				tags[rand.Intn(len(tags))],
				tags[rand.Intn(len(tags))],
			},
		}
	}

	return posts
}

func generateComments(num int, users []*store.User, posts []*store.Post) []*store.Comment {
	cms := make([]*store.Comment, num)

	for i := 0; i < num; i++ {
		cms[i] = &store.Comment{
			PostID:  posts[rand.Intn(len(posts))].ID,
			UserID:  users[rand.Intn(len(users))].ID,
			Content: comments[rand.Intn(len(comments))],
		}
	}

	return cms
}
