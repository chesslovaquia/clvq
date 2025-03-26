#!/bin/sh
set -eu
exec docker run -it --rm -u devel \
	--name clvq \
	--hostname clvq.local \
	-v "${PWD}:/opt/go/src/clvq" \
	--workdir /opt/go/src/clvq \
	-p 127.0.0.1:8044:8044 \
	chesslovaquia/clvq
