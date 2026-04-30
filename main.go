package main

import (
	"embed"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/charmbracelet/huh"
)

//go:embed templates
var templatesFS embed.FS

var version = "dev"

var tplVarRe = regexp.MustCompile(`\{\{(\w+)\}\}`)

func main() {
	if len(os.Args) > 1 && (os.Args[1] == "--version" || os.Args[1] == "-v") {
		fmt.Println(version)
		os.Exit(0)
	}

	var site, locale, scraperType string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Site name (e.g. amazon, kaufland)").
				Value(&site).
				Validate(func(v string) error {
					if strings.TrimSpace(v) == "" {
						return errors.New("site name is required")
					}
					return nil
				}),
			huh.NewInput().
				Title("Locale (e.g. fr, co.uk, com.br)").
				Value(&locale).
				Validate(func(v string) error {
					if strings.TrimSpace(v) == "" {
						return errors.New("locale is required")
					}
					return nil
				}),
			huh.NewSelect[string]().
				Title("Scraper type").
				Options(
					huh.NewOption("products/delivery", "products/delivery"),
					huh.NewOption("products/drive", "products/drive"),
					huh.NewOption("retailOutlets", "retailOutlets"),
					huh.NewOption("productDetails", "productDetails"),
				).
				Value(&scraperType),
		),
	)

	if err := form.Run(); err != nil {
		fmt.Println("\nAborted.")
		os.Exit(0)
	}

	if err := generate(strings.TrimSpace(site), strings.TrimSpace(locale), scraperType); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func generate(site, locale, scraperType string) error {
	targetPath := filepath.Join("src", "sites", site, locale, scraperType)

	if _, err := os.Stat(targetPath); err == nil {
		return fmt.Errorf("\nDirectory already exists: %s\nAborting to avoid overwriting existing scraper.", targetPath)
	}

	if err := os.MkdirAll(targetPath, 0755); err != nil {
		return err
	}

	templateKey := strings.ReplaceAll(scraperType, "/", "-")
	vars := map[string]string{
		"SITE":         site,
		"LOCALE":       locale,
		"SCRAPER_PATH": "sites/" + site + "/" + locale + "/" + scraperType,
	}

	for _, filename := range []string{"config.js", "index.js", "transformers.js"} {
		tplPath := "templates/" + templateKey + "/" + filename + ".tpl"
		tplBytes, err := templatesFS.ReadFile(tplPath)
		if err != nil {
			return err
		}
		rendered := render(string(tplBytes), vars)
		if err := os.WriteFile(filepath.Join(targetPath, filename), []byte(rendered), 0644); err != nil {
			return err
		}
	}

	fmt.Printf("\nCreated scraper at: %s\n", targetPath)
	fmt.Println("  config.js")
	fmt.Println("  index.js")
	fmt.Println("  transformers.js")
	fmt.Println("\nNext steps:")
	fmt.Println("  1. Implement getAllUrl() and getDataFromUrl() logic in index.js")
	fmt.Println("  2. Fill in config.js: browser, dataSelectors...")
	fmt.Println("  3. Add transformer functions to transformers.js as needed")

	return nil
}

func render(template string, vars map[string]string) string {
	return tplVarRe.ReplaceAllStringFunc(template, func(match string) string {
		key := match[2 : len(match)-2]
		if val, ok := vars[key]; ok {
			return val
		}
		return ""
	})
}
