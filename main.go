package main

/**
Title: Page title, and title in left sidebar
Slug: slug-of-url
Parent: The name you wish the parent series to be called
Order: number in terms of parent order
Description: Small strap-line description which appears under the title
MetaPropertyTitle: Title for social sharing
MetaDescription: Description ~ 150 - 200 words of the page for SEO.
MetaPropertyDescription: SHORT description for social media sharing.
MetaOgURL: https://www.matolat.com/slug-of-url
*/

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

type BlogPost struct {
	Title                   string
	Slug                    string
	Parent                  string
	Content                 template.HTML
	Description             string
	Order                   int
	Headers                 []string // these are the in page h2 tags
	MetaDescription         string
	MetaPropertyTitle       string
	MetaPropertyDescription string
	MetaOgURL               string
}

type SidebarData struct {
	Categories []Category
}

type Category struct {
	Name  string
	Pages []BlogPost
	Order int
}

//var BaseURL = "http://localhost:8080"

var (
	BaseURL      = "http://localhost:8080"
	currentPosts []BlogPost
	sidebarData  SidebarData
)

func main() {
	file, err := os.OpenFile("gin.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer file.Close()

	// Set Gin's default logger to write to the file
	gin.DefaultWriter = file
	gin.DefaultErrorWriter = file

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	// Load initial sidebar data
	sidebarData, err = loadSidebarData("./markdown")
	if err != nil {
		log.Fatal(err)
	}

	// Load initial posts
	currentPosts, err = loadMarkdownPosts("./markdown")
	if err != nil {
		log.Fatal(err)
	}

	// Register the sidebar template as a partial
	r.SetFuncMap(template.FuncMap{
		"loadSidebar": func() SidebarData {
			return sidebarData
		},
		"dict": dict,
	})

	// Load HTML templates
	r.LoadHTMLGlob("templates/*")

	// Serve static assets
	r.Static("/static", "./static")

	// Start server in a goroutine
	go func() {
		if err := r.Run(); err != nil {
			log.Fatalf("Failed to run server: %v", err)
		}
	}()

	// Start a goroutine to listen for console input
	go listenForCommands()

	// Set up routes
	setupRoutes(r)

	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", gin.H{
			"Title": "Page Not Found",
		})
	})

	// Block the main goroutine
	select {}
}

func listenForCommands() {
	for {
		var command string
		_, err := fmt.Scanln(&command)
		if err != nil {
			log.Println("Error reading input:", err)
			continue
		}
		if command == "refreshMD" || command == "refreshmd" || command == "REFRESHMD" {
			log.Println("Refreshing markdown files...")
			refreshMarkdownData()
		}
	}
}

func refreshMarkdownData() {
	// Reload posts
	newPosts, err := loadMarkdownPosts("./markdown")
	if err != nil {
		log.Printf("Error refreshing markdown posts: %v", err)
		return
	}
	currentPosts = newPosts

	// Reload sidebar data
	newSidebarData, err := loadSidebarData("./markdown")
	if err != nil {
		log.Printf("Error refreshing sidebar data: %v", err)
		return
	}
	sidebarData = newSidebarData

	log.Println("Markdown files refreshed successfully.")
}

func setupRoutes(r *gin.Engine) {
	// single route for the home page
	r.GET("/", func(c *gin.Context) {
		indexPath := "./markdown/index.md"
		indexContent, err := os.ReadFile(indexPath)
		if err != nil {
			log.Printf("Error occurred during operation: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}

		post, err := parseMarkdownFile(indexContent)
		if err != nil {
			log.Printf("Error occurred during operation: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}

		sidebarLinks := createSidebarLinks(post.Headers)

		c.HTML(http.StatusOK, "index.html", gin.H{
			"Title":                   post.Title,
			"Content":                 post.Content,
			"SidebarData":             sidebarData,
			"Headers":                 post.Headers,
			"SidebarLinks":            sidebarLinks,
			"CurrentSlug":             post.Slug,
			"MetaDescription":         post.MetaDescription,
			"MetaPropertyTitle":       post.MetaPropertyTitle,
			"MetaPropertyDescription": post.MetaPropertyDescription,
			"MetaOgURL":               post.MetaOgURL,
		})
	})

	for _, post := range currentPosts {
		localPost := post
		if localPost.Slug != "" {
			// Define the route for each post
			r.GET("/"+localPost.Slug, func(c *gin.Context) {
				// Find the current post based on the slug
				var currentPost *BlogPost
				for _, post := range currentPosts {
					if post.Slug == localPost.Slug {
						currentPost = &post
						break
					}
				}

				if currentPost == nil {
					c.String(http.StatusNotFound, "Post not found")
					return
				}

				// Create sidebar links for the current post
				sidebarLinks := createSidebarLinks(currentPost.Headers)

				// Render the post with the most current data
				c.HTML(http.StatusOK, "layout.html", gin.H{
					"Title":                   currentPost.Title,
					"Content":                 currentPost.Content,
					"SidebarData":             sidebarData,
					"Headers":                 currentPost.Headers,
					"Description":             currentPost.Description,
					"SidebarLinks":            sidebarLinks,
					"CurrentSlug":             currentPost.Slug,
					"MetaDescription":         currentPost.MetaDescription,
					"MetaPropertyTitle":       currentPost.MetaPropertyTitle,
					"MetaPropertyDescription": currentPost.MetaPropertyDescription,
					"MetaOgURL":               currentPost.MetaOgURL,
				})
			})
		} else {
			log.Printf("Warning: Post titled '%s' has an empty slug and will not be accessible via a unique URL.\n", localPost.Title)
		}
	}
}

func loadMarkdownPosts(dir string) ([]BlogPost, error) {
	var posts []BlogPost
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".md") {
			path := dir + "/" + file.Name()
			content, err := os.ReadFile(path)
			if err != nil {
				return nil, err
			}

			post, err := parseMarkdownFile(content)
			if err != nil {
				return nil, err
			}

			posts = append(posts, post)
		}
	}

	return posts, nil
}

func parseMarkdownFile(content []byte) (BlogPost, error) {
	sections := strings.SplitN(string(content), "---", 2)
	if len(sections) < 2 {
		return BlogPost{}, errors.New("invalid markdown format")
	}

	metadata := sections[0]
	mdContent := sections[1]

	// deal with rogue \r's
	metadata = strings.ReplaceAll(metadata, "\r", "")
	mdContent = strings.ReplaceAll(mdContent, "\r", "")

	title, slug, parent, description, order, metaDescriptionStr,
		metaPropertyTitleStr, metaPropertyDescriptionStr,
		metaOgURLStr := parseMetadata(metadata)

	htmlContent := mdToHTML([]byte(mdContent))
	headers := extractHeaders([]byte(mdContent))

	return BlogPost{
		Title:                   title,
		Slug:                    slug,
		Parent:                  parent,
		Description:             description,
		Content:                 template.HTML(htmlContent),
		Headers:                 headers,
		Order:                   order,
		MetaDescription:         metaDescriptionStr,
		MetaPropertyTitle:       metaPropertyTitleStr,
		MetaPropertyDescription: metaPropertyDescriptionStr,
		MetaOgURL:               metaOgURLStr,
	}, nil
}

func extractHeaders(content []byte) []string {
	var headers []string
	//match only level 2 markdown headers
	re := regexp.MustCompile(`(?m)^##\s+(.*)`)
	matches := re.FindAllSubmatch(content, -1)

	for _, match := range matches {
		// match[1] contains header text without the '##'
		headers = append(headers, string(match[1]))
	}

	return headers
}

func mdToHTML(md []byte) []byte {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	parser := parser.NewWithExtensions(extensions)

	opts := html.RendererOptions{
		Flags: html.CommonFlags | html.HrefTargetBlank,
	}
	renderer := html.NewRenderer(opts)

	doc := parser.Parse(md)

	output := markdown.Render(doc, renderer)

	return output
}

func parseMetadata(metadata string) (
	title string,
	slug string,
	parent string,
	description string,
	order int,
	metaDescription string,
	metaPropertyTitle string,
	metaPropertyDescription string,
	metaOgURL string,
) {
	re := regexp.MustCompile(`(?m)^(\w+):\s*(.+)`)
	matches := re.FindAllStringSubmatch(metadata, -1)

	metaDataMap := make(map[string]string)
	for _, match := range matches {
		if len(match) == 3 {
			metaDataMap[match[1]] = match[2]
		}
	}

	title = metaDataMap["Title"]
	slug = metaDataMap["Slug"]
	parent = metaDataMap["Parent"]
	description = metaDataMap["Description"]
	orderStr := metaDataMap["Order"]
	metaDescriptionStr := metaDataMap["MetaDescription"]
	metaPropertyTitleStr := metaDataMap["MetaPropertyTitle"]
	metaPropertyDescriptionStr := metaDataMap["MetaPropertyDescription"]
	metaOgURLStr := metaDataMap["MetaOgURL"]

	orderStr = strings.TrimSpace(orderStr)
	order, err := strconv.Atoi(orderStr)
	if err != nil {
		log.Printf("Error converting order from string: %v", err)
		order = 9999
	}

	return title, slug, parent, description, order, metaDescriptionStr,
		metaPropertyTitleStr, metaPropertyDescriptionStr, metaOgURLStr
}

func loadSidebarData(dir string) (SidebarData, error) {
	var sidebar SidebarData
	categoriesMap := make(map[string]*Category)

	posts, err := loadMarkdownPosts(dir)
	if err != nil {
		return sidebar, err
	}

	for _, post := range posts {
		if post.Parent != "" {
			if _, exists := categoriesMap[post.Parent]; !exists {
				categoriesMap[post.Parent] = &Category{
					Name:  post.Parent,
					Pages: []BlogPost{post},
					Order: post.Order,
				}
			} else {
				categoriesMap[post.Parent].Pages = append(categoriesMap[post.Parent].Pages, post)
			}
		}
	}

	// convert map to slice
	for _, cat := range categoriesMap {
		sidebar.Categories = append(sidebar.Categories, *cat)
	}

	// sort categories by order
	sort.Slice(sidebar.Categories, func(i, j int) bool {
		return sidebar.Categories[i].Order < sidebar.Categories[j].Order
	})

	return sidebar, nil
}

func createSidebarLinks(headers []string) template.HTML {
	var linksHTML string
	for _, header := range headers {
		sanitizedHeader := sanitizeHeaderForID(header)
		link := fmt.Sprintf(`<li><a href="#%s">%s</a></li>`, sanitizedHeader, header)
		linksHTML += link
	}
	return template.HTML(linksHTML)
}

func sanitizeHeaderForID(header string) string {
	// lowercase
	header = strings.ToLower(header)

	// replace spaces with hyphens
	header = strings.ReplaceAll(header, " ", "-")

	// remove any characters that are not alphanumeric or hyphens
	header = regexp.MustCompile(`[^a-z0-9\-]`).ReplaceAllString(header, "")

	return header
}

func dict(values ...interface{}) (map[string]interface{}, error) {
	if len(values)%2 != 0 {
		return nil, errors.New("invalid dict call")
	}
	dict := make(map[string]interface{}, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, errors.New("dict keys must be strings")
		}
		dict[key] = values[i+1]
	}
	return dict, nil
}
