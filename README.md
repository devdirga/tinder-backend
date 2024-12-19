## Details on the structure of the service

### Authentication (`auth`)
Handles user registration, login, and session management. Utilizes JWT for secure authentication.

**Endpoints:**
- `POST /signup`: User registration.
- `POST /signin`: User login.
- `GET  /verification/:token`: Verification registration.

### User Profile (`profile`)
Manages user profiles, including viewing and editing personal information.

**Endpoints:**
- `GET /me`: Retrieve user profile.
- `POST /me`: Update user profile.

### Swipe Data (`swipe`)
Implements the swiping feature, allowing users to get profile.

**Endpoints:**
- `GET /swipe`: Load one data profile.

### Swiping (`swipe`)
Implements the swiping feature, allowing users to like or pass on other users.

**Endpoints:**
- `POST /swipe`: Like/Pass a user.

### Matchmaking (`match`)
Handles the logic for determining matches between users.
And send email notification


## Instructions on how to run the service

### Prerequisites
Ensure you have the following installed:
- **Go** (1.2 or later)
- **PostgreSQL** (v13 or later)
- **Git**
- **Docker** (optional, for message queue)

## Installation

### Clone the Repository
```bash
git clone https://github.com/devdirga/tinder-backend.git
cd tinder-service
```
---

### Environment Configuration
Create a `.config.json` file in the project root and populate it with the following variables:
```bash
{
  "DB": "host=localhost user=postgres password=mysecretpassword dbname=tinder port=5432 sslmode=disable",
  "IsDebug": true,
  "IsQueue": true,
  "IsConcurrent":true,
  "Secret": "secret",
  "GoogleSmtpKey": "azdg rkiv wnqe vuil ",
  "URL": "http://localhost:5000/",
  "Quota": 10,
  "KafkaUrl": "localhost:9092",
  "KafkaTopic": "test-topic"
}
```

### Backend Setup
Install Go dependencies:
```bash
go mod tidy
```

Run the backend service:
```bash
go run main.go
```
The service will run on `http://localhost:5000` by default.

## Deployment

CI/CD using Github Actions

Create new repository secrets to add your server access
```bash
VPS_HOST=
VPS_PRIVATE_KEY=
VPS_USER=
```

Setup your path in your server by change in file ```bash .github/workflows/deploy.yml  ```

Every time you commit to the repository, GitHub Actions will build and deploy to your server, using SSH to copy the Golang binary file.

## Implementation of messsage broker (KafKa)
1. Create a docker-compose.yml file
```bash
version: '3.8'
services:
  zookeeper:
    image: confluentinc/cp-zookeeper:latest
    container_name: zookeeper
    environment:
      ZOOKEEPER_CLIENT_PORT: 2181
      ZOOKEEPER_TICK_TIME: 2000
    ports:
      - "2181:2181"

  kafka:
    image: confluentinc/cp-kafka:latest
    container_name: kafka
    depends_on:
      - zookeeper
    ports:
      - "9092:9092"
    environment:
      KAFKA_BROKER_ID: 1
      KAFKA_ZOOKEEPER_CONNECT: zookeeper:2181
      KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://localhost:9092
      KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR: 1
```

2. Run the following command in the directory where the docker-compose.yml is located:
```bash
docker-compose up -d
```

3. Access the Kafka container:
```bash
docker exec -it kafka bash
```

4. Create a topic:
```bash
kafka-topics --create --topic test-topic --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1
```

