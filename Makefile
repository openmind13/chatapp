# start postgres db server and run application
start:
	clear && sudo service postgresql start && go build -o webchat *.go && ./webchat

# run application
build:
	clear && go build -o webchat *.go && ./webchat

start_db_server:
	sudo service postgresql start

setup_db:
	psql -f ./models/setupdb.sql



.DEFAULT_GOAL := build