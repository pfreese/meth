package gtf

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Transcript struct {
	chrom    string
	lower	int
	upper	int
	strand	string
	geneID	string
	transcriptID	string
	transcriptType	string
}

type Gene struct {
	chrom    string
	lower	int
	upper	int
	strand	string
	geneID	string
	geneType	string
	transcripts	[]string
}

func (gene *Gene) AddTrans(transcript Transcript) []string {
	gene.transcripts = append(gene.transcripts, transcript.transcriptID)
	return gene.transcripts
}

// ParseGTF parses the
// 60725 genes
// 199,348 transcripts
func ParseGTF(gtfPath string) error {
	csvFile, _ := os.Open(gtfPath)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.LazyQuotes = true
	reader.Comma = '\t'
	nGenes := 0
	nTrans := 0
	genes := make(map[string]Gene)
	for lineNum := 1; ; lineNum++ {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		if len(line) != 9 {
			return fmt.Errorf("line %v does not have 9 columns", lineNum)
		}
		fmt.Println(line)
		fmt.Println(fmt.Sprintf("%T", line))
		//fmt.Println(len(line))
		fmt.Print("\n\n\n")
		if lineNum > 100 {
			break
		}
		if line[2] == "gene" {
			nGenes += 1
			gene := parseGeneLine(line)
			genes[gene.geneID] = gene
			fmt.Println("geenezzz")
		} else if line[2] == "transcript" {
			nTrans += 1
			trans := parseTransLine(line)
			fmt.Println(trans)
			// Add this transcript to its gene
			transGene := genes[trans.geneID]
			fmt.Println("BEFOREEEEEE")
			fmt.Println(transGene)
			transGene.AddTrans(trans)
			//genes[trans.geneID].AddTrans(trans)
			fmt.Println(transGene)
			genes[trans.geneID] = transGene
		}
	}
	fmt.Println(nGenes, "genes")
	fmt.Println(nTrans, "transcripts")
	fmt.Println(genes)
	return nil
}

// parseAnnotField parses the annotation (9th) column of a .gtf file line
// and returns a map like:
// map[gene_status:KNOWN
// 		transcript_status:KNOWN
// 		havana_gene:OTTHUMG00000002481.1
// 		transcript_type:antisense
// 		transcript_name:RP11-34P13.14-001
// 		level:2
// 		tag:basic
// 		gene_id:ENSG00000239906.1
// 		transcript_id:ENST00000493797.1
// 		gene_type:antisense
// 		gene_name:RP11-34P13.14
// 		transcript_support_level:2
// 		havana_transcript:OTTHUMT00000007038.1]
func parseAnnotField(annot string) map[string]string {
	annots := make(map[string]string)
	s := strings.Split(annot, ";")
	for _, field := range s {
		// Strip leading/trailing spaces, after which the key/value annotation
		// is separated by a space
		trimmed := strings.Trim(field, " ")
		keyVal := strings.Split(trimmed, " ")
		if len(keyVal) == 2 {
			annots[keyVal[0]] = strings.Trim(keyVal[1], "\"")
		}
	}
	return annots
}



//
// [chr1 HAVANA transcript 139790 140339 . - . gene_id "ENSG00000239906.1"; transcript_id "ENST00000493797.1"; gene_type "antisense"; gene_status "KNOWN"; gene_name "RP11-34P13.14"; transcript_type "antisense"; transcript_status "KNOWN"; transcript_name "RP11-34P13.14-001"; level 2; tag "basic"; transcript_support_level "2"; havana_gene "OTTHUMG00000002481.1"; havana_transcript "OTTHUMT00000007038.1";]
func parseTransLine(line []string) Transcript {
	lower, _ := strconv.Atoi(line[3])
	upper, _ := strconv.Atoi(line[4])
	annots := parseAnnotField(line[8])
	trans := Transcript{
		chrom:        line[0],
		lower:        lower,
		upper:        upper,
		strand:       line[6],
		geneID:       annots["gene_id"],
		transcriptID: annots["transcript_id"],
		transcriptType: annots["transcript_type"],
	}
	return trans
}



// parseGeneLine parses a line corresponding to a gene:
// [chr1 HAVANA gene 139790 140339 . - . gene_id "ENSG00000239906.1"; gene_type "antisense"; gene_status "KNOWN"; gene_name "RP11-34P13.14"; level 2; havana_gene "OTTHUMG00000002481.1";]
func parseGeneLine(line []string) Gene {
	lower, _ := strconv.Atoi(line[3])
	upper, _ := strconv.Atoi(line[4])
	annots := parseAnnotField(line[8])
	gene := Gene{
		chrom:        line[0],
		lower:        lower,
		upper:        upper,
		strand:       line[6],
		geneID:       annots["gene_id"],
		transcripts: []string{},
	}
	return gene
}

