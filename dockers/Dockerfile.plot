FROM freqtradeorg/freqtrade:develop_plot

# Add Postgres support
RUN pip install --user psycopg2
