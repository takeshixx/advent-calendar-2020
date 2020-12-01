web-build:
	sudo docker build --tag advent-website -f _website/Dockerfile _website

web-up: web-build
	sudo docker save advent-website | pv | ssh xmas sudo docker load

run:
	sudo docker-compose up -d

sync:
	scp docker-compose.yml xmas:

day01-build:
	sudo docker build --tag day01 -f xmas-cookie2/Dockerfile xmas-cookie2

day01-up: day01-build
	sudo docker save day01 | pv | ssh xmas sudo docker load

day02-build:
	sudo docker build --tag day02 -f dtls/Dockerfile dtls

day02-up: day02-build
	sudo docker save day02 | pv | ssh xmas sudo docker load

day03-build:
	sudo docker build --tag day03 -f elf/Dockerfile elf

day03-up: day03-build
	sudo docker save day03 | pv | ssh xmas sudo docker load