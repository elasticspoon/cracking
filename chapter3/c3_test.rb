require "minitest/autorun"
require_relative "c3"

class TestChapter3 < Minitest::Test
  def test_stacks
    ss = StackOfStacks.new(3)
    (1..10).each do |i|
      ss.push(i)
    end

    6.downto(4).each do |i|
      assert ss.pop_at(1) == i
    end

    10.downto(7).each do |i|
      assert ss.pop == i
    end

    3.downto(1).each do |i|
      assert ss.pop == i
    end
  end

  def test_sort_stack_empty
    assert sort_stack(nil).nil?
    assert sort_stack([]).nil?
  end

  def test_sort_stack_sorted
    assert_equal sort_stack([3, 2, 1]), [1, 2, 3].reverse
  end

  def test_sort_stack_unsorted_short
    assert_equal sort_stack([3, 2, 1]), [1, 2, 3].reverse
  end

  def test_sort_stack_unsorted_long
    assert_equal sort_stack([5, 3, 8, 7, 11, 10, 2, 1, 4, 6, 9]), [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11].reverse
  end

  def test_animal_shelter
    as = AnimalShelter.new
    as.enqueue(1, :cat)
    as.enqueue(2, :cat)
    as.enqueue(3, :dog)
    as.enqueue(4, :cat)
    as.enqueue(5, :dog)
    as.enqueue(6, :dog)
    as.enqueue(7, :cat)
    as.enqueue(8, :cat)
    as.enqueue(9, :dog)
    as.enqueue(10, :dog)

    assert_equal as.dequeue_any, 1
    assert_equal as.dequeue_dog, 3
    assert_equal as.dequeue_any, 2
    assert_equal as.dequeue_dog, 5
    assert_equal as.dequeue_cat, 4
    assert_equal as.dequeue_cat, 7
    assert_equal as.dequeue_cat, 8
    assert_equal as.dequeue_dog, 6
    assert_equal as.dequeue_dog, 9
    assert_equal as.dequeue_any, 10
  end
end
