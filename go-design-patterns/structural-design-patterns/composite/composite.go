package composite

type Component interface {
	Add(child Component)
	Name() string
	Child() []Component
	Print(prefix string) string
}

type Directory struct {
	name   string
	childs []Component
}

func (d *Directory) Add(child Component) {
	d.childs = append(d.childs, child)
}

func (d *Directory) Name() string {
	return d.name
}

func (d *Directory) Child() []Component {
	return d.childs
}

func (d *Directory) Print(prefix string) string {
	result := prefix + "/" + d.Name() + "\n"
	for _, val := range d.Child() {
		result += val.Print(prefix + "/" + d.Name())
	}
	return result
}

type File struct {
	name string
}

func (f *File) Add(_ Component) {

}

func (f *File) Name() string {
	return f.name
}

func (f *File) Child() []Component {
	return []Component{}
}

func (f *File) Print(prefix string) string {
	return prefix + "/" + f.Name() + "\n"
}

func NewDirectory(name string) *Directory {
	return &Directory{name: name}
}

func NewFile(name string) *File {
	return &File{name: name}
}
