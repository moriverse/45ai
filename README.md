# AI Portrait Generator

A full-stack application for generating AI-powered artistic portraits, built with Uni-app (frontend) and Golang (backend).

## Project Structure

```
.
├── frontend/           # Uni-app frontend (WeChat Mini Program)
└── backend/           # Golang backend
```

## Frontend (Uni-app)

The frontend is built using:
- Uni-app (Vue 3 Syntax)
- unibest boilerplate
- Pinia for state management
- UnoCSS for styling

### Setup

1. Install dependencies:
```bash
cd frontend
npm install
```

2. Run development server:
```bash
npm run dev:mp-weixin
```

3. Open WeChat DevTools and import the project from the `dist/dev/mp-weixin` directory.

## Backend (Golang)

The backend is built using:
- Golang 1.21+
- Gin web framework
- GORM for database operations
- MySQL 8.0+

### Setup

1. Install dependencies:
```bash
cd backend
go mod download
```

2. Configure the application:
- Copy `configs/config.example.yaml` to `configs/config.yaml`
- Update the configuration with your credentials

3. Run the server:
```bash
go run cmd/server/main.go
```

## Features

- User authentication via WeChat
- Template selection
- Photo upload and AI generation
- Payment integration with WeChat Pay
- Task history and management

## Development

### Prerequisites

- Node.js 16+
- Go 1.21+
- MySQL 8.0+
- WeChat Developer Account
- WeChat Pay Merchant Account

### Environment Variables

Create a `.env` file in the backend directory with the following variables:

```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=your_user
DB_PASSWORD=your_password
DB_NAME=ai_photo
JWT_SECRET=your_jwt_secret
WECHAT_APP_ID=your_app_id
WECHAT_MCH_ID=your_mch_id
WECHAT_API_V3_KEY=your_api_v3_key
```

## License

MIT
