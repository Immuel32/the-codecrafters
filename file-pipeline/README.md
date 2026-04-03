The project is a text transformation program. It reads a text file, applies a set of transformation rules, writes the processed lines to an output file, and prints summary to the terminal. But the transformation is carried out based on my group's selected 5 rule.

And the rules are as follows:
- Trim all leading and trailing whitespace
- Replace TODO: with ✦ ACTION
- Convert ALL CAPS lines to Title Case
- Convert all lowercase lines to uppercase
- Reverse the words in any line that contains the word REVERSE

The output format:
- Each processed line is prefixed with a line number starting from 1
- The output file begins with a header
- There is a summary output to the terminal at the end

How to run the program
- Make sure all the files are in the folder as the main.go
- Create an input text file call test_input.txt where you will save the text for testing your program
- Use go run main.go test_input.txt test_output.txt

Example of the input test and the output test(processed)

***Input***
WELCOME TO THE GOPHER FIELD REPORT
 TODO: Collect all soil samples 
 TODO: Record water readings 
 THIS LINE HAS EXTRA SPACES 
 NORMAL operations should continue 
 Check this line 
 END OF REPORT 
 MIXED Case Line With TODO: action reverse needed

 ***Output***
 Gopher's Sentinel Field Report - Processed
1. welcome to the gopher field report
2. action: collect all soil samples
3. action: record water readings
4. this line has extra spaces
5. normal operations should continue
6. check this line
7. end of report
8. dexim esac enil htiw :noitca noitca esrever dedeen

-----summary-----
Lines read    : 8
Lines written : 8
Lines removed : 0
Rules applied : TrimWhitespace, ReplaceTODO, AllCapsToTitle, Lower, reverseWord


