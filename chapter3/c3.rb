class StackOfStacks
  private attr_reader :stacks, :max_size
  def initialize(max_size)
    @stacks = []
    @max_size = max_size
  end

  def push(value)
    stacks.push([]) if stacks.empty? || stacks.last.size == max_size
    stacks.last.push(value)
  end

  def pop
    return nil if stacks.empty?
    target = stacks.last
    value = target.pop
    stacks.pop if target.empty?

    value
  end

  def pop_at(index)
    return nil if stacks.empty? || index >= stacks.size || index < 0
    value = stacks[index].pop
    stacks.delete_at(index) if stacks[index].empty?
    value
  end
end

def sort_stack(stl)
  lts = []
  return nil if stl.nil? || stl.empty?

  temp = stl.pop
  Kernel.loop do
    until stl.empty?
      if stl.last < temp
        lts.push(stl.pop)
      else
        lts.push(temp)
        temp = stl.pop
      end
    end
    num_swaps = 0
    until lts.empty?
      if lts.last > temp
        stl.push(lts.pop)
        num_swaps += 1
      else
        stl.push(temp)
        temp = lts.pop
      end
    end
    break if num_swaps <= 1
  end

  stl.push(temp)
  stl
end

class Animal
  attr_accessor :name, :next_animal, :prev_animal

  def initialize(name, next_animal: nil, prev_animal: nil)
    @name = name
    @next_animal = next_animal
    @prev_animal = prev_animal
  end
end

class Cat < Animal
  attr_accessor :next_cat, :prev_cat

  def initialize(name, next_animal: nil, prev_animal: nil, prev_cat: nil, next_cat: nil)
    super(name, next_animal:, prev_animal:)
    @next_cat = next_cat
    @prev_cat = prev_cat
  end
end

class Dog < Animal
  attr_accessor :next_dog, :prev_dog

  def initialize(name, next_animal: nil, prev_animal: nil, prev_dog: nil, next_dog: nil)
    super(name, next_animal:, prev_animal:)
    @next_dog = next_dog
    @prev_dog = prev_dog
  end
end

class AnimalShelter
  private attr_accessor :cats, :dogs, :animals, :oldest_cat, :oldest_dog, :oldest_animal

  def initialize
    @cats = nil
    @dogs = nil
    @animals = nil
    @oldest_cat = nil
    @oldest_dog = nil
    @oldest_animal = nil
  end

  def enqueue(animal, type)
    case type
    when :cat
      enqueue_cat(animal)
    when :dog
      enqueue_dog(animal)
    else
      raise ArgumentError, "Invalid type"
    end
  end

  def dequeue_any
    return nil if animals.nil?

    animal = oldest_animal

    if animal.is_a?(Cat)
      dequeue_cat
    elsif animal.is_a?(Dog)
      dequeue_dog
    end
    animal.name
  end

  def dequeue_cat
    return nil if cats.nil?
    cat = oldest_cat
    self.oldest_cat = cat.prev_cat
    oldest_cat&.next_cat = cat.next_cat
    cat.prev_animal&.next_animal = cat.next_animal
    cat.next_animal&.prev_animal = cat.prev_animal
    if cat == oldest_animal
      self.oldest_animal = cat.prev_animal
      oldest_animal&.next_animal = cat.next_animal
    end

    cat.name
  end

  def dequeue_dog
    return nil if dogs.nil?
    dog = oldest_dog
    self.oldest_dog = dog.prev_dog
    oldest_dog&.next_dog = dog.next_dog
    dog.prev_animal&.next_animal = dog.next_animal
    dog.next_animal&.prev_animal = dog.prev_animal
    if dog == oldest_animal
      self.oldest_animal = dog.prev_animal
      oldest_animal&.next_animal = dog.next_animal
    end
    dog.name
  end

  def print_shelter
    a = "animals: "
    animal = animals
    until animal.nil?
      a += "#{animal.name} -> "
      animal = animal.next_animal
    end
    puts a

    d = "dogs: "
    dog = dogs
    until dog.nil?
      d += "#{dog.name} -> "
      dog = dog.next_dog
    end
    puts d

    c = "cats: "
    cat = cats
    until cat.nil?
      c += "#{cat.name} -> "
      cat = cat.next_cat
    end
    puts c
    puts "oldest cat: #{oldest_cat&.name}"
    puts "oldest dog: #{oldest_dog&.name}"
    puts "oldest animal: #{oldest_animal&.name}"
  end

  private

  def enqueue_cat(cat)
    cat = Cat.new(cat, next_animal: animals, next_cat: cats)
    if cats.nil?
      self.oldest_cat = cat
    else
      cats.prev_cat = cat
    end
    self.cats = cat

    enqueue_animal(cat)
  end

  def enqueue_dog(dog)
    dog = Dog.new(dog, next_animal: animals, next_dog: dogs)
    if dogs.nil?
      self.oldest_dog = dog
    else
      dogs.prev_dog = dog
    end
    self.dogs = dog
    enqueue_animal(dog)
  end

  def enqueue_animal(animal)
    if animals.nil?
      self.oldest_animal = animal
    else
      animals.prev_animal = animal
    end
    self.animals = animal
  end
end
