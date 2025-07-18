# NestMate - Personal Productivity App

A unified, cross-platform personal productivity and life management application that combines expense tracking, to-do management, and note-taking capabilities.

## Project Structure

```
nestmate/
├── backend/          # Go backend with clean architecture
├── web/             # React web application
├── mobile/          # React Native mobile application
├── docker/          # Docker configurations
├── .github/         # GitHub Actions CI/CD
└── docs/            # Documentation
```

## Getting Started

### Prerequisites

- Go 1.21+
- Node.js 18+
- Docker & Docker Compose
- React Native CLI (for mobile development)

### Development Setup

1. Clone the repository
2. Run `docker-compose up -d` to start development environment
3. Follow individual module setup instructions in their respective directories

## Architecture

The application follows clean architecture principles with:
- **Backend**: Go with clean architecture (domain, application, infrastructure, interfaces)
- **Web**: React with TypeScript
- **Mobile**: React Native with TypeScript
- **Database**: PouchDB (local) + Firestore (cloud sync)
- **Authentication**: Firebase Auth

## Modules

- **Expense Tracking**: Income/expense management with categorization
- **PDF Parser**: Automatic bank statement parsing
- **To-Do Management**: Task organization with labels and priorities
- **Notes**: Markdown-based note-taking with search and tags