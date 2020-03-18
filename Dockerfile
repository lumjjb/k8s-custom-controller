FROM ubuntu:18.04
RUN apt update
RUN apt install -y curl
RUN curl -LO https://storage.googleapis.com/kubernetes-release/release/`curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt`/bin/linux/amd64/kubectl
RUN chmod +x kubectl
COPY isecl-k8s-controller-1.0-SNAPSHOT /isecl-k8s-controller-1.0-SNAPSHOT
ENTRYPOINT ["isecl-k8s-controller-1.0-SNAPSHOT"]