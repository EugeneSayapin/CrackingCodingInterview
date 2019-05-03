package main

type animal interface {
	makeSound() string
	getType() string
}

type cat struct {
}

type dog struct {
}

func (_ dog) makeSound() string {
	return "bough"
}
func (_ dog) getType() string {
	return "dog"
}

func (_ cat) makeSound() string {
	return "mew"
}

func (_ cat) getType() string {
	return "cat"
}

type node struct {
	animal animal
	number int
	next   *node
}

type animalQueue struct {
	head *node
	tail *node
}

func (this *animalQueue) enqueue(animal animal, counter int) {
	node := &node{animal: animal, number: counter, next: nil}
	if this.tail != nil {
		this.tail.next = node
	}
	if this.head == nil {
		this.head = node
	}
	this.tail = node
}

func (this *animalQueue) dequeue() animal {
	result := this.head.animal
	this.head = this.head.next
	return result
}

type shelter struct {
	dogQueue *animalQueue
	catQueue *animalQueue
	counter  int
}

func (this *shelter) enqueue(animal animal) {
	this.counter++
	switch animal.getType() {
	case "cat":
		this.catQueue.enqueue(animal, this.counter)
	case "dog":
		this.dogQueue.enqueue(animal, this.counter)
	}
}

func (this *shelter) dequeueAny() animal {
	if this.catQueue.head.number < this.dogQueue.head.number {
		return this.catQueue.dequeue()
	} else {
		return this.dogQueue.dequeue()
	}
}

func (this *shelter) dequeueDog() animal {
	return this.dogQueue.dequeue()
}

func (this *shelter) dequeuCat() animal {
	return this.catQueue.dequeue()
}

func main() {
	shelter := &shelter{catQueue: &animalQueue{}, dogQueue: &animalQueue{}}
	shelter.enqueue(&cat{})
	shelter.enqueue(&cat{})
	shelter.enqueue(&dog{})
	shelter.enqueue(&dog{})
	shelter.dequeueDog()
	shelter.dequeueAny()
	shelter.dequeueAny()
}
