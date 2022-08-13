from vaderSentiment.vaderSentiment import SentimentIntensityAnalyzer

def sentiment_scores(sentence):
    sid_obj = SentimentIntensityAnalyzer()
    sentiment_dict = sid_obj.polarity_scores(sentence)
	
    print(sentiment_dict)
	
    print(sentiment_dict['neg']*100, "% Negative")
	
    print(sentiment_dict['neu']*100, "% Neutral")
	
    print(sentiment_dict['pos']*100, "% Positive")

    return sentiment_dict