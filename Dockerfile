FROM python:3.7
MAINTAINER Arun Kumar<tigeraniya@gmail.com>

ADD . /code
RUN pip3 install -r /code/proj/requirements.txt

EXPOSE 8000

WORKDIR /code/proj

CMD ["python", "server.py"]
