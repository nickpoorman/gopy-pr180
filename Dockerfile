FROM ubuntu:18.04

# Tools
RUN apt-get update && apt-get install -y \
    g++ \
    gdb \
    git \
    make \
    vim \
    wget \
    python3 python3-dev \
    python3-pip python3-venv python3-wheel python3-setuptools \
    build-essential \
    pkg-config \
    libssl-dev libffi-dev \
    && rm -rf /var/lib/apt/lists/*

# Go installation
RUN cd /tmp && \
    wget https://dl.google.com/go/go1.12.7.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.12.7.linux-amd64.tar.gz && \
    rm go1.12.7.linux-amd64.tar.gz
ENV PATH="/usr/local/go/bin:${PATH}"

RUN python3 -m pip install --upgrade \
    Cython \
    ipython \
    numpy \
    pyarrow \
    pybindgen \
    setuptools \
    wheel

# link the pkgconfig files we need
RUN rm -rf /usr/lib/pkgconfig && \
    ln -s /usr/lib/x86_64-linux-gnu/pkgconfig /usr/lib/pkgconfig

# gopy
ENV GO111MODULE=off
RUN go get golang.org/x/tools/cmd/goimports && \
    go get github.com/go-python/gopy && \
    cd ~/go/src/github.com/go-python/gopy && \
    git fetch origin pull/180/head:pr180 && \
    git checkout pr180 && \
    go install
ENV PATH="/root/go/bin:${PATH}"

# Python needs to know where our .so files are.
ENV LD_LIBRARY_PATH=/out

WORKDIR /src/gopy-pr180