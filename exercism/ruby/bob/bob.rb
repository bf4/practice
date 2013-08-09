class Bob

  attr_reader :drivel

  # @param drivel [String] The drivel Bob is responding to
  # @return response [String] Bob's response to the drivel
  def hey(drivel)
    @drivel = drivel
    case
    when silence?   then "Fine. Be that way!"
    when shouting?  then "Woah, chill out!"
    when asking?    then "Sure."
    else                 "Whatever."
    end
  end

  private

  def shouting?
    drivel.upcase == drivel
  end

  def stating?
    drivel.end_with?('.')
  end

  def asking?
    drivel.end_with?('?')
  end

  def demanding?
    drivel.end_with?('!')
  end

  def silence?
    drivel.nil? || drivel.strip == ''
  end

end
