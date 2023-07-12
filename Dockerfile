FROM ubuntu:22.04

# Install dependencies
RUN apt update
RUN apt install -y wget curl git unzip

# Install node
RUN curl -fsSL https://deb.nodesource.com/setup_18.x | bash -
RUN apt-get install -y nodejs
RUN node -v

# Install bun
RUN npm install -g bun

# Install go
RUN wget https://go.dev/dl/go1.21rc2.linux-arm64.tar.gz
RUN tar -C /usr/local -xzf go1.21rc2.linux-arm64.tar.gz
RUN rm go1.21rc2.linux-arm64.tar.gz
ENV PATH=$PATH:/usr/local/go/bin
