FROM opensuse:tumbleweed

RUN zypper --non-interactive ref && zypper --non-interactive dup && zypper --non-interactive in python3 python3-devel libpcap-devel python3-numpy python3-Pillow

COPY step1 /
COPY step0 /
COPY ring /
COPY run.sh /
COPY run_single.sh /

CMD ["/run.sh"]
