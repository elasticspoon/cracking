require "minitest/autorun"
require_relative "chapter_1"

class TestChapter1 < Minitest::Test
  # test compress method
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

  # test set_0_matrix method
  def test_set_0_matrix_2_2_no_0s
    expected = [[1, 1], [1, 1]]
    arrayy = [[1, 1], [1, 1]]
    actual = set_0_matrix(array)

    assert_equal expected, actual
  end

  def test_set_0_matrix_2_2_1_0
    expected = [[0, 0], [0, 1]]
    arrayy = [[0, 1], [1, 1]]
    actual = set_0_matrix(array)

    assert_equal expected, actual
  end

  def test_set_0_matrix_2_2_2_0
    expected = [[0, 0], [0, 0]]
    arrayy = [[0, 1], [0, 1]]
    actual = set_0_matrix(array)

    assert_equal expected, actual
  end

  def test_set_0_matrix_5_5_2_0
    expected =
      [[0, 0, 0, 0, 0],
        [0, 3, 3, 0, 3],
        [0, 0, 0, 0, 0],
        [0, 3, 3, 0, 3],
        [0, 3, 3, 0, 3]]
    arrayy =
      [[0, 3, 3, 3, 3],
        [3, 3, 3, 3, 3],
        [3, 3, 3, 0, 3],
        [3, 3, 3, 3, 3],
        [3, 3, 3, 3, 3]]
    actual = set_0_matrix(array)

    assert_equal expected, actual
  end

  def test_set_0_matrix_5_5_1_0
    expected =
      [[0, 0, 0, 0, 0],
        [0, 3, 3, 3, 3],
        [0, 3, 3, 3, 3],
        [0, 3, 3, 3, 3],
        [0, 3, 3, 3, 3]]
    arrayy =
      [[0, 3, 3, 3, 3],
        [3, 3, 3, 3, 3],
        [3, 3, 3, 3, 3],
        [3, 3, 3, 3, 3],
        [3, 3, 3, 3, 3]]
    actual = set_0_matrix(array)

    assert_equal expected, actual
  end

  # test is_rotation method
  def test_is_rotation_1_char_true
    assert_equal true, is_rotation?("a", "a")
  end

  def test_is_rotation_2_char_true
    assert_equal true, is_rotation?("ab", "ba")
  end

  def test_is_rotation_waterbottle_true
    assert_equal true, is_rotation?("waterbottle", "erbottlewat")
  end

  def test_is_rotation_different_length_false
    assert_equal false, is_rotation?("waterbote", "erbwat")
  end

  def test_is_rotation_different_chars_false
    assert_equal false, is_rotation?("abcdef", "erbwat")
  end
end
