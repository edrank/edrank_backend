import matplotlib.pyplot as plt 
import seaborn as sns

colors = sns.color_palette('pastel')


def create_piechart(ticks,vals, q):
	plt.figure(figsize=(4,4), tight_layout = True)
	explode_list = [0.025,0.025, 0.025]

	plt.pie(vals, labels=ticks, autopct = "%.0f %%", explode = explode_list, colors = colors)
	plt.title(f"Feedback counts - Quarter %i" %q, weight ="bold")
	
	if(q == 1):
		plt.savefig("piechart_1.png")
	elif(q == 2):
		plt.savefig("piechart_2.png")
	plt.show()


def display_feedbacks():
	print("Top 3 Positive Feedbacks : ")
	pos = list(obj_dict[k3].values())[0]
		
	for num, i in enumerate(pos, start = 1):
		print(num, i)

	print("\nFew Critical Feedbacks : ")
	neg = list(obj_dict[k3].values())[1]
	for num, i in enumerate(neg, start = 1):
		print(num, i)


if __name__ == '__main__':
	
	obj_dict = {"teacher" : {	#basic details
								"t_id" : "T001",
								"name" : "Mark Johnson", 
								"designation" : "Assistant Professor",
								"course" : "Computer Applications",
								"email" : "markjohnson01@gmail.com",
								"contact" : "9876543456" 
							},

				"college" : {	#college details
								"name" : "ABC college of Technology",
								"address" : "Salem, Tamil Nadu"
							},

				"assessment_period" : { #semester_wise
										"from" : "July, 2019",
										"to" : "December, 2019"
										},

				"quarter_1" : 12,
						

				"quarter_2" : 7,

				"q1_feedback" : { #feedback_counts
							"total_feedbacks" : 175,
							"positive" : 95,
							"negative" : 35,
							"neutral" : 45
							},

				"q2_feedback" : { #feedback_counts
							"total_feedbacks" : 180,
							"positive" : 120,
							"negative" : 40,
							"neutral" : 20
							},

				"top_feedbacks" : { #text_feedbacks
									"positive" : ["Manner of teaching is so wonderful and refreshing!! She’s patient and supportive, but really knows how to motivate her students.",
												"You are amazing at what you do! Your passion and dedication is beyond words!",
												"She is a calm and thoughtful tutor who really cares about her students"],

									"negative" : ["Syllabus is not completed on time.",
												"Teacher is very strict!!",
												"She doesn't cover the entire Syllabus."]
								}
				
				}

	

	k1 = "q1_feedback"

	if k1 in obj_dict:
		ticks = list(obj_dict[k1].keys())
		vals = list(obj_dict[k1].values())

		create_piechart(ticks[1:],vals[1:], 1)

	k2 = "q2_feedback"

	if k2 in obj_dict:
		ticks = list(obj_dict[k2].keys())
		vals = list(obj_dict[k2].values())

		create_piechart(ticks[1:],vals[1:], 2)

	k3 = "top_feedbacks"

	if k3 in obj_dict:
		display_feedbacks()

