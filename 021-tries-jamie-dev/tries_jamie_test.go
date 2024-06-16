package main

import (
	"fmt"
	"testing"
)

func BenchmarkTries(b *testing.B) {
	myTrie := InitTrie()

	toAdd := []string{
		"antidisestablishmentarianism",
		"floccinaucinihilipilification",
		"pseudopseudohypoparathyroidism",
		"supercalifragilisticexpialidocious",
		"hippopotomonstrosesquipedaliophobia",
		"pneumonoultramicroscopicsilicovolcanoconiosis",
		"thyroparathyroidectomized",
		"psychoneuroendocrinological",
		"hepaticocholangiogastrostomy",
		"spectrophotofluorometrically",
		"incomprehensibilities",
		"uncharacteristically",
		"immunoelectrophoretically",
		"otorhinolaryngological",
		"electroencephalographically",
		"radioimmunoelectrophoresis",
		"thermophosphorescence",
		"trinitrophenylmethylnitramine",
		"thyroparathyroidectomized",
		"prostatovesiculectomy",
		"gastroenterocolitis",
		"tetrahydrocannabinol",
		"deoxyribonucleoprotein",
		"phosphatidylethanolamine",
		"cholecystogastrostomy",
		"psychophysicotherapeutics",
		"laryngotracheobronchitis",
		"pneumoencephalographically",
		"anatomicophysiologicochemical",
		"intercontinentalballistic",
	}

	for _, v := range toAdd {
		myTrie.Insert(v)
	}

	result := []string{}

	for i:=0 ; i < b.N ; i++ {
		result = myTrie.DFS()
	}

	fmt.Println(result)
}
