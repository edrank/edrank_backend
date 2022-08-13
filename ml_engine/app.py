from flask import jsonify,request, Flask
import pymysql
from flaskext.mysql import MySQL
# from vaderSentiment import sentiment_scores


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

@app.route('/')
@app.route('/analyze', methods =['POST'])
def analyzeSentiment():
    if request.method == 'POST' and 'text_to_analyze' in request.form:
        text_to_analyze = request.form["text_to_analyze"]
        # sentiment_dict = sentiment_scores(text_to_analyze)
        return jsonify(
            score=91.33
        )
    else:
        return jsonify(
            error="bad request"
        )

if __name__ == "__main__":
    app.run(host="0.0.0.0", debug=True, port=5002)