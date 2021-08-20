FROM ubuntu:18.04
RUN apt-get update -qq && DEBIAN_FRONTEND="noninteractive" && apt-get install --no-install-recommends -qqy \
     build-essential \
     curl \
     htop \
     unzip \
     wget \
     vim \
     net-tools \
     lsb-release \
     git \
     sudo \
     ca-certificates \
     supervisor

RUN mkdir -p /var/log/supervisor
COPY supervisord.conf /etc/supervisor/conf.d/supervisord.conf


# Create user
ARG user_id=1000
ARG user_name=developer
ARG groups="developer:1000"
ARG home=/home/developer
RUN echo "ALL    ALL = (ALL) NOPASSWD: ALL" > /etc/sudoers && mkdir -p "$(dirname $home)"
RUN useradd --shell /bin/bash -u $user_id -o -c "" -m $user_name -d $home
RUN usermod -aG root $user_name

WORKDIR /app/
COPY . /app/
RUN chown -R developer:developer /app

ENV USER=$user_name
USER $user_name

RUN /bin/bash /app/utils/go-protoc-install.sh 2>&1 /dev/null
RUN mkdir -p /app/logs
RUN mkdir -p /app/state

EXPOSE 4101
EXPOSE 4102
#CMD ["/bin/bash"]
CMD ["/usr/bin/supervisord"]
