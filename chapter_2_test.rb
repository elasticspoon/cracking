require "minitest/autorun"
require_relative "chapter_2"

class TestChapter2 < Minitest::Test
  def create_list(arr)
    head = nil
    arr.reverse_each do |val|
      head = Node.new(val, head)
    end
    head
  end

  def test_palindrome_even_len
    list = create_list([1, 2, 3, 3, 2, 1])
    assert palindrome?(list)
  end

  def test_palindrome_odd_len
    list = create_list([1, 2, 3, 2, 1])
    assert palindrome?(list)
  end

  def test_not_palindrome_odd_len
    list = create_list([1, 2, 3, 4, 5])
    assert !palindrome?(list)
  end
end
