FROM python:3

RUN apt-get update -y
RUN apt-get install python3 python3-pip python3-dev git libssl-dev libffi-dev build-essential netcat -y
RUN python3 -m pip install --upgrade pip
RUN python3 -m pip install --upgrade pipenv

ENV SECRET "w0r2DmJz9ahZluVyajphYRNkDs864xrL"
ENV IS_PRODUCTION True

ADD webapp /webapp

WORKDIR /webapp

RUN pipenv install 
RUN pipenv run python source/manage.py makemigrations && pipenv run python source/manage.py migrate && pipenv run python source/manage.py collectstatic

HEALTHCHECK --interval=900s --timeout=10s --retries=3 CMD curl 127.0.0.1:80 2>/dev/null | grep 'You are a guest' || exit 1
ENTRYPOINT pipenv run python source/manage.py runserver --insecure 0.0.0.0:80