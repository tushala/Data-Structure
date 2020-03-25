package main

type Heap struct {
	s []int
}
func (self *Heap) swap(i,j int){
	self.s[i],self.s[j] = self.s[j], self.s[i]
}
func New() *Heap {
	s := make([]int, 0)
	return &Heap{s:s}
}
func (self *Heap)Put(x int){
	if len(self.s) == 0{
		self.s = append(self.s, x)
		return
	}
	self.put(x, len(self.s))
}
func (self *Heap) put(x, length int){
	
}
