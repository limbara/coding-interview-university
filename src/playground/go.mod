module data/playground

go 1.24.6

replace data/array => ../array

require (
	data/array v0.0.0-00010101000000-000000000000
	data/linked_list v0.0.0-00010101000000-000000000000
)

replace data/linked_list => ../linked_list
