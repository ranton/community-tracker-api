# community-tracker-api

# How to run the Backend Service using Docker

### Environment Setup
Create copy of .env.example and name it .env \
Complete the env values before running the container\
for the first time.\
Running the container with incomplete .env values \
will cause docker to use default values for the database.



### How to start the docker container

In the project directory, you can run:

### `docker compose up`

Runs the service.

### `docker compose up --build`

Use the build flag to build the container before running the service. \
Use this command to to apply code changes. \
Stop the container for before rebuilding.


### `docker compose up -d`

Use the d flag to run the service as a daemon (backgound application). \


### `docker compose down`

Use this command to stop the container.

## Authentication Setup
To able to use the JWT authentication, \
a jwt secret should be generated. \
The key for it on .env is JWT_SECRET \
You can get a sign key in https://www.grc.com/passwords.htm \
use 64 random hexadecimal characters