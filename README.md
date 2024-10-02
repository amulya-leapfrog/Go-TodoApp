- Clone the repository and open it in your code editor.
- Run the 'docker-compose up' command in your project's root directory terminal to pull the Postgres Docker image.
- Connect your DB client to the Postgres server.
- Create a .env file referencing .env.example.
- Migrate the database schema by running the migration with Goose: goose postgres "user=<username> password=<password> dbname=<database name> host=<host> port=<port>" up. You can also use your preferred migration library.
- Run the app using the 'go run main.go' command.




