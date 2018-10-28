package gtf

import "testing"

func TestParseTransLine(t *testing.T) {
	line := []string{
		"chr1",
		"HAVANA",
		"transcript",
		"139790",
		"140339",
		".",
		"-",
		".",
		"gene_id \"ENSG00000239906.1\"; transcript_id \"ENST00000493797.1\"; gene_type \"antisense\"; gene_status \"KNOWN\"; gene_name \"RP11-34P13.14\"; transcript_type \"antisense\"; transcript_status \"KNOWN\"; transcript_name \"RP11-34P13.14-001\"; level 2; tag \"basic\"; transcript_support_level \"2\"; havana_gene \"OTTHUMG00000002481.1\"; havana_transcript \"OTTHUMT00000007038.1\", tag \"ncRNA_host\"; havana_gene \"OTTHUMG00000002480.3\";",
	}
	actual := parseTransLine(line)
	expected := Transcript{
		chrom:	"chr1",
		lower:	139790,
		upper:	140339,
		strand:	"-",
		geneID:	"ENSG00000239906.1",
		transcriptID:	"ENST00000493797.1",
		transcriptType:	"antisense",
	}
	if actual != expected {
		t.Errorf("Transcript was incorrect, got: %v, want: %v.", actual, expected)
	}
}
