FROM envoyproxy/envoy-dev:3b45d6cf465b8ce3c268edaa5c70722811dbb39f

RUN apt-get update && apt-get install -y gettext && rm -rf /var/lib/apt/lists/*

COPY . /etc/envoy/

RUN chmod go+r /etc/envoy/

RUN chmod +x /etc/envoy/entry-point.sh

RUN chmod -R 777 /etc/envoy

CMD ["/etc/envoy/entry-point.sh"]
