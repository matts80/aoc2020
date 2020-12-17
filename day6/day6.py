if __name__ == '__main__':
	
	fname = 'input.txt'

	answers = ""
	answer_sets = []

	total_counts = 0
	total_unique = 0

	with open(fname) as fh:
		for line in fh:
			line = line.strip()

			if line != "":
				answers += line
				answer_sets.append(set(list(line)))
			else:
				total_counts += len(set(list(answers)))
				answers = ""

				# calculate the intersection of this group's answers
				total_unique += len(answer_sets[0].intersection(*answer_sets))

				# reset answer sets
				answer_sets = []

	print("Total counts: {}".format(total_counts))
	print("Total unique answers: {}".format(total_unique))
