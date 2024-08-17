# Chatley
Description:


# Instructions
1. Create a .env file in the root folder of the project
2. Add the port you wish to use in the .env file. Example = 'PORT="8080"'  
3. Run the command 'openssl rand -base64 64' in your terminal and place the output in the .env file like this:
JWT_SECRET=insert_key_here
4. Run 'go build && ./chatley' in terminal

Installation:
1. Run the following commands in your terminal:
        'go get "github.com/gorilla/mux"' and 
        'go get "github.com/joho/godotenv' in the terminal
        'go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest'
        'go install github.com/pressly/goose/v3/cmd/goose@latest'
        
2. Create an .env file and add the following lines to it:
        PORT="(insert port number you will be using here)"
        DB_URL="(insert database URL for the database you will be using)"
        TEST_DB_URL="(insert database URL for the database you will be using for testing)"