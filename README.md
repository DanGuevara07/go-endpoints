# go-endpoints

# Steps to Run

1. Install dependencies with npm i
2. Copy .env.example in .env new file
3. Start database in a docker container with command docker-compose up -d or docker-compose up
4. Run the command make start, to start service
5. Access the URL http://0.0.0.0:3000/local/go-endpoints/users/ to try endpoints
    - GET /users will show array of users
    - POST /users with JSON containing "name" field will create a new user
    - DELETE /users/{id} will delete user if id exist on database
    - PATCH /users/{id} will update user if id exist and a JSON with "name" field is sent
    - GET /users/{id} will get an user if the id exist on database.
    - User fields:
        - ID
        - Created at
        - Updated at
        - name
        Only name can be changed, the other 3 fields are generated when a user is created.