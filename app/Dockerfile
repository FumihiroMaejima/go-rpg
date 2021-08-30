FROM php:8.0.3-fpm-alpine

# tinker(psysh)
ARG PSYSH_DIR=/usr/local/share/psysh
ARG PSYSH_PHP_MANUAL=$PSYSH_DIR/php_manual.sqlite
ARG PHP_MANUAL_URL=http://psysh.org/manual/ja/php_manual.sqlite

# composer
ENV COMPOSER_ALLOW_SUPERUSER 1
ENV COMPOSER_HOME /composer

# phpcs(for php8)
ENV PHP_CS_FIXER_IGNORE_ENV 1

RUN apk update && \
    apk add --update --no-cache --virtual=.build-dependencies \
    autoconf \
    gcc \
    g++ \
    make \
    tzdata \
    git && \
    apk add --update --no-cache \
    icu-dev \
    libzip-dev \
    oniguruma-dev && \
    pecl install xdebug && \
    git clone https://github.com/phpredis/phpredis.git /usr/src/php/ext/redis && \
    apk del .build-dependencies && \
    docker-php-ext-install intl pdo_mysql mbstring zip bcmath redis && \
    docker-php-ext-enable xdebug && \
    mkdir $PSYSH_DIR && curl $PHP_MANUAL_URL -o $PSYSH_PHP_MANUAL && \
    curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/bin --filename=composer && \
    composer config -g process-timeout 3600 && \
    composer config -g repos.packagist composer https://packagist.jp


WORKDIR /var/www/html
