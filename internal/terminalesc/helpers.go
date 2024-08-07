package terminalesc

// StringToCodeSequence turns an arbitrary string into a CodeSequence.
func StringToCodeSequence(s string) CodeSequence {
	seq := make(CodeSequence, len(s))

	for i, v := range s {
		seq[i] = Code(v)
	}

	return seq
}
