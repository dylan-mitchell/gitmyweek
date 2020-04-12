# gitmyweek

A simple command line tool to show a summary of your Git commits.

To use

`go build`

`./gitmyweek -dir="Path/To/Development/Dir" -author="Github Username" -start="YYYY-MM-DD"`

If no start date is given then, the default is the last week.

This tool only works if you have all of your Git repos in one directory. For example I have all of mine in `~/Development`.

So I would do `./gitmyweek -dir="/Users/dylanmitchell/Development" -author="dylan-mitchell" -start="2020-02-05"`

Output:

```
Monday
	memex
		Added ability to generate yearly summary
	ParseTakeout
		Add year to yearly summary
		Added functions to get total summary
Tuesday
	memex
		Added screenshots
		Updated package.json
		Added total summary and about sections
Wednesday
	ParseTakeout
		Added logic to parse youtube watch history
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
		Add year to yearly summary
		Added functions to get total summary
	memex
		Added ability to generate yearly summary
Sunday
	ParseTakeout
		Add year to yearly summary
		Added functions to get total summary
	memex
		Added ability to generate yearly summary
```
