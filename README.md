A Simple Dashboard in Htmx & Go

set the respective environment variable refer [.env.example](.env.example)

It uses postgres as its database (set it up)

Once the database is ready perform the migration so the datbase is updated with our schema

```
goose -dir "sql/schema" postgres <DB_URL> up
```

You would require a new user there is function in main.go that creates a user you can uncomment it to create a user

build process first build the css and js using the build.js script (it build the css from tailwindcss and minifies js)

After that build and run the go project
```
go build -o tmp/main .; ./tmp/main
```

This project works even if there is no JS enabled (progressively enhanced)
