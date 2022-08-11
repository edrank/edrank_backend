from flask import jsonify, Flask
import pymysql
from flaskext.mysql import MySQL


app = Flask(__name__)

mysql = MySQL()

# MySQL configurations
app.config['MYSQL_DATABASE_HOST'] = '185.210.145.52'
app.config['MYSQL_DATABASE_USER'] = 'u341279852_edrank_dev'
app.config['MYSQL_DATABASE_PASSWORD'] = '#Edrank@SIH22#'
app.config['MYSQL_DATABASE_DB'] = 'u341279852_edrank_db'

mysql.init_app(app)

@app.route('/')
def users():
    conn = mysql.connect()

    cursor = conn.cursor(pymysql.cursors.DictCursor)
    cursor.execute('SELECT * FROM teachers')

    rows = cursor.fetchall()

    resp = jsonify(rows)
    resp.status_code = 200

    return resp

if __name__ == "__main__":
    app.run(host="0.0.0.0", debug=True, port=5002)