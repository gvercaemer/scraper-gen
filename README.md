# scraper-gen

Interactive CLI to scaffold Click2buy scrapers.

Generates the boilerplate structure for a new scraper — `config.js`, `index.js`, and `transformers.js` — from a guided prompt.

## Requirements

- macOS (arm64 or x86_64)
- A Click2buy scraper project with a `src/sites/` directory

## Installation

```sh
brew tap gvercaemer/tap
brew install scraper-gen
```

## Usage

Run from the root of your scraper project:

```sh
scraper-gen
```

You'll be prompted for:

- **Site name** — e.g. `amazon`, `kaufland`
- **Locale** — e.g. `fr`, `co.uk`, `com.br`
- **Scraper type** — one of `products/delivery`, `products/drive`, `retailOutlets`, `productDetails`

The generated files are placed at:

```
src/sites/{site}/{locale}/{type}/
├── config.js
├── index.js
└── transformers.js
```

## Upgrade

```sh
brew upgrade scraper-gen
```
