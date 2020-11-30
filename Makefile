web-build:
	sudo docker build --tag advent-website -f _website/Dockerfile _website

web-up: web-build
	sudo docker save advent-website | pv | ssh xmas sudo docker load

run:
	sudo docker-compose up -d