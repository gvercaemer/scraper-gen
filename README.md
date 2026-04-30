# scraper-gen

> Interactive CLI to scaffold Click2buy scrapers in seconds.

Generates the boilerplate structure for a new scraper — `config.js`, `index.js`, and `transformers.js` — from a guided prompt or a single command.

---

## Requirements

- macOS (arm64 or x86_64)
- A Click2buy scraper project with a `src/sites/` directory

---

## Installation

```sh
brew tap gvercaemer/tap
brew install scraper-gen
```

---

## Usage

### Interactive mode

Run from the root of your scraper project:

```sh
scraper-gen
```

You'll be prompted for:

| Field | Examples |
|---|---|
| **Site name** | `amazon`, `kaufland`, `carrefour` |
| **Locale** | `fr`, `co.uk`, `com.br` |
| **Scraper type(s)** | pick one or more with `space`, confirm with `enter` |

Scraper types available:

- `products/delivery`
- `products/drive`
- `retailOutlets`
- `productDetails`

### Non-interactive mode

Skip the prompts entirely by passing flags — useful for scripts and CI:

```sh
scraper-gen --site=amazon --locale=fr --type=retailOutlets,products/drive
```

| Flag | Description |
|---|---|
| `--site` | Site name |
| `--locale` | Locale |
| `--type` | Comma-separated list of scraper types |
| `--dry-run` | Preview what would be generated without writing any files |
| `--version` | Print the current version |

### Dry run

See what would be created before committing:

```sh
scraper-gen --site=amazon --locale=fr --type=retailOutlets --dry-run

# Would create:
#   src/sites/amazon/fr/retailOutlets/config.js
#   src/sites/amazon/fr/retailOutlets/index.js
#   src/sites/amazon/fr/retailOutlets/transformers.js
```

---

## Output structure

Generated files are placed at:

```
src/sites/{site}/{locale}/{type}/
├── config.js        ← name, model, url, dataSelectors
├── index.js         ← getAllUrl() and getDataFromUrl() scaffolds
└── transformers.js  ← one transformer stub per field
```

Multiple types selected in a single run each get their own directory:

```
src/sites/amazon/fr/
├── retailOutlets/
│   ├── config.js
│   ├── index.js
│   └── transformers.js
└── products/
    └── drive/
        ├── config.js
        ├── index.js
        └── transformers.js
```

---

## Upgrade

```sh
brew update && brew upgrade scraper-gen
```

---

## Development

Go 1.22+ is required to build from source:

```sh
git clone https://github.com/gvercaemer/scraper-gen
cd scraper-gen
go build .
```

Release binaries are built and published automatically via GitHub Actions when a new tag is pushed.
