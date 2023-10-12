package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"strings"
	"time"

	"github.com/adrg/frontmatter"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/yaml.v2"

	"github.com/HidemaruOwO/VirtualSlime-API/src/lib"
	"github.com/HidemaruOwO/nuts/log"
)

var ISDEBUG bool = false

// type Frontmatter struct {
// Title    string `json:"title"`
// Excerpt  string `json:"excerpt"`
// Category string `json:"category"`
// }
type Frontmatter struct {
	Title    string `yaml:"title"`
	Excerpt  string `yaml:"excerpt"`
	Category string `yaml:"category"`
}

type Post struct {
	Slug        string      `json:"slug"`
	Frontmatter Frontmatter `json:"frontmatter"`
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	lib.InitEnv()

	// Check Envrionment
	log.Info("🔍 Checking $APP_ENV")
	if os.Getenv("APP_ENV") == "production" {
		log.Info("😀 Loading Production mode")
		// Donwload articles cache
		if err := lib.DownloadCache(); err != nil {
			log.Error(err)
		}
		nextDay := time.Now().Add(24 * time.Hour).Truncate(24 * time.Hour)
		tomorrowFromNow := nextDay.Sub(time.Now())

		go func() {
			time.Sleep(tomorrowFromNow)
			if err := lib.DownloadCache(); err != nil {
				log.Error(err)
			}
		}()
	} else {
		log.Info("🛠️ Loading Development mode")
		log.Info("🔍 Checking $VIRTUALSLIME_DIR")
		if virtualSlimeDir := os.Getenv("VIRTUALSLIME_DIR"); virtualSlimeDir == "" {
			log.Critical(fmt.Errorf("Please set $VIRTUALSLIME_DIR"))
		}
	}
	log.Info("🔍 Checking $DEBUG")
	if os.Getenv("DEBUG") == "true" {
		ISDEBUG = true
	}
	log.Info("DEBUG: %t", ISDEBUG)

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if os.Getenv("APP_ENV") == "production" {
				c.Response().Header().Set("Access-Control-Allow-Origin", lib.DOMAIN)
				c.Response().Header().Set("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE")
				c.Response().Header().Set("Access-Control-Allow-Headers", "*")
			} else {
				c.Response().Header().Set("Access-Control-Allow-Origin", "*")
				c.Response().Header().Set("Access-Control-Allow-Methods", "*")
				c.Response().Header().Set("Access-Control-Allow-Headers", "*")
			}
			return next(c)
		}
	})

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Running Virtual Slime API")
	})

	e.GET("/v1/posts", func(c echo.Context) error {
		var posts []Post

		if os.Getenv("APP_ENV") == "production" {
			cache := lib.ReadCache()
			log.Debug(ISDEBUG, string(cache))

			if err := json.Unmarshal(cache, &posts); err != nil {
				log.Critical(err)
			}
		} else {
			postDir := filepath.Join(os.Getenv("VIRTUALSLIME_DIR"), "posts")

			log.Info(postDir)

			files, err := os.ReadDir(postDir)
			if err != nil {
				log.Critical(err)
			}

			for _, file := range files {
				filename := file.Name()
				slug := filepath.Base(filename[:len(filename)-len(filepath.Ext(filename))])

				var matter Frontmatter
				markdownWithMeta, err := os.ReadFile(postDir + "/" + filename)
				if err != nil {
					log.Critical(err)
				}

				formats := []*frontmatter.Format{
					frontmatter.NewFormat("---", "---", yaml.Unmarshal),
				}

				_, err = frontmatter.Parse(strings.NewReader(string(markdownWithMeta)), &matter, formats...)

				if err != nil {
					log.Critical(err)
				}

				// log.Debug(ISDEBUG, matter)

				post := Post{
					Slug:        slug,
					Frontmatter: matter,
				}

				posts = append(posts, post)
			}
		}

		query := c.QueryParam("q")
		query = strings.ToLower(query)

		var results []Post

		for _, post := range posts {
			fm := post.Frontmatter
			if strings.Contains(strings.ToLower(fm.Title), query) ||
				strings.Contains(strings.ToLower(fm.Excerpt), query) ||
				strings.Contains(strings.ToLower(fm.Category), query) {
				results = append(results, post)
			}
		}

		return c.JSON(http.StatusOK, map[string]interface{}{"results": results})
	})

	port := lib.PORT
	if port == "" {
		port = "3000"
	}

	err := e.Start(":" + port)
	if err != nil {
		log.Critical(err)
	}
}
