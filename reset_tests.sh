cd sql/schema
goose postgres postgres://brendancoen:@localhost:5432/gator?sslmode=disable down
goose postgres postgres://brendancoen:@localhost:5432/gator?sslmode=disable up
cd ../..
