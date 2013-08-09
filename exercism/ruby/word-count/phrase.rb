class Phrase

  def initialize(word)
    @word = word
  end

  def word_count
    WordCounter.new(@word).word_count
  end

end

WordCounter = Struct.new(:sentance) do

  def words
    @words ||= sentance.split(non_word_expression).reject(&:empty?)
  end

  def word_count
    words.each_with_object(hash_with_default(0)) do |word, counter|
      counter[word.downcase] += 1
    end
  end

  private

  def hash_with_default(value)
    Hash.new {|hash,k| hash[k] = value}
  end

  # see http://www.geocities.jp/kosako3/oniguruma/doc/RE.txt
  def non_word_expression
    /[^A-Za-z0-9]/o
  end
end
