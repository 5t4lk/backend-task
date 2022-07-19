△ Potential corrections:

1. The functions latestArticles & retrieveSingleArticleByID are very similar.
The only difference is that they generate a link to read the data from the
website using http.Get. I wanted to combine these into one function somehow,
but haven't figured out how.
2. The process of function userChoiceAndLogic is a little bit strange on my
opinion. I'm sure there is some way to optimise the logic of this function.
3. Names of structs. Can be better.

△ Advantages in my opinion:

1. Func check(). Helps make error handling easier and reduces the amount of code.
2. Func repeat(). For ease of use, the software asks the user if they want to
repeat the process
3. Func enter(). Makes it easier to read the input.
4. Incorrect ID handling.

△ What Works / What does not:

●✓	Be written in Golang - https://golang.org/doc/
●✕	Use MongoDB to store news articles - (https://www.mongodb.com/download-center/community)
●✕	At regular intervals poll the endpoint for new news articles
●✓	Transform the XML feeds of news articles into appropriate model(s) and save them in the database
●✓	Provide two REST endpoints which return JSON:
	a.	Retrieve a list of all articles
	b.	Retrieve a single article by ID
●✓	Comments where appropriate
●✓	Don’t spend more than 1 day on the task - If it's not fully complete (but should compile!) That is fine.
●✓	Send over README file listing what works and what doesn't - as well as instructions (if any) on how to run

△ Why didn't do?

I had no experience with MongoDB and no experience with sending queries with
update checks. I figured out how to run the local MongoDB database on my
computer, but didn't understand how to add data from the software to the
database. Maybe I would have understood, but unfortunately time is running out.
