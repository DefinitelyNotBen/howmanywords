# howmanywords
Quick program to practice string manipulations in Go

howMany() function will read a string given to it and return the number of words in a sentence

Limits are as follows 
 - Words are separated by white space
 - A words must only contain letters
    - 'wrfgewrgw' is a word
    - 'gvef$gs' is not a word
    - 'rge3' is not a word
 - Words separated by a '-' are considered 1 word
    - 'hello-world' is 1 word
 - If a word is followed by one of the following characters {',' '.' '?' '!'} then it is still counted and that is the end of the sentence
    - 'hello. world' only the first word would be counted
 
