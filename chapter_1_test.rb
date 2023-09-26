require "minitest/autorun"
require_relative "chapter_1"

class TestChapter1 < Minitest::Test
  def test_compress_single_char_string
    assert_equal "a", compress("a")
  end

  def test_compress_two_char_string
    assert_equal "aa", compress("aa")
  end

  def test_compress_three_char_string
    assert_equal "a3", compress("a3")
  end

  def test_compress_four_a_one_b_string
    assert_equal "a4b1", compress("aaaab")
  end
end
