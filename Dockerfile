FROM aecgeeks/ifcopenshell:latest AS develop
WORKDIR /work

USER root

RUN sudo apt update -y && apt-get install -y curl
RUN curl -fsSL https://golang.org/dl/go1.21.5.linux-amd64.tar.gz | tar -C /usr/local -xz
ENV PATH $PATH:/usr/local/go/bin

CMD /bin/bash
