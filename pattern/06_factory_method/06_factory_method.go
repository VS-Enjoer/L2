package main

import "fmt"

/*
Фабричный метод — это порождающий паттерн проектирования, который определяет общий интерфейс для создания объектов в
суперклассе, позволяя подклассам изменять тип создаваемых объектов.
*/

// Document - интерфейс документа
type Document interface {
	Open()
	Save()
}

// DocumentFactory - интерфейс фабрики документов
type DocumentFactory interface {
	CreateDocument() Document
}

// TextDocument - конкретная реализация текстового документа
type TextDocument struct{}

func (t *TextDocument) Open() {
	fmt.Println("Opening Text Document")
}

func (t *TextDocument) Save() {
	fmt.Println("Saving Text Document")
}

// TextDocumentFactory - фабрика для создания текстовых документов
type TextDocumentFactory struct{}

func (t *TextDocumentFactory) CreateDocument() Document {
	return &TextDocument{}
}

// SpreadsheetDocument - конкретная реализация таблицы
type SpreadsheetDocument struct{}

func (s *SpreadsheetDocument) Open() {
	fmt.Println("Opening Spreadsheet Document")
}

func (s *SpreadsheetDocument) Save() {
	fmt.Println("Saving Spreadsheet Document")
}

// SpreadsheetDocumentFactory - фабрика для создания таблиц
type SpreadsheetDocumentFactory struct{}

func (s *SpreadsheetDocumentFactory) CreateDocument() Document {
	return &SpreadsheetDocument{}
}

func main() {
	// Используем фабричные методы для создания документов
	textDocumentFactory := &TextDocumentFactory{}
	textDocument := textDocumentFactory.CreateDocument()
	textDocument.Open()
	textDocument.Save()

	spreadsheetDocumentFactory := &SpreadsheetDocumentFactory{}
	spreadsheetDocument := spreadsheetDocumentFactory.CreateDocument()
	spreadsheetDocument.Open()
	spreadsheetDocument.Save()
}
