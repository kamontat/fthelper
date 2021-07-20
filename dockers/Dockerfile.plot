FROM freqtradeorg/freqtrade:develop_plot

# Install postgres compiler support
USER root
RUN apt-get update \
  && apt-get -y install libpq-dev gcc
USER ftuser

# Add Postgres support
RUN pip install --user psycopg2

# Cleanup
USER root
RUN apt-get clean \
  && apt-get autoclean -y \
  && apt-get autoremove -y
USER ftuser
