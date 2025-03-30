# e-pocket

e-pocket is a digital wallet application that allows users to make payments seamlessly.

## ✨ Key Features

- **Payments**: Make payments directly from the app to various provider.
- **Reports**: Generate detailed reports to analyze your financial habits.

## 🚀 Quick Start

### Prerequisites
- Docker & Docker Compose
- Node.js (v18+)
- PostgreSQL (via Docker)

### Installation
```bash
git clone https://github.com/stephanvebrian/e-pocket.git
cd e-pocket
docker-compose up --build
```

### 🛠 Development

#### Services Overview

#### Project Structure
```text
e-pocket/
├── pocket-ui/                # Next.js frontend
├── pocket-engine/            # Go Backend service
├── pocket-db/                # PostgreSQL data & migrations
└── docker-compose.<env>.yml  # Orchestration
```

## Apps Notes
### 🖥 pocket-ui (Next.js)

#### WSL2 Hot Reload Issues
Due to filesystem limitations in WSL2, Hot Module Replacement (HMR) may break. We've implemented these workarounds:

- **`WATCHPACK_POLLING` enabled** in Docker environment
- **Related Issues**:
  - [next.js#36774](https://github.com/vercel/next.js/issues/36774)
  - [WSL#4739](https://github.com/microsoft/WSL/issues/4739#issuecomment-534049240)


