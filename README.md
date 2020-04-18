# gitmyweek

A simple command line tool to show a summary of your Git commits. This can be useful for standup, or Monday morning remembering what it was you did last week.

To use

`go build`

`./gitmyweek -dir="Path/To/Development/Dir" -author="Github Username" -start="YYYY-MM-DD"`

If no start date is given then, the default is the last week.

This tool only works if you have all of your Git repos in one directory. For example I have all of mine in `~/Development`.

So I would do `./gitmyweek -dir="/Users/dylanmitchell/Development" -author="dylan-mitchell" -start="2020-02-06"`

Output:

```
Thursday
	ParseTakeout
		Added model files
		Reorganized project to be the same package
		Added more funcs to access db
		Added models to put items in sqlite db
Friday
	ParseTakeout
		Added function to get years
Saturday
	ParseTakeout
		Dont accept items that are over 250 chars
		Unencode results
		Leave results url encoded
		Added function to search items
Sunday
	ParseTakeout
		Ignore empty channels
		Added functions to get yearly summaries
	memex
		init commit
Monday
	ParseTakeout
		Add year to yearly summary
		Added functions to get total summary
	memex
		Added ability to generate yearly summary
Tuesday
	memex
		Added screenshots
		Updated package.json
		Added total summary and about sections
```
