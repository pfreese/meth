package brainmap

import (
	"fmt"
	"os"
	"path"
)

// All downloaded file basenames begin with this fStart
const fStart = "human_VISp_2018-10-04"

// hVISpConfig holds paths to the 4 downloaded files:
// 		human_VISp_2018-10-04_genes-rows.csv
// 		human_VISp_2018-10-04_exon-matrix.csv
//		human_VISp_2018-10-04_intron-matrix.csv
// 		human_VISp_2018-10-04_samples-columns.csv
type hVISpConfig struct {
	genes    string
	exon	string
	intron	string
	samples	string
}

// ProcesshVISpConfig stores the file paths of the 4 downloaded data files.
// Files are within the dataDIR specified in the configuration.
func ProcesshVISpConfig(config Configuration) (hVISpConfig, error) {
	hConfig := hVISpConfig{}

	// The genes file
	genesF := path.Join(config.DataDIR, fmt.Sprintf("%s_genes-rows.csv", fStart))
	if _, err := os.Stat(genesF); os.IsNotExist(err) {
		fmt.Println(err)
		return hVISpConfig{}, err
	}
	hConfig.genes = genesF

	// The exons file
	exonsF := path.Join(config.DataDIR, fmt.Sprintf("%s_exon-matrix.csv", fStart))
	if _, err := os.Stat(exonsF); os.IsNotExist(err) {
		fmt.Println(err)
		return hVISpConfig{}, err
	}
	hConfig.exon = exonsF

	// The introns file
	intronsF := path.Join(config.DataDIR, fmt.Sprintf("%s_intron-matrix.csv", fStart))
	if _, err := os.Stat(intronsF); os.IsNotExist(err) {
		fmt.Println(err)
		return hVISpConfig{}, err
	}
	hConfig.intron = intronsF

	// The sample columns file
	sampColF := path.Join(config.DataDIR, fmt.Sprintf("%s_samples-columns.csv", fStart))
	if _, err := os.Stat(sampColF); os.IsNotExist(err) {
		fmt.Println(err)
		return hVISpConfig{}, err
	}
	hConfig.samples = sampColF
	
	return hConfig, nil
}
