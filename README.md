# Payment API
API for manage transactions, payments and accounts

Powered by:
**IrisGo** - https://github.com/kataras/iris
**GORM** - https://github.com/jinzhu/gorm

# Directory structure

**main.go** -> Application bootstraper
**Dockerfile**
**docker-compose.yml**
**app/**
	**-- http/**
		**---- controllers/** -> All application controllers
		**---- resources/** -> Requests/Responses representations
	**-- adapters/** -> Adapters for models/resources conversion and vice-versa
	**-- models/** -> Database entities
	**-- repositories/** -> Application repositories
	**-- services/** -> Application services
**infra/**
	
# Running the application

 1. Clone the repository
 2. Run the command:
 

    *docker-compose up --build*
3. For database check  with external tool, connect on port 3307, exposed by the container. No volumes were created, but can be adjusted on docker-compose.yml if needed
