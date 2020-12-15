web-build:
	sudo docker build --tag advent-website -f _website/Dockerfile _website

web-up: web-build sync
	sudo docker save advent-website | pv | ssh xmas sudo docker load

web: web-up
	ssh xmas "sudo docker-compose stop website && sudo docker-compose up -d website"

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

day04-build:
	sudo docker build --tag day04 -f xmas-socks/Dockerfile xmas-socks

day04-up: day04-build
	sudo docker save day04 | pv | ssh xmas sudo docker load

day05-build:
	sudo docker build --tag day05 -f proto/Dockerfile proto

day05-up: day05-build
	sudo docker save day05 | pv | ssh xmas sudo docker load

day06-build:
	sudo docker build --tag day06 -f xmas-cloud/Dockerfile xmas-cloud

day06-up: day06-build
	sudo docker save day06 | pv | ssh xmas sudo docker load
	scp -r xmas-cloud/flag xmas:data/day06/

day07-build:
	sudo docker build --tag day07 -f PCAP_poly/Dockerfile PCAP_poly

day07-up: day07-build
	sudo docker save day07 | pv | ssh xmas sudo docker load

day08-build:
	sudo docker build --tag day08 -f HSFZ/Dockerfile HSFZ

day08-up: day08-build
	sudo docker save day08 | pv | ssh xmas sudo docker load

day09-build:
	sudo docker build --tag day09 -f xmas-from/Dockerfile xmas-from

day09-up: day09-build
	sudo docker save day09 | pv | ssh xmas sudo docker load

day10-build:
	sudo docker build --tag day10 -f redstar/Dockerfile redstar

day10-up: day10-build
	sudo docker save day10 | pv | ssh xmas sudo docker load

day11-build:
	sudo docker build --tag day11 -f xmas-karaoke/Dockerfile xmas-karaoke

day11-up: day11-build
	sudo docker save day11 | pv | ssh xmas sudo docker load

day12-build:
	sudo docker build --tag day12 -f xmasgreetings/Dockerfile xmasgreetings

day12-up: day12-build
	sudo docker save day12 | pv | ssh xmas sudo docker load

day13-build:
	sudo docker build --tag day13 -f xmas-webasm/Dockerfile xmas-webasm

day13-up: day13-build
	sudo docker save day13 | pv | ssh xmas sudo docker load

day14-build:
	sudo docker build --tag day14 -f WebRace/Dockerfile WebRace

day14-up: day14-build
	sudo docker save day14 | pv | ssh xmas sudo docker load

day15-build:
	sudo docker build --tag day15 -f nts/Dockerfile nts

day15-up: day15-build
	sudo docker save day15 | pv | ssh xmas sudo docker load

day16-build:
	sudo docker build --tag day16 -f ip-https/Dockerfile ip-https

day16-up: day16-build
	sudo docker save day16 | pv | ssh xmas sudo docker load

#####################################

day01-restart:
	ssh xmas "sudo docker-compose stop day01 && sudo docker-compose up -d day01"

day02-restart:
	ssh xmas "sudo docker-compose stop day02 && sudo docker-compose up -d day02"

day03-restart:
	ssh xmas "sudo docker-compose stop day03 && sudo docker-compose up -d day03"

day04-restart:
	ssh xmas "sudo docker-compose stop day04 && sudo docker-compose up -d day04"

day05-restart:
	ssh xmas "sudo docker-compose stop day05 && sudo docker-compose up -d day05"

day06-restart:
	ssh xmas "sudo docker-compose stop day06 && sudo docker-compose up -d day06"

day07-restart:
	ssh xmas "sudo docker-compose stop day07 && sudo docker-compose up -d day07"

day08-restart:
	ssh xmas "sudo docker-compose stop day08 && sudo docker-compose up -d day08"

day09-restart:
	ssh xmas "sudo docker-compose stop day09 && sudo docker-compose up -d day09"

day10-restart:
	ssh xmas "sudo docker-compose stop day10 && sudo docker-compose up -d day10"

day11-restart:
	ssh xmas "sudo docker-compose stop day11 && sudo docker-compose up -d day11"

day12-restart:
	ssh xmas "sudo docker-compose stop day12 && sudo docker-compose up -d day12"

day13-restart:
	ssh xmas "sudo docker-compose stop day13 && sudo docker-compose up -d day13"

day14-restart:
	ssh xmas "sudo docker-compose stop day14 && sudo docker-compose up -d day14"

day15-restart:
	ssh xmas "sudo docker-compose stop day15 && sudo docker-compose up -d day15"

day16-restart:
	ssh xmas "sudo docker-compose stop day16 && sudo docker-compose up -d day16"

day17-restart:
	ssh xmas "sudo docker-compose stop day17 && sudo docker-compose up -d day17"

day18-restart:
	ssh xmas "sudo docker-compose stop day18 && sudo docker-compose up -d day18"

day19-restart:
	ssh xmas "sudo docker-compose stop day19 && sudo docker-compose up -d day19"

day20-restart:
	ssh xmas "sudo docker-compose stop day20 && sudo docker-compose up -d day20"

day21-restart:
	ssh xmas "sudo docker-compose stop day21 && sudo docker-compose up -d day21"

day22-restart:
	ssh xmas "sudo docker-compose stop day22 && sudo docker-compose up -d day22"

day23-restart:
	ssh xmas "sudo docker-compose stop day23 && sudo docker-compose up -d day23"

day24-restart:
	ssh xmas "sudo docker-compose stop day24 && sudo docker-compose up -d day24"