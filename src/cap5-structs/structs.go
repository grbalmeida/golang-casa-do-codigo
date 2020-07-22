package main

import "fmt"

type Arquivo struct {
	nome       string
	tamanho    float64
	caracteres int
	palavras   int
	linhas     int
}

func (arq *Arquivo) TamanhoMedioDePalavra() float64 {
	return float64(arq.caracteres) / float64(arq.palavras)
}

func (arq *Arquivo) MediaDePalavrasPorLinha() float64 {
	return float64(arq.palavras) / float64(arq.linhas)
}

func main() {
	arquivo := Arquivo{"artigo.txt", 12.68, 12986, 1862, 220}
	fmt.Println(arquivo)

	codigoFonte := Arquivo{tamanho: 1.12, nome: "programa.go"}
	fmt.Println(codigoFonte)

	fmt.Printf("%s\t%.2fKB\n", arquivo.nome, arquivo.tamanho)
	fmt.Printf("%s\t%.2fKB\n",
		codigoFonte.nome,
		codigoFonte.tamanho)

	ponteiroArquivo := &Arquivo{tamanho: 7.29, nome: "arquivo.txt"}

	fmt.Printf("%s\t%.2fKB\n",
		ponteiroArquivo.nome,
		ponteiroArquivo.tamanho)

	// Mutabilidade

	codigoFonte.tamanho = 23.42

	fmt.Printf("%s\t%.2fKB\n",
		codigoFonte.nome,
		codigoFonte.tamanho)

	fmt.Printf("Média de palavras por linha: %.2f\n",
		arquivo.MediaDePalavrasPorLinha())

	fmt.Printf("Tamanho médio de palavra: %.2f",
		arquivo.TamanhoMedioDePalavra())
}
