.PHONY: docker-build docker-run build-hi

docker-build:
	docker build -t gopy/pr180 .

docker-run:
	docker run \
	-v $(PWD):/src/gopy-pr180 \
	-v $(PWD)/out:/out \
	-it --rm --workdir=/src/gopy-pr180 \
	gopy/pr180

build-hi:
	gopy build -vm="python3" -output=/out github.com/go-python/gopy/_examples/hi
