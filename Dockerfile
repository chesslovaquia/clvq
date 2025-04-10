FROM debian:trixie-20250317-slim

LABEL maintainer="Jeremías Casteglione <jrmsdev@gmail.com>"
LABEL version="250326"

USER root:root
WORKDIR /root

ENV USER root
ENV HOME /root

ENV DEBIAN_FRONTEND noninteractive

ENV APT_INSTALL bash openssl ca-certificates build-essential golang

RUN apt-get clean \
	&& apt-get update -yy \
	&& apt-get dist-upgrade -yy --purge \
	&& apt-get install -yy --no-install-recommends ${APT_INSTALL} \
	&& apt-get clean \
	&& apt-get autoremove -yy --purge \
	&& rm -rf /var/lib/apt/lists/* \
		/var/cache/apt/archives/*.deb \
		/var/cache/apt/*cache.bin

ENV APT_INSTALL_EXTRA media-types less wget

RUN apt-get clean \
	&& apt-get update -yy \
	&& apt-get install -yy --no-install-recommends ${APT_INSTALL_EXTRA} \
	&& apt-get clean \
	&& apt-get autoremove -yy --purge \
	&& rm -rf /var/lib/apt/lists/* \
		/var/cache/apt/archives/*.deb \
		/var/cache/apt/*cache.bin

ARG DEVEL_UID=1000
ARG DEVEL_GID=1000

ENV DEVEL_UID ${DEVEL_UID}
ENV DEVEL_GID ${DEVEL_GID}

RUN groupadd -o -g ${DEVEL_GID} devel \
	&& useradd -o -d /home/devel -m -c 'devel' -g ${DEVEL_GID} -u ${DEVEL_UID} devel \
	&& chmod -v 0750 /home/devel

RUN printf 'umask %s\n' '027' >>/home/devel/.profile
RUN printf "export PS1='%s '\n" '\u@\h:\W\$' >>/home/devel/.profile

RUN install -v -m 0750 -o devel -g devel -d /opt/bin
RUN install -v -m 0750 -o devel -g devel -d /opt/src
RUN install -v -m 0750 -o devel -g devel -d /opt/pkg

RUN install -v -m 0750 -o devel -g devel -d /opt/src/clvq

RUN ln -vsf /opt/bin/clvq /usr/local/bin/clvq
RUN ln -vsf /opt/bin/clvq-site /usr/local/bin/clvq-site

COPY . /opt/src/clvq

ENV SRCD /opt/src/clvq

RUN chown -R devel:devel ${SRCD}

RUN install -v -m 0755 ${SRCD}/docker/user-login.sh /usr/local/bin/user-login.sh

USER devel:devel
WORKDIR /home/devel

ENV USER devel
ENV HOME /home/devel

ENV GOPATH /opt

RUN go version

WORKDIR ${SRCD}
RUN make all

ENTRYPOINT /usr/local/bin/user-login.sh
