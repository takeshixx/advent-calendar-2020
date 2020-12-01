web-build:
	sudo docker build --tag advent-website -f _website/Dockerfile _website

web-up: web-build
	sudo docker save advent-website | pv | ssh xmas sudo docker load

run:
	sudo docker-compose up -d

day01-build:
	sudo docker build --tag day01 -f xmas-cookie2/Dockerfile xmas-cookie2

day01-up: day01-build
	sudo docker save day01 | pv | ssh xmas sudo docker load