from elasticapm.contrib.flask import ElasticAPM
from flask import Flask

app = Flask(__name__)

app.config['ELASTIC_APM'] = {
    'SERVICE_NAME': 'flask',
    'SECRET_TOKEN': 'o8U7XSZsJ0lA83ad0N4iO053',
    'SERVER_URL': 'https://192.168.2.225:8200',
    'VERIFY_SERVER_CERT': False,
    'ENVIRONMENT': 'production',
}

apm = ElasticAPM(app)


@app.route("/")
def index():
    return "Hello"


if __name__ == "__main__":
    app.run(debug=True)
