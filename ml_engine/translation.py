#pip install googletrans==3.1.0a0

from googletrans import Translator

translator = Translator()
text = 'Hey whatsup?'
result = translator.translate(text, src='en', dest='hi')

print(result.text)

