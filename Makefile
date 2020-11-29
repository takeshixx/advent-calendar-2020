web-build:
	sudo docker build --tag advent-website -f website/Dockerfile website

web-up:
	sudo docker save advent-website | pv | ssh xmas sudo docker load

run:
	sudo docker-compose up -d