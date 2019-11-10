package full_justify

import "testing"

func TestFullJustify(t *testing.T) {
	var words, hope, ret []string
	var maxWidth int

	words = []string{"This", "is", "an", "example", "of", "text", "justification."}
	maxWidth = 16
	hope = []string{"This    is    an", "example  of text", "justification.  "}
	ret = fullJustify(words, maxWidth)
	t.Logf("\nwords=%v maxWidth=%d \n hope=%v \n  ret=%v", words, maxWidth, hope, ret)

	words = []string{"What", "must", "be", "acknowledgment", "shall", "be"}
	maxWidth = 16
	hope = []string{"What   must   be", "acknowledgment  ", "shall be        "}
	ret = fullJustify(words, maxWidth)
	t.Logf("\nwords=%v maxWidth=%d \nhope=%v \n ret=%v", words, maxWidth, hope, ret)

	words = []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain",
		"to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}
	maxWidth = 20
	hope = []string{"Science  is  what we", "understand      well", "enough to explain to", "a  computer.  Art is", "everything  else  we", "do                  "}
	ret = fullJustify(words, maxWidth)
	t.Logf("\nwords=%v maxWidth=%d \nhope=%v \n ret=%v", words, maxWidth, hope, ret)

	words = []string{"ask", "not", "what", "your", "country", "can", "do", "for", "you", "ask", "what", "you", "can", "do", "for", "your", "country"}
	maxWidth = 16
	hope = []string{"ask   not   what", "your country can", "do  for  you ask", "what  you can do", "for your country"}
	ret = fullJustify(words, maxWidth)
	t.Logf("\nwords=%v maxWidth=%d \nhope=%v \n ret=%v", words, maxWidth, hope, ret)
}
