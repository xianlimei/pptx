/*
Package pptx creates and adds slides to powerpoint presentations.

The package can be used to create simple powerpoint pptx files or
to append slides to an existing presentation.

The file format is a compressed zip archive with xml files and images.
The package unpacks the archive in memory, adds files and manipulates
existing xml files to add new slides.
It uses github.com/beevik/etree for xml operations.

The aim of the package is not to support all of powerpoint's possibilities
but to understand just enough to be able to add new slides to an existing
presentation.

Supported are:

Textboxes
	Position
	Font, Fontsize
	Multiline text
	mark as title

Images
	Position

*/
package pptx
