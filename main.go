package main

import (
	"embed"
	"errors"
	"flag"
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

var scraperTypeOptions = []string{
	"products/delivery",
	"products/drive",
	"retailOutlets",
	"productDetails",
}

func main() {
	flags := flag.NewFlagSet("scraper-gen", flag.ExitOnError)
	versionFlag := flags.Bool("version", false, "print version")
	vFlag := flags.Bool("v", false, "print version")
	siteFlag := flags.String("site", "", "site name (e.g. amazon)")
	localeFlag := flags.String("locale", "", "locale (e.g. fr, co.uk)")
	typeFlag := flags.String("type", "", "comma-separated scraper types")
	dryRun := flags.Bool("dry-run", false, "print what would be generated without writing files")

	_ = flags.Parse(os.Args[1:])

	if *versionFlag || *vFlag {
		fmt.Println(version)
		os.Exit(0)
	}

	site := strings.TrimSpace(*siteFlag)
	locale := strings.TrimSpace(*localeFlag)
	var scraperTypes []string
	if *typeFlag != "" {
		for _, t := range strings.Split(*typeFlag, ",") {
			scraperTypes = append(scraperTypes, strings.TrimSpace(t))
		}
	}

	// Non-interactive path: all flags provided
	if site != "" && locale != "" && len(scraperTypes) > 0 {
		runGenerate(site, locale, scraperTypes, *dryRun)
		return
	}

	// Interactive path
	var scraperTypesInteractive []string

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
			huh.NewMultiSelect[string]().
				Title("Scraper type(s)").
				Options(
					huh.NewOption("products/delivery", "products/delivery"),
					huh.NewOption("products/drive", "products/drive"),
					huh.NewOption("retailOutlets", "retailOutlets"),
					huh.NewOption("productDetails", "productDetails"),
				).
				Validate(func(v []string) error {
					if len(v) == 0 {
						return errors.New("select at least one scraper type")
					}
					return nil
				}).
				Value(&scraperTypesInteractive),
		),
	)

	if err := form.Run(); err != nil {
		fmt.Println("\nAborted.")
		os.Exit(0)
	}

	runGenerate(strings.TrimSpace(site), strings.TrimSpace(locale), scraperTypesInteractive, *dryRun)
}

func runGenerate(site, locale string, scraperTypes []string, dryRun bool) {
	if dryRun {
		fmt.Println("\nWould create:")
		for _, t := range scraperTypes {
			base := filepath.Join("src", "sites", site, locale, t)
			fmt.Printf("  %s/config.js\n", base)
			fmt.Printf("  %s/index.js\n", base)
			fmt.Printf("  %s/transformers.js\n", base)
		}
		return
	}

	// Check all target paths before writing anything
	for _, t := range scraperTypes {
		targetPath := filepath.Join("src", "sites", site, locale, t)
		if _, err := os.Stat(targetPath); err == nil {
			fmt.Fprintf(os.Stderr, "\nDirectory already exists: %s\nAborting to avoid overwriting existing scraper.\n", targetPath)
			os.Exit(1)
		}
	}

	for _, t := range scraperTypes {
		if err := generate(site, locale, t); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	}
}

func generate(site, locale, scraperType string) error {
	targetPath := filepath.Join("src", "sites", site, locale, scraperType)

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
