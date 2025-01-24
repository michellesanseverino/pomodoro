# pomodoro

## What is the project
    A pomodoro site/app where you can time your focus time and also play music from youtube, maintaining only one tab open in the browser. It will be able to personalize the focus and rest time, the number of sessions, the genre of the music. 
    For organization, it will feature a to do list and tags to be able to manage time spend in each place

## Architecture

pomodoro/
├── cmd/
│   └── server/             # Entry point for the application
│       └── main.go         # Main application file
├── config/                 # Configuration files (e.g., environment variables)
│   └── config.go           # Configuration loading logic
├── internal/               # Internal application logic (not exposed externally)
│   ├── api/                # API handlers and routes
│   │   ├── pomodoro.go     # Handlers for Pomodoro logic
│   │   ├── youtube.go      # Handlers for YouTube integration
│   │   └── routes.go       # API routes definition
│   ├── services/           # Business logic and service layer
│   │   ├── pomodoro.go     # Core Pomodoro timer logic
│   │   └── youtube.go      # Logic for interacting with YouTube API
│   ├── models/             # Data models
│   │   ├── pomodoro.go     # Pomodoro session model
│   │   └── user.go         # User model
│   ├── repository/         # Database interaction logic
│   │   ├── pomodoro.go     # Pomodoro session persistence
│   │   └── user.go         # User data persistence
│   └── utils/              # Utility functions
│       ├── logger.go       # Logging utilities
│       └── http.go         # HTTP utilities
├── pkg/                    # Public reusable packages
│   └── youtube/            # YouTube API wrapper
│       ├── client.go       # YouTube API client
│       └── auth.go         # YouTube authentication logic
├── web/                    # Frontend assets and templates
│   ├── static/             # Static files (CSS, JS, images)
│   │   ├── css/
│   │   └── js/
│   ├── templates/          # HTML templates
│   │   ├── base.html       # Base layout
│   │   └── index.html      # Home page
│   └── public/             # Public assets
│       ├── favicon.ico
│       └── manifest.json
├── scripts/                # Helper scripts (e.g., for setup, deployment)
│   ├── start.sh            # Start server script
│   └── deploy.sh           # Deployment script
├── test/                   # Tests
│   ├── integration/        # Integration tests
│   │   ├── api_test.go
│   │   └── youtube_test.go
│   └── unit/               # Unit tests
│       ├── pomodoro_test.go
│       └── youtube_test.go
├── .env                    # Environment variables
├── .gitignore              # Git ignore file
├── go.mod                  # Go module file
├── go.sum                  # Go module dependencies
└── README.md               # Project documentation


## Tecnologies to be used
    - Database 
    - Integration
    - Figma
    - Golang
    - HTML/CSS/JavaScript
    - React


