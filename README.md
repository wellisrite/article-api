
# Article API

The Article API is a simple web service that allows you to manage articles and perform various operations such as fetching articles and creating new articles.

## Table of Contents

- [Description](#description)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Running](#running)
- [Usage](#usage)
- [Endpoints](#endpoints)
- [Contributing](#contributing)
- [License](#license)

## Description

The Article API provides HTTP endpoints to interact with articles. It supports creating new articles and retrieving a list of articles. Articles are stored in a database and cached for better performance.

## Getting Started

### Prerequisites

- Docker: Ensure you have Docker installed on your machine.

### Installation

1. Clone this repository:

   ```bash
   git clone https://github.com/yourusername/article-api.git
   cd article-api
   ```
   
Running
Create a .env file in the project root directory and provide values for the environment variables:

```plaintext
ARTICLE_API_APP_ENV=
ARTICLE_API_APP_SECRET=
ARTICLE_API_APP_API_PRIVATE_KEY=
ARTICLE_API_APP_API_PUBLIC_KEY=
ARTICLE_API_APP_VERSION=
ARTICLE_API_APP_CORS_DOMAIN=

# WEBSERVER
ARTICLE_API_WEBSERVER_LISTEN_ADDRESS=
ARTICLE_API_WEBSERVER_MAX_CONNECTION=
ARTICLE_API_WEBSERVER_READ_TIMEOUT=
ARTICLE_API_WEBSERVER_WRITE_TIMEOUT=
ARTICLE_API_WEBSERVER_GRACEFUL_TIMEOUT=

# DB
ARTICLE_API_DB_DIALECT=
ARTICLE_API_DB_HOST= 
ARTICLE_API_DB_USER=
ARTICLE_API_DB_PASSWORD=
ARTICLE_API_DB_NAME=
ARTICLE_API_DB_PORT=
ARTICLE_API_DB_SSL_MODE=
ARTICLE_API_DB_MAX_CONNECTION=
ARTICLE_API_DB_MAX_IDLE_CONNECTION=
ARTICLE_API_DB_MAX_LIFETIME_CONNECTION=

# REDIS
ARTICLE_API_REDIS_HOST=
ARTICLE_API_REDIS_PORT=
ARTICLE_API_REDIS_PASSWORD=
```

Start the services using Docker Compose:

```bash
  Copy code
  docker-compose up
  This will start the Article API service and a Redis container for caching.
```

Access the API at http://localhost:4002.

Usage
Once the Article API and Redis container are up and running, you can interact with the API using HTTP requests. Here are some example requests:

Create a new article:

```bash
curl -X POST http://localhost:4002/articles -H "Content-Type: application/json" -d '{"author": "John", "title": "My Article", "body": "This is the content of the article"}'
```

Get articles filtered by author:

```bash
curl http://localhost:4002/articles?author=John
```

Get articles filtered by author and query:

```bash
curl http://localhost:4002/articles?author=John&query=Sample
Endpoints
```

POST /articles: Create a new article.

Request Body:

```json
Copy code
{
  "author": "John",
  "title": "My Article",
  "body": "This is the content of the article"
}
```

GET /articles: Get a list of articles. Optional query parameters: author to filter by author.

Contributing
Contributions are welcome! Feel free to open issues and pull requests for improvements or bug fixes.

License
This project is licensed under the MIT License - see the LICENSE file for details.
