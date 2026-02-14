# Introduction
### General overview
This is a repo for a website for my band. The initial goal is to sell merch, but as I go I will be adding things like email updates, tour dates, music videos and more! 

On the more technical side, this project is also to showcase my skills in CI/CD, automated testing, SQL, and of course full stack development.

## Tech stack

#### Backend
- Framework: None, golang's standard libraries
- ORM: None, I write raw SQL then use SQLC to generate type-safe Go code to interact with the edatabase
- Database: Postgresql
- Deployment environment: Ubuntu server 24

#### Frontend
- Framework: Vue3 + Typescript
- State management: Pinia
- Styles: Vanilla CSS
- UI component library: PrimeVue

#### Test and CI/CD
- Test and deployment pipeline: Github Actions
- Api testing: Bruno
- End-to-end testing: Cypress
- Unit tests: Go standard testing library

## Setup
Clone the repo
```bash
git clone <repo_url>
```

### Backend
Check out the Go docs for installation instructions: https://go.dev/doc/install

Change to the backend directory
```bash 
cd backend
``` 
Fetch Go dependencies
```bash 
go get
```
Install goose and sqlc
```bash 
go install github.com/pressly/goose/v3/cmd/goose@latest
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```
You will need to create bo your .env file and docker-compose.yml files. I have provided examples for both in the repository with an extra **.example** extension. Copy both of these files and enter the relevant information. 

Once done, you should be able to run 
```bash
docker-compose up -d
```

Now from your backend directory, run 
```bash 
go run .
```

This should start your backend server at http://localhost:8080 (or whatever port you specified)

### Frontend

This one is a bit easier. Change directories to your frontend directory (I am assuming you are still in the backend directory)
```bash 
cd ../frontend
```  
Install the frontend dependencies
```bash 
npm install
```  

Run the dev server
```bash 
npm run dev
```  

### Testing
Coming soon! 